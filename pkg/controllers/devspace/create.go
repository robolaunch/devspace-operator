package devspace

import (
	"context"

	"github.com/robolaunch/devspace-operator/internal/node"
	"github.com/robolaunch/devspace-operator/internal/resources"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *DevspaceReconciler) createPVC(ctx context.Context, instance *devv1alpha1.Devspace, pvcNamespacedName *types.NamespacedName) error {

	pvc := resources.GetPersistentVolumeClaim(instance, pvcNamespacedName)

	err := ctrl.SetControllerReference(instance, pvc, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, pvc)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: PVC " + pvc.Name + " is created.")
	return nil
}

func (r *DevspaceReconciler) createJob(ctx context.Context, instance *devv1alpha1.Devspace, jobNamespacedName *types.NamespacedName) error {

	activeNode, err := r.reconcileCheckNode(ctx, instance)
	if err != nil {
		return err
	}

	job := resources.GetLoaderJob(instance, jobNamespacedName, node.HasGPU(*activeNode))

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

func (r *DevspaceReconciler) createDevSuite(ctx context.Context, instance *devv1alpha1.Devspace, rdsNamespacedName *types.NamespacedName) error {

	devSuite := resources.GetDevSuite(instance, rdsNamespacedName)

	err := ctrl.SetControllerReference(instance, devSuite, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, devSuite)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: Devspace dev suite " + devSuite.Name + " is created.")
	return nil
}

func (r *DevspaceReconciler) createWorkspaceManager(ctx context.Context, instance *devv1alpha1.Devspace, wsmNamespacedName *types.NamespacedName) error {

	workspaceManager := resources.GetWorkspaceManager(instance, wsmNamespacedName)

	err := ctrl.SetControllerReference(instance, workspaceManager, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, workspaceManager)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: Workspace manager " + workspaceManager.Name + " is created.")
	return nil
}
