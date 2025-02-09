/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package devspace_vdi

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/go-logr/logr"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
)

// DevSpaceVDIReconciler reconciles a DevSpaceVDI object
type DevSpaceVDIReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	DynamicClient dynamic.Interface
}

//+kubebuilder:rbac:groups=dev.roboscale.io,resources=devspacevdis,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=dev.roboscale.io,resources=devspacevdis/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=dev.roboscale.io,resources=devspacevdis/finalizers,verbs=update

//+kubebuilder:rbac:groups=core,resources=persistentvolumeclaims,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete

var logger logr.Logger

func (r *DevSpaceVDIReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger = log.FromContext(ctx)

	instance, err := r.reconcileGetInstance(ctx, req.NamespacedName)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	if !instance.DeletionTimestamp.IsZero() {
		return ctrl.Result{}, nil
	}

	err = r.reconcileCheckStatus(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.reconcileUpdateInstanceStatus(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.reconcileCheckResources(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.reconcileUpdateInstanceStatus(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *DevSpaceVDIReconciler) reconcileCheckStatus(ctx context.Context, instance *devv1alpha1.DevSpaceVDI) error {

	switch instance.Status.PVCStatus.Created {
	case true:

		switch instance.Status.ServiceTCPStatus.Resource.Created {
		case true:

			switch instance.Status.ServiceUDPStatus.Created {
			case true:

				switch instance.Status.PodStatus.Resource.Created {
				case true:

					switch instance.Status.PodStatus.Resource.Phase {
					case string(corev1.PodRunning):

						switch instance.Spec.Ingress {
						case true:

							switch instance.Status.IngressStatus.Created {
							case true:

								instance.Status.Phase = devv1alpha1.DevSpaceVDIPhaseRunning

							case false:

								instance.Status.Phase = devv1alpha1.DevSpaceVDIPhaseCreatingIngress
								err := r.reconcileCreateIngress(ctx, instance)
								if err != nil {
									return err
								}
								instance.Status.IngressStatus.Created = true

							}

						case false:

							instance.Status.Phase = devv1alpha1.DevSpaceVDIPhaseRunning

						}

					}

				case false:

					instance.Status.Phase = devv1alpha1.DevSpaceVDIPhaseCreatingPod
					err := r.reconcileCreatePod(ctx, instance)
					if err != nil {
						return err
					}
					instance.Status.PodStatus.Resource.Created = true

				}

			case false:

				instance.Status.Phase = devv1alpha1.DevSpaceVDIPhaseCreatingUDPService
				err := r.reconcileCreateServiceUDP(ctx, instance)
				if err != nil {
					return err
				}
				instance.Status.ServiceUDPStatus.Created = true

			}

		case false:

			instance.Status.Phase = devv1alpha1.DevSpaceVDIPhaseCreatingTCPService
			err := r.reconcileCreateServiceTCP(ctx, instance)
			if err != nil {
				return err
			}
			instance.Status.ServiceTCPStatus.Resource.Created = true

		}

	case false:

		instance.Status.Phase = devv1alpha1.DevSpaceVDIPhaseCreatingPVC
		err := r.reconcileCreatePVC(ctx, instance)
		if err != nil {
			return err
		}
		instance.Status.PVCStatus.Created = true

	}

	return nil
}

func (r *DevSpaceVDIReconciler) reconcileCheckResources(ctx context.Context, instance *devv1alpha1.DevSpaceVDI) error {

	err := r.reconcileCheckPVC(ctx, instance)
	if err != nil {
		return err
	}

	err = r.reconcileCheckServices(ctx, instance)
	if err != nil {
		return err
	}

	err = r.reconcileCheckPod(ctx, instance)
	if err != nil {
		return err
	}

	err = r.reconcileCheckIngress(ctx, instance)
	if err != nil {
		return err
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DevSpaceVDIReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&devv1alpha1.DevSpaceVDI{}).
		Owns(&corev1.PersistentVolumeClaim{}).
		Owns(&corev1.Service{}).
		Owns(&corev1.Pod{}).
		Owns(&networkingv1.Ingress{}).
		Complete(r)
}
