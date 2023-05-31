package robot_ide

import (
	"context"

	robotErr "github.com/robolaunch/robot-operator/internal/error"
	"github.com/robolaunch/robot-operator/internal/label"
	robotv1alpha1 "github.com/robolaunch/robot-operator/pkg/api/roboscale.io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *DevSpaceIDEReconciler) reconcileGetInstance(ctx context.Context, meta types.NamespacedName) (*robotv1alpha1.DevSpaceIDE, error) {
	instance := &robotv1alpha1.DevSpaceIDE{}
	err := r.Get(ctx, meta, instance)
	if err != nil {
		return &robotv1alpha1.DevSpaceIDE{}, err
	}

	return instance, nil
}

func (r *DevSpaceIDEReconciler) reconcileUpdateInstanceStatus(ctx context.Context, instance *robotv1alpha1.DevSpaceIDE) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		instanceLV := &robotv1alpha1.DevSpaceIDE{}
		err := r.Get(ctx, types.NamespacedName{
			Name:      instance.Name,
			Namespace: instance.Namespace,
		}, instanceLV)

		if err == nil {
			instance.ResourceVersion = instanceLV.ResourceVersion
		}

		err1 := r.Status().Update(ctx, instance)
		return err1
	})
}

func (r *DevSpaceIDEReconciler) reconcileGetTargetRobot(ctx context.Context, instance *robotv1alpha1.DevSpaceIDE) (*robotv1alpha1.Robot, error) {
	robot := &robotv1alpha1.Robot{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: instance.Namespace,
		Name:      label.GetTargetRobot(instance),
	}, robot)
	if err != nil {
		return nil, err
	}

	return robot, nil
}

func (r *DevSpaceIDEReconciler) reconcileGetTargetRobotVDI(ctx context.Context, instance *robotv1alpha1.DevSpaceIDE) (*robotv1alpha1.RobotVDI, error) {
	robotVDI := &robotv1alpha1.RobotVDI{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: instance.Namespace,
		Name:      label.GetTargetRobotVDI(instance),
	}, robotVDI)
	if err != nil {
		return nil, err
	}

	return robotVDI, nil
}

func (r *DevSpaceIDEReconciler) reconcileCheckNode(ctx context.Context, instance *robotv1alpha1.Robot) (*corev1.Node, error) {

	tenancyMap := label.GetTenancyMap(instance)

	requirements := []labels.Requirement{}
	for k, v := range tenancyMap {
		newReq, err := labels.NewRequirement(k, selection.In, []string{v})
		if err != nil {
			return nil, err
		}
		requirements = append(requirements, *newReq)
	}

	nodeSelector := labels.NewSelector().Add(requirements...)

	nodes := &corev1.NodeList{}
	err := r.List(ctx, nodes, &client.ListOptions{
		LabelSelector: nodeSelector,
	})
	if err != nil {
		return nil, err
	}

	if len(nodes.Items) == 0 {
		return nil, &robotErr.NodeNotFoundError{
			ResourceKind:      instance.Kind,
			ResourceName:      instance.Name,
			ResourceNamespace: instance.Namespace,
		}
	} else if len(nodes.Items) > 1 {
		return nil, &robotErr.MultipleNodeFoundError{
			ResourceKind:      instance.Kind,
			ResourceName:      instance.Name,
			ResourceNamespace: instance.Namespace,
		}
	}

	instance.Status.NodeName = nodes.Items[0].Name

	return &nodes.Items[0], nil
}
