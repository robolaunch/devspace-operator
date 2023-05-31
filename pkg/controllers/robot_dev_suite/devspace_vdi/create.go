package devspace_vdi

import (
	"context"

	"github.com/robolaunch/devspace-operator/internal/resources"
	robotv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *DevSpaceVDIReconciler) reconcileCreatePVC(ctx context.Context, instance *robotv1alpha1.DevSpaceVDI) error {

	robot, err := r.reconcileGetTargetRobot(ctx, instance)
	if err != nil {
		return err
	}

	vdiPVC := resources.GetDevSpaceVDIPVC(instance, instance.GetDevSpaceVDIPVCMetadata(), *robot)

	err = ctrl.SetControllerReference(instance, vdiPVC, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, vdiPVC)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: VDI PVC is created.")

	return nil
}

func (r *DevSpaceVDIReconciler) reconcileCreateServiceTCP(ctx context.Context, instance *robotv1alpha1.DevSpaceVDI) error {

	vdiService := resources.GetDevSpaceVDIServiceTCP(instance, instance.GetDevSpaceVDIServiceTCPMetadata())

	err := ctrl.SetControllerReference(instance, vdiService, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, vdiService)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: VDI TCP service is created.")

	return nil
}

func (r *DevSpaceVDIReconciler) reconcileCreateServiceUDP(ctx context.Context, instance *robotv1alpha1.DevSpaceVDI) error {

	vdiService := resources.GetDevSpaceVDIServiceUDP(instance, instance.GetDevSpaceVDIServiceUDPMetadata())

	err := ctrl.SetControllerReference(instance, vdiService, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, vdiService)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: VDI UDP service is created.")

	return nil
}

func (r *DevSpaceVDIReconciler) reconcileCreatePod(ctx context.Context, instance *robotv1alpha1.DevSpaceVDI) error {

	robot, err := r.reconcileGetTargetRobot(ctx, instance)
	if err != nil {
		return err
	}

	activeNode, err := r.reconcileCheckNode(ctx, robot)
	if err != nil {
		return err
	}

	vdiPod := resources.GetDevSpaceVDIPod(instance, instance.GetDevSpaceVDIPodMetadata(), *robot, *activeNode)

	err = ctrl.SetControllerReference(instance, vdiPod, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, vdiPod)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: VDI pod is created.")

	return nil
}

func (r *DevSpaceVDIReconciler) reconcileCreateIngress(ctx context.Context, instance *robotv1alpha1.DevSpaceVDI) error {

	robot, err := r.reconcileGetTargetRobot(ctx, instance)
	if err != nil {
		return err
	}

	vdiIngress := resources.GetDevSpaceVDIIngress(instance, instance.GetDevSpaceVDIPodMetadata(), *robot)

	err = ctrl.SetControllerReference(instance, vdiIngress, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, vdiIngress)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: VDI ingress is created.")

	return nil
}
