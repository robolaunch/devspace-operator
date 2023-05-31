package devspace_ide

import (
	"context"

	"github.com/robolaunch/devspace-operator/internal/handle"
	"github.com/robolaunch/devspace-operator/internal/reference"
	robotv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
)

func (r *DevSpaceIDEReconciler) reconcileCheckService(ctx context.Context, instance *robotv1alpha1.DevSpaceIDE) error {

	serviceQuery := &corev1.Service{}
	err := r.Get(ctx, *instance.GetDevSpaceIDEServiceMetadata(), serviceQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.ServiceStatus = robotv1alpha1.OwnedServiceStatus{}
		} else {
			return err
		}
	} else {
		robot, err := r.reconcileGetTargetRobot(ctx, instance)
		if err != nil {
			return err
		}

		instance.Status.ServiceStatus.Resource.Created = true
		reference.SetReference(&instance.Status.ServiceStatus.Resource.Reference, serviceQuery.TypeMeta, serviceQuery.ObjectMeta)
		if instance.Spec.Ingress {
			instance.Status.ServiceStatus.URL = robotv1alpha1.GetRobotServiceDNS(*robot, "https://", "/ide/")
		} else if instance.Spec.ServiceType == corev1.ServiceTypeNodePort {
			// TODO: Address with Node IP and port will be generated.
			instance.Status.ServiceStatus.URL = "http://<NODE-IP>:<PORT>"
		}
	}

	return nil
}

func (r *DevSpaceIDEReconciler) reconcileCheckPod(ctx context.Context, instance *robotv1alpha1.DevSpaceIDE) error {

	podQuery := &corev1.Pod{}
	err := r.Get(ctx, *instance.GetDevSpaceIDEPodMetadata(), podQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.PodStatus = robotv1alpha1.OwnedPodStatus{}
		} else {
			return err
		}
	} else {

		err = handle.HandlePod(ctx, r, *podQuery)
		if err != nil {
			return err
		}

		instance.Status.PodStatus.Resource.Created = true
		reference.SetReference(&instance.Status.PodStatus.Resource.Reference, podQuery.TypeMeta, podQuery.ObjectMeta)
		instance.Status.PodStatus.Resource.Phase = string(podQuery.Status.Phase)
		instance.Status.PodStatus.IP = podQuery.Status.PodIP
	}

	return nil
}

func (r *DevSpaceIDEReconciler) reconcileCheckIngress(ctx context.Context, instance *robotv1alpha1.DevSpaceIDE) error {

	if instance.Spec.Ingress {
		ingressQuery := &networkingv1.Ingress{}
		err := r.Get(ctx, *instance.GetDevSpaceIDEIngressMetadata(), ingressQuery)
		if err != nil {
			if errors.IsNotFound(err) {
				instance.Status.IngressStatus = robotv1alpha1.OwnedResourceStatus{}
			} else {
				return err
			}
		} else {
			instance.Status.IngressStatus.Created = true
			reference.SetReference(&instance.Status.IngressStatus.Reference, ingressQuery.TypeMeta, ingressQuery.ObjectMeta)
		}
	}

	return nil
}
