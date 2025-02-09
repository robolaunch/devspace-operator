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

package workspace_manager

import (
	"context"
	goErr "errors"
	"time"

	devspaceErr "github.com/robolaunch/devspace-operator/internal/error"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/go-logr/logr"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
)

// WorkspaceManagerReconciler reconciles a WorkspaceManager object
type WorkspaceManagerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=dev.roboscale.io,resources=workspacemanagers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=dev.roboscale.io,resources=workspacemanagers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=dev.roboscale.io,resources=workspacemanagers/finalizers,verbs=update

//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete

var logger logr.Logger

func (r *WorkspaceManagerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger = log.FromContext(ctx)

	instance, err := r.reconcileGetInstance(ctx, req.NamespacedName)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	err = r.reconcileCheckUpdates(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Check target devspace's other attached objects to see if devspace's resources are released
	err = r.reconcileCheckOtherAttachedResources(ctx, instance)
	if err != nil {
		var e *devspaceErr.DevSpaceResourcesHasNotBeenReleasedError
		if goErr.As(err, &e) {
			return ctrl.Result{
				Requeue:      true,
				RequeueAfter: 3 * time.Second,
			}, nil
		}
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

func (r *WorkspaceManagerReconciler) reconcileCheckStatus(ctx context.Context, instance *devv1alpha1.WorkspaceManager) error {

	switch instance.Status.ClonerJobStatus.Created {
	case true:

		switch instance.Status.ClonerJobStatus.Phase {
		case string(devv1alpha1.JobSucceeded):

			instance.Status.Phase = devv1alpha1.WorkspaceManagerPhaseReady

		case string(devv1alpha1.JobActive):

			instance.Status.Phase = devv1alpha1.WorkspaceManagerPhaseConfiguringWorkspaces

		case string(devv1alpha1.JobFailed):

			instance.Status.Phase = devv1alpha1.WorkspaceManagerPhaseFailed

		}

	case false:

		instance.Status.Phase = devv1alpha1.WorkspaceManagerPhaseConfiguringWorkspaces
		err := r.createClonerJob(ctx, instance, instance.GetClonerJobMetadata())
		if err != nil {
			return err
		}
		instance.Status.ClonerJobStatus.Created = true

	}

	return nil
}

func (r *WorkspaceManagerReconciler) reconcileCheckResources(ctx context.Context, instance *devv1alpha1.WorkspaceManager) error {

	err := r.reconcileCheckClonerJob(ctx, instance)
	if err != nil {
		return err
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *WorkspaceManagerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&devv1alpha1.WorkspaceManager{}).
		Owns(&batchv1.Job{}).
		Complete(r)
}
