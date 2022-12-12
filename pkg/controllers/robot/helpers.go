package robot

import (
	"context"
	"sort"

	robotv1alpha1 "github.com/robolaunch/robot-operator/api/v1alpha1"
	"github.com/robolaunch/robot-operator/internal"
	robotErr "github.com/robolaunch/robot-operator/internal/error"
	label "github.com/robolaunch/robot-operator/internal/label"
	nodePkg "github.com/robolaunch/robot-operator/internal/node"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *RobotReconciler) reconcileGetInstance(ctx context.Context, meta types.NamespacedName) (*robotv1alpha1.Robot, error) {
	instance := &robotv1alpha1.Robot{}
	err := r.Get(ctx, meta, instance)
	if err != nil {
		return &robotv1alpha1.Robot{}, err
	}

	return instance, nil
}

func (r *RobotReconciler) reconcileUpdateInstanceStatus(ctx context.Context, instance *robotv1alpha1.Robot) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		instanceLV := &robotv1alpha1.Robot{}
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

func (r *RobotReconciler) reconcileCheckNode(ctx context.Context, instance *robotv1alpha1.Robot) (*corev1.Node, error) {

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

func (r *RobotReconciler) reconcileCheckImage(ctx context.Context, instance *robotv1alpha1.Robot) error {

	node, err := r.reconcileCheckNode(ctx, instance)
	if err != nil {
		return err
	}

	if instance.Status.Image == "" {
		instance.Status.Image = nodePkg.GetImage(*node, *instance)
	}

	return nil
}

func (r *RobotReconciler) reconcileAttachBuildObject(ctx context.Context, instance *robotv1alpha1.Robot) error {

	// Get attached build objects for this robot
	requirements := []labels.Requirement{}
	newReq, err := labels.NewRequirement(internal.TARGET_ROBOT, selection.In, []string{instance.Name})
	if err != nil {
		return err
	}
	requirements = append(requirements, *newReq)

	robotSelector := labels.NewSelector().Add(requirements...)

	buildManagerList := robotv1alpha1.BuildManagerList{}
	err = r.List(ctx, &buildManagerList, &client.ListOptions{Namespace: instance.Namespace, LabelSelector: robotSelector})
	if err != nil {
		return err
	}

	if len(buildManagerList.Items) == 0 {
		instance.Status.AttachedBuildObject.Reference = corev1.ObjectReference{}
		return nil
	}

	// Sort attached build objects for this robot according to their creation timestamps
	sort.SliceStable(buildManagerList.Items[:], func(i, j int) bool {
		return buildManagerList.Items[i].CreationTimestamp.String() > buildManagerList.Items[j].CreationTimestamp.String()
	})

	selectedBuildManager := buildManagerList.Items[0]

	instance.Status.AttachedBuildObject.Reference = corev1.ObjectReference{
		Kind:            selectedBuildManager.Kind,
		Namespace:       selectedBuildManager.Namespace,
		Name:            selectedBuildManager.Name,
		UID:             selectedBuildManager.UID,
		APIVersion:      selectedBuildManager.APIVersion,
		ResourceVersion: selectedBuildManager.ResourceVersion,
	}

	return nil
}
