package workspace_manager

import (
	"context"

	"github.com/robolaunch/devspace-operator/internal/label"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
)

func (r *WorkspaceManagerReconciler) reconcileGetInstance(ctx context.Context, meta types.NamespacedName) (*devv1alpha1.WorkspaceManager, error) {
	instance := &devv1alpha1.WorkspaceManager{}
	err := r.Get(ctx, meta, instance)
	if err != nil {
		return &devv1alpha1.WorkspaceManager{}, err
	}

	return instance, nil
}

func (r *WorkspaceManagerReconciler) reconcileUpdateInstanceStatus(ctx context.Context, instance *devv1alpha1.WorkspaceManager) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		instanceLV := &devv1alpha1.WorkspaceManager{}
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

func (r *WorkspaceManagerReconciler) reconcileGetTargetDevSpace(ctx context.Context, instance *devv1alpha1.WorkspaceManager) (*devv1alpha1.DevSpace, error) {
	devspace := &devv1alpha1.DevSpace{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: instance.Namespace,
		Name:      label.GetTargetDevSpace(instance),
	}, devspace)
	if err != nil {
		return nil, err
	}

	return devspace, nil
}

func (r *WorkspaceManagerReconciler) reconcileCleanup(ctx context.Context, instance *devv1alpha1.WorkspaceManager) error {

	err := r.reconcileDeleteClonerJob(ctx, instance)
	if err != nil {
		return err
	}

	err = r.reconcileDeleteCleanupJob(ctx, instance)
	if err != nil {
		return err
	}

	err = r.createCleanupJob(ctx, instance, instance.GetCleanupJobMetadata())
	if err != nil {
		return err
	}

	err = r.reconcileCheckCleanupJob(ctx, instance)
	if err != nil {
		return err
	}

	err = r.reconcileDeleteCleanupJob(ctx, instance)
	if err != nil {
		return err
	}
	return nil
}
