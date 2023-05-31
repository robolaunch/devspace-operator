package workspace_manager

import (
	"context"

	"github.com/robolaunch/devspace-operator/internal"
	robotErr "github.com/robolaunch/devspace-operator/internal/error"
	"github.com/robolaunch/devspace-operator/internal/label"
	"github.com/robolaunch/devspace-operator/internal/reference"
	robotv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *WorkspaceManagerReconciler) reconcileCheckClonerJob(ctx context.Context, instance *robotv1alpha1.WorkspaceManager) error {

	clonerJobQuery := &batchv1.Job{}
	err := r.Get(ctx, *instance.GetClonerJobMetadata(), clonerJobQuery)
	if err != nil && errors.IsNotFound(err) {
		instance.Status.ClonerJobStatus = robotv1alpha1.OwnedResourceStatus{}
	} else if err != nil {
		return err
	} else {
		reference.SetReference(&instance.Status.ClonerJobStatus.Reference, clonerJobQuery.TypeMeta, clonerJobQuery.ObjectMeta)
		switch 1 {
		case int(clonerJobQuery.Status.Succeeded):
			instance.Status.ClonerJobStatus.Phase = string(robotv1alpha1.JobSucceeded)
		case int(clonerJobQuery.Status.Active):
			instance.Status.ClonerJobStatus.Phase = string(robotv1alpha1.JobActive)
		case int(clonerJobQuery.Status.Failed):
			instance.Status.ClonerJobStatus.Phase = string(robotv1alpha1.JobFailed)
		}
	}

	return nil
}

func (r *WorkspaceManagerReconciler) reconcileCheckCleanupJob(ctx context.Context, instance *robotv1alpha1.WorkspaceManager) error {

	isActive := true
	for isActive {

		cleanupJobQuery := &batchv1.Job{}
		err := r.Get(ctx, *instance.GetCleanupJobMetadata(), cleanupJobQuery)
		if err != nil && errors.IsNotFound(err) {
			instance.Status.CleanupJobStatus = robotv1alpha1.OwnedResourceStatus{}
		} else if err != nil {
			return err
		} else {
			reference.SetReference(&instance.Status.CleanupJobStatus.Reference, cleanupJobQuery.TypeMeta, cleanupJobQuery.ObjectMeta)
			switch 1 {
			case int(cleanupJobQuery.Status.Succeeded):
				instance.Status.CleanupJobStatus.Phase = string(robotv1alpha1.JobSucceeded)
				isActive = false
			case int(cleanupJobQuery.Status.Active):
				instance.Status.CleanupJobStatus.Phase = string(robotv1alpha1.JobActive)
			case int(cleanupJobQuery.Status.Failed):
				instance.Status.CleanupJobStatus.Phase = string(robotv1alpha1.JobFailed)
				isActive = false
			}
		}

	}

	return nil
}

func (r *WorkspaceManagerReconciler) reconcileCheckOtherAttachedResources(ctx context.Context, instance *robotv1alpha1.WorkspaceManager) error {

	// Get attached build manager objects for this robot
	requirements := []labels.Requirement{}
	targetReq, err := labels.NewRequirement(internal.TARGET_ROBOT_LABEL_KEY, selection.In, []string{label.GetTargetRobot(instance)})
	if err != nil {
		return err
	}

	ownedReq, err := labels.NewRequirement(internal.ROBOT_DEV_SUITE_OWNED, selection.DoesNotExist, []string{})
	if err != nil {
		return err
	}
	requirements = append(requirements, *targetReq, *ownedReq)

	robotSelector := labels.NewSelector().Add(requirements...)

	devSuiteList := robotv1alpha1.DevSuiteList{}
	err = r.List(ctx, &devSuiteList, &client.ListOptions{Namespace: instance.Namespace, LabelSelector: robotSelector.Add()})
	if err != nil {
		return err
	}

	for _, rds := range devSuiteList.Items {

		if rds.Status.Active {
			return &robotErr.RobotResourcesHasNotBeenReleasedError{
				ResourceKind:      instance.Kind,
				ResourceName:      instance.Name,
				ResourceNamespace: instance.Namespace,
			}
		}

		if rds.Status.Phase != robotv1alpha1.DevSuitePhaseInactive {
			return &robotErr.RobotResourcesHasNotBeenReleasedError{
				ResourceKind:      instance.Kind,
				ResourceName:      instance.Name,
				ResourceNamespace: instance.Namespace,
			}
		}
	}

	return nil
}

func (r *WorkspaceManagerReconciler) reconcileCheckUpdates(ctx context.Context, instance *robotv1alpha1.WorkspaceManager) error {

	switch instance.Spec.UpdateNeeded {
	case true:

		instance.Spec.UpdateNeeded = false
		err := r.Update(ctx, instance, &client.UpdateOptions{})
		if err != nil {
			return err
		}

		instance.Status.Version++
		instance.Status.Phase = robotv1alpha1.WorkspaceManagerPhaseConfiguringWorkspaces
		instance.Status.CleanupJobStatus = robotv1alpha1.OwnedResourceStatus{}
		instance.Status.ClonerJobStatus = robotv1alpha1.OwnedResourceStatus{}

		err = r.reconcileUpdateInstanceStatus(ctx, instance)
		if err != nil {
			return err
		}

		err = r.reconcileCleanup(ctx, instance)
		if err != nil {
			return err
		}

	}

	return nil
}
