package workspace_manager

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/robolaunch/devspace-operator/internal/resources"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

func (r *WorkspaceManagerReconciler) createClonerJob(ctx context.Context, instance *devv1alpha1.WorkspaceManager, jobNamespacedName *types.NamespacedName) error {

	devspace, err := r.reconcileGetTargetDevSpace(ctx, instance)
	if err != nil {
		return err
	}

	job := resources.GetClonerJob(instance, jobNamespacedName, devspace)

	err = ctrl.SetControllerReference(instance, job, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, job)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: Job " + job.Name + " is created.")
	return nil
}

func (r *WorkspaceManagerReconciler) createCleanupJob(ctx context.Context, instance *devv1alpha1.WorkspaceManager, jobNamespacedName *types.NamespacedName) error {

	devspace, err := r.reconcileGetTargetDevSpace(ctx, instance)
	if err != nil {
		return err
	}

	job := resources.GetCleanupJob(instance, jobNamespacedName, devspace)

	err = ctrl.SetControllerReference(instance, job, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, job)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: Job " + job.Name + " is created.")
	return nil
}
