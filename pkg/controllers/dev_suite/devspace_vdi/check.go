package devspace_vdi

import (
	"context"

	"github.com/robolaunch/devspace-operator/internal/handle"
	"github.com/robolaunch/devspace-operator/internal/reference"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
)

func (r *DevSpaceVDIReconciler) reconcileCheckPVC(ctx context.Context, instance *devv1alpha1.DevSpaceVDI) error {

	pvcQuery := &corev1.PersistentVolumeClaim{}
	err := r.Get(ctx, *instance.GetDevSpaceVDIPVCMetadata(), pvcQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.PVCStatus = devv1alpha1.OwnedResourceStatus{}
		} else {
			return err
		}
	} else {
		instance.Status.PVCStatus.Created = true
		reference.SetReference(&instance.Status.PVCStatus.Reference, pvcQuery.TypeMeta, pvcQuery.ObjectMeta)
	}

	return nil
}

func (r *DevSpaceVDIReconciler) reconcileCheckServices(ctx context.Context, instance *devv1alpha1.DevSpaceVDI) error {

	serviceTCPQuery := &corev1.Service{}
	err := r.Get(ctx, *instance.GetDevSpaceVDIServiceTCPMetadata(), serviceTCPQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.ServiceTCPStatus = devv1alpha1.OwnedServiceStatus{}
		} else {
			return err
		}
	} else {
		devspace, err := r.reconcileGetTargetDevSpace(ctx, instance)
		if err != nil {
			return err
		}

		instance.Status.ServiceTCPStatus.Resource.Created = true
		reference.SetReference(&instance.Status.ServiceTCPStatus.Resource.Reference, serviceTCPQuery.TypeMeta, serviceTCPQuery.ObjectMeta)
		if instance.Spec.Ingress {
			instance.Status.ServiceTCPStatus.URL = devv1alpha1.GetDevSpaceServiceDNS(*devspace, "https://", "/vdi/")
		} else if instance.Spec.ServiceType == corev1.ServiceTypeNodePort {
			// TODO: Address with Node IP and port will be generated.
			instance.Status.ServiceTCPStatus.URL = "http://<NODE-IP>:<PORT>"
		}
	}

	serviceUDPQuery := &corev1.Service{}
	err = r.Get(ctx, *instance.GetDevSpaceVDIServiceUDPMetadata(), serviceUDPQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.ServiceUDPStatus = devv1alpha1.OwnedResourceStatus{}
		} else {
			return err
		}
	} else {
		instance.Status.ServiceUDPStatus.Created = true
		reference.SetReference(&instance.Status.ServiceUDPStatus.Reference, serviceUDPQuery.TypeMeta, serviceUDPQuery.ObjectMeta)
	}

	return nil
}

func (r *DevSpaceVDIReconciler) reconcileCheckPod(ctx context.Context, instance *devv1alpha1.DevSpaceVDI) error {

	vdiPodQuery := &corev1.Pod{}
	err := r.Get(ctx, *instance.GetDevSpaceVDIPodMetadata(), vdiPodQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.PodStatus = devv1alpha1.OwnedPodStatus{}
		} else {
			return err
		}
	} else {

		err = handle.HandlePod(ctx, r, *vdiPodQuery)
		if err != nil {
			return err
		}

		instance.Status.PodStatus.Resource.Created = true
		reference.SetReference(&instance.Status.PodStatus.Resource.Reference, vdiPodQuery.TypeMeta, vdiPodQuery.ObjectMeta)
		instance.Status.PodStatus.Resource.Phase = string(vdiPodQuery.Status.Phase)
		instance.Status.PodStatus.IP = vdiPodQuery.Status.PodIP
	}

	return nil
}

func (r *DevSpaceVDIReconciler) reconcileCheckIngress(ctx context.Context, instance *devv1alpha1.DevSpaceVDI) error {

	if instance.Spec.Ingress {
		ingressQuery := &networkingv1.Ingress{}
		err := r.Get(ctx, *instance.GetDevSpaceVDIIngressMetadata(), ingressQuery)
		if err != nil {
			if errors.IsNotFound(err) {
				instance.Status.IngressStatus = devv1alpha1.OwnedResourceStatus{}
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
