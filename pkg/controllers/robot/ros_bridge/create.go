package ros_bridge

import (
	"context"

	"github.com/robolaunch/robot-operator/internal/resources"
	robotv1alpha1 "github.com/robolaunch/robot-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *ROSBridgeReconciler) createService(ctx context.Context, instance *robotv1alpha1.ROSBridge, svcNamespacedName *types.NamespacedName) error {

	svc := resources.GetBridgeService(instance, instance.GetBridgeServiceMetadata())

	err := ctrl.SetControllerReference(instance, svc, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, svc)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: Service " + svc.Name + " is created.")
	return nil
}

func (r *ROSBridgeReconciler) createPod(ctx context.Context, instance *robotv1alpha1.ROSBridge, podNamespacedName *types.NamespacedName) error {

	robot, err := r.reconcileGetOwner(ctx, instance)
	if err != nil {
		return err
	}

	pod := resources.GetBridgePod(instance, instance.GetBridgePodMetadata(), *robot)

	err = ctrl.SetControllerReference(instance, pod, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, pod)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: Pod " + pod.Name + " is created.")
	return nil
}

func (r *ROSBridgeReconciler) createIngress(ctx context.Context, instance *robotv1alpha1.ROSBridge, ingressNamespacedName *types.NamespacedName) error {

	robot, err := r.reconcileGetOwner(ctx, instance)
	if err != nil {
		return err
	}

	ingress := resources.GetBridgeIngress(instance, instance.GetBridgeIngressMetadata(), *robot)

	err = ctrl.SetControllerReference(instance, ingress, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, ingress)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: Ingress " + ingress.Name + " is created.")
	return nil
}
