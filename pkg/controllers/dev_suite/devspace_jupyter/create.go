package devspace_jupyter

import (
	"context"

	"github.com/robolaunch/devspace-operator/internal/label"
	"github.com/robolaunch/devspace-operator/internal/resources"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *DevSpaceJupyterReconciler) reconcileCreateService(ctx context.Context, instance *devv1alpha1.DevSpaceJupyter) error {

	jupyterService := resources.GetDevSpaceJupyterService(instance, instance.GetDevSpaceJupyterServiceMetadata())

	err := ctrl.SetControllerReference(instance, jupyterService, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, jupyterService)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: Jupyter service is created.")

	return nil
}

func (r *DevSpaceJupyterReconciler) reconcileCreatePod(ctx context.Context, instance *devv1alpha1.DevSpaceJupyter) error {

	devspace, err := r.reconcileGetTargetDevSpace(ctx, instance)
	if err != nil {
		return err
	}

	devSpaceVDI := &devv1alpha1.DevSpaceVDI{}
	if label.GetTargetDevSpaceVDI(instance) != "" {
		devSpaceVDI, err = r.reconcileGetTargetDevSpaceVDI(ctx, instance)
		if err != nil {
			return err
		}
	}

	activeNode, err := r.reconcileCheckNode(ctx, devspace)
	if err != nil {
		return err
	}

	jupyterPod := resources.GetDevSpaceJupyterPod(instance, instance.GetDevSpaceJupyterPodMetadata(), *devspace, *devSpaceVDI, *activeNode)

	err = ctrl.SetControllerReference(instance, jupyterPod, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, jupyterPod)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: Jupyter pod is created.")

	return nil
}

func (r *DevSpaceJupyterReconciler) reconcileCreateIngress(ctx context.Context, instance *devv1alpha1.DevSpaceJupyter) error {

	devspace, err := r.reconcileGetTargetDevSpace(ctx, instance)
	if err != nil {
		return err
	}

	jupyterIngress := resources.GetDevSpaceJupyterIngress(instance, instance.GetDevSpaceJupyterIngressMetadata(), *devspace)

	err = ctrl.SetControllerReference(instance, jupyterIngress, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, jupyterIngress)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: Jupyter ingress is created.")

	return nil
}
