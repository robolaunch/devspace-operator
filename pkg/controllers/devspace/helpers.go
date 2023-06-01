package devspace

import (
	"context"

	devspaceErr "github.com/robolaunch/devspace-operator/internal/error"
	label "github.com/robolaunch/devspace-operator/internal/label"
	nodePkg "github.com/robolaunch/devspace-operator/internal/node"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *DevSpaceReconciler) reconcileGetInstance(ctx context.Context, meta types.NamespacedName) (*devv1alpha1.DevSpace, error) {
	instance := &devv1alpha1.DevSpace{}
	err := r.Get(ctx, meta, instance)
	if err != nil {
		return &devv1alpha1.DevSpace{}, err
	}

	return instance, nil
}

func (r *DevSpaceReconciler) reconcileUpdateInstanceStatus(ctx context.Context, instance *devv1alpha1.DevSpace) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		instanceLV := &devv1alpha1.DevSpace{}
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

func (r *DevSpaceReconciler) reconcileCheckNode(ctx context.Context, instance *devv1alpha1.DevSpace) (*corev1.Node, error) {

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
		return nil, &devspaceErr.NodeNotFoundError{
			ResourceKind:      instance.Kind,
			ResourceName:      instance.Name,
			ResourceNamespace: instance.Namespace,
		}
	} else if len(nodes.Items) > 1 {
		return nil, &devspaceErr.MultipleNodeFoundError{
			ResourceKind:      instance.Kind,
			ResourceName:      instance.Name,
			ResourceNamespace: instance.Namespace,
		}
	}

	instance.Status.NodeName = nodes.Items[0].Name

	return &nodes.Items[0], nil
}

func (r *DevSpaceReconciler) reconcileCheckImage(ctx context.Context, instance *devv1alpha1.DevSpace) error {

	node, err := r.reconcileCheckNode(ctx, instance)
	if err != nil {
		return err
	}

	if instance.Status.Image == "" {
		instance.Status.Image, err = nodePkg.GetImage(*node, *instance)
		if err != nil {
			return err
		}
	}

	return nil
}
