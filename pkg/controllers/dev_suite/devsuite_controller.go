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

package dev_suite

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/go-logr/logr"
	"github.com/robolaunch/devspace-operator/internal"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
)

// DevSuiteReconciler reconciles a DevSuite object
type DevSuiteReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	DynamicClient dynamic.Interface
}

//+kubebuilder:rbac:groups=dev.roboscale.io,resources=devsuites,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=dev.roboscale.io,resources=devsuites/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=dev.roboscale.io,resources=devsuites/finalizers,verbs=update

var logger logr.Logger

func (r *DevSuiteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger = log.FromContext(ctx)

	instance, err := r.reconcileGetInstance(ctx, req.NamespacedName)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// Check target robot's attached object, update activity status
	err = r.reconcileCheckTargetDevspace(ctx, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.Phase = devv1alpha1.DevSuitePhaseDevspaceNotFound
			instance.Status.Active = false
		} else {
			return ctrl.Result{}, err
		}
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

func (r *DevSuiteReconciler) reconcileCheckStatus(ctx context.Context, instance *devv1alpha1.DevSuite) error {

	switch instance.Status.Active {
	case true:

		switch instance.Spec.VDIEnabled {
		case true:

			switch instance.Status.DevSpaceVDIStatus.Resource.Created {
			case true:

				switch instance.Status.DevSpaceVDIStatus.Resource.Phase {
				case string(devv1alpha1.DevSpaceVDIPhaseRunning):

					switch instance.Spec.IDEEnabled {
					case true:

						switch instance.Status.DevSpaceIDEStatus.Resource.Created {
						case true:

							switch instance.Status.DevSpaceIDEStatus.Resource.Phase {
							case string(devv1alpha1.DevSpaceIDEPhaseRunning):

								instance.Status.Phase = devv1alpha1.DevSuitePhaseRunning

							}

						case false:

							instance.Status.Phase = devv1alpha1.DevSuitePhaseCreatingDevSpaceIDE
							err := r.reconcileCreateDevSpaceIDE(ctx, instance)
							if err != nil {
								return err
							}
							instance.Status.DevSpaceIDEStatus.Resource.Created = true

						}

					case false:

						instance.Status.Phase = devv1alpha1.DevSuitePhaseRunning

					}

				}

			case false:

				instance.Status.Phase = devv1alpha1.DevSuitePhaseCreatingDevSpaceVDI
				err := r.reconcileCreateDevSpaceVDI(ctx, instance)
				if err != nil {
					return err
				}
				instance.Status.DevSpaceVDIStatus.Resource.Created = true

			}

		case false:

			switch instance.Spec.IDEEnabled {
			case true:

				switch instance.Status.DevSpaceIDEStatus.Resource.Created {
				case true:

					switch instance.Status.DevSpaceIDEStatus.Resource.Phase {
					case string(devv1alpha1.DevSpaceIDEPhaseRunning):

						instance.Status.Phase = devv1alpha1.DevSuitePhaseRunning

					}

				case false:

					instance.Status.Phase = devv1alpha1.DevSuitePhaseCreatingDevSpaceIDE
					err := r.reconcileCreateDevSpaceIDE(ctx, instance)
					if err != nil {
						return err
					}
					instance.Status.DevSpaceIDEStatus.Resource.Created = true

				}

			case false:

				instance.Status.Phase = devv1alpha1.DevSuitePhaseRunning

			}

		}

	case false:

		instance.Status.Phase = devv1alpha1.DevSuitePhaseDeactivating

		err := r.reconcileDeleteDevSpaceIDE(ctx, instance)
		if err != nil {
			return err
		}

		err = r.reconcileDeleteDevSpaceVDI(ctx, instance)
		if err != nil {
			return err
		}

		instance.Status.Phase = devv1alpha1.DevSuitePhaseInactive

	}

	return nil
}

func (r *DevSuiteReconciler) reconcileCheckResources(ctx context.Context, instance *devv1alpha1.DevSuite) error {

	err := r.reconcileCheckDevSpaceVDI(ctx, instance)
	if err != nil {
		return err
	}

	err = r.reconcileCheckDevSpaceIDE(ctx, instance)
	if err != nil {
		return err
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DevSuiteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&devv1alpha1.DevSuite{}).
		Owns(&devv1alpha1.DevSpaceVDI{}).
		Owns(&devv1alpha1.DevSpaceIDE{}).
		Watches(
			&source.Kind{Type: &devv1alpha1.Devspace{}},
			handler.EnqueueRequestsFromMapFunc(r.watchDevspaces),
		).
		Complete(r)
}

func (r *DevSuiteReconciler) watchDevspaces(o client.Object) []reconcile.Request {

	robot := o.(*devv1alpha1.Devspace)

	// Get attached build objects for this robot
	requirements := []labels.Requirement{}
	newReq, err := labels.NewRequirement(internal.TARGET_ROBOT_LABEL_KEY, selection.In, []string{robot.Name})
	if err != nil {
		return []reconcile.Request{}
	}
	requirements = append(requirements, *newReq)

	robotSelector := labels.NewSelector().Add(requirements...)

	devSuiteList := devv1alpha1.DevSuiteList{}
	err = r.List(context.TODO(), &devSuiteList, &client.ListOptions{Namespace: robot.Namespace, LabelSelector: robotSelector})
	if err != nil {
		return []reconcile.Request{}
	}

	requests := make([]reconcile.Request, len(devSuiteList.Items))
	for i, item := range devSuiteList.Items {

		requests[i] = reconcile.Request{
			NamespacedName: types.NamespacedName{
				Name:      item.Name,
				Namespace: item.Namespace,
			},
		}

	}

	return requests
}
