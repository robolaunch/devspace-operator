package devspace_ide

import (
	"context"

	"github.com/robolaunch/devspace-operator/internal/label"
	"github.com/robolaunch/devspace-operator/internal/resources"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *DevSpaceIDEReconciler) reconcileCreateService(ctx context.Context, instance *devv1alpha1.DevSpaceIDE) error {

	ideService := resources.GetDevSpaceIDEService(instance, instance.GetDevSpaceIDEServiceMetadata())

	err := ctrl.SetControllerReference(instance, ideService, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, ideService)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: IDE service is created.")

	return nil
}

func (r *DevSpaceIDEReconciler) reconcileCreatePod(ctx context.Context, instance *devv1alpha1.DevSpaceIDE) error {

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

	idePod := resources.GetDevSpaceIDEPod(instance, instance.GetDevSpaceIDEPodMetadata(), *devspace, *devSpaceVDI, *activeNode)

	err = ctrl.SetControllerReference(instance, idePod, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, idePod)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: IDE pod is created.")

	return nil
}

func (r *DevSpaceIDEReconciler) reconcileCreateIngress(ctx context.Context, instance *devv1alpha1.DevSpaceIDE) error {

	devspace, err := r.reconcileGetTargetDevSpace(ctx, instance)
	if err != nil {
		return err
	}

	ideIngress := resources.GetDevSpaceIDEIngress(instance, instance.GetDevSpaceIDEIngressMetadata(), *devspace)

	err = ctrl.SetControllerReference(instance, ideIngress, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, ideIngress)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: IDE ingress is created.")

	return nil
}
