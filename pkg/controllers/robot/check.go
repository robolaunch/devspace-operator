package robot

import (
	"context"
	"reflect"

	"github.com/robolaunch/devspace-operator/internal/reference"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
)

func (r *DevspaceReconciler) reconcileCheckPVCs(ctx context.Context, instance *devv1alpha1.Devspace) error {

	pvcVarQuery := &corev1.PersistentVolumeClaim{}
	err := r.Get(ctx, *instance.GetPVCVarMetadata(), pvcVarQuery)
	if err != nil && errors.IsNotFound(err) {
		instance.Status.VolumeStatuses.Var = devv1alpha1.OwnedResourceStatus{}
	} else if err != nil {
		return err
	} else {
		instance.Status.VolumeStatuses.Var.Created = true
		reference.SetReference(&instance.Status.VolumeStatuses.Var.Reference, pvcVarQuery.TypeMeta, pvcVarQuery.ObjectMeta)
	}

	pvcOptQuery := &corev1.PersistentVolumeClaim{}
	err = r.Get(ctx, *instance.GetPVCOptMetadata(), pvcOptQuery)
	if err != nil && errors.IsNotFound(err) {
		instance.Status.VolumeStatuses.Opt = devv1alpha1.OwnedResourceStatus{}
	} else if err != nil {
		return err
	} else {
		instance.Status.VolumeStatuses.Opt.Created = true
		reference.SetReference(&instance.Status.VolumeStatuses.Opt.Reference, pvcOptQuery.TypeMeta, pvcOptQuery.ObjectMeta)
	}

	pvcEtcQuery := &corev1.PersistentVolumeClaim{}
	err = r.Get(ctx, *instance.GetPVCEtcMetadata(), pvcEtcQuery)
	if err != nil && errors.IsNotFound(err) {
		instance.Status.VolumeStatuses.Etc = devv1alpha1.OwnedResourceStatus{}
	} else if err != nil {
		return err
	} else {
		instance.Status.VolumeStatuses.Etc.Created = true
		reference.SetReference(&instance.Status.VolumeStatuses.Etc.Reference, pvcEtcQuery.TypeMeta, pvcEtcQuery.ObjectMeta)
	}

	pvcUsrQuery := &corev1.PersistentVolumeClaim{}
	err = r.Get(ctx, *instance.GetPVCUsrMetadata(), pvcUsrQuery)
	if err != nil && errors.IsNotFound(err) {
		instance.Status.VolumeStatuses.Usr = devv1alpha1.OwnedResourceStatus{}
	} else if err != nil {
		return err
	} else {
		instance.Status.VolumeStatuses.Usr.Created = true
		reference.SetReference(&instance.Status.VolumeStatuses.Usr.Reference, pvcUsrQuery.TypeMeta, pvcUsrQuery.ObjectMeta)
	}

	pvcWorkspaceQuery := &corev1.PersistentVolumeClaim{}
	err = r.Get(ctx, *instance.GetPVCWorkspaceMetadata(), pvcWorkspaceQuery)
	if err != nil && errors.IsNotFound(err) {
		instance.Status.VolumeStatuses.Workspace = devv1alpha1.OwnedResourceStatus{}
	} else if err != nil {
		return err
	} else {
		instance.Status.VolumeStatuses.Workspace.Created = true
		reference.SetReference(&instance.Status.VolumeStatuses.Workspace.Reference, pvcWorkspaceQuery.TypeMeta, pvcWorkspaceQuery.ObjectMeta)
	}

	return nil
}

func (r *DevspaceReconciler) reconcileCheckLoaderJob(ctx context.Context, instance *devv1alpha1.Devspace) error {

	if instance.Status.Phase != devv1alpha1.DevspacePhaseEnvironmentReady {
		loaderJobQuery := &batchv1.Job{}
		err := r.Get(ctx, *instance.GetLoaderJobMetadata(), loaderJobQuery)
		if err != nil && errors.IsNotFound(err) {
			instance.Status.LoaderJobStatus = devv1alpha1.OwnedResourceStatus{}
		} else if err != nil {
			return err
		} else {
			reference.SetReference(&instance.Status.LoaderJobStatus.Reference, loaderJobQuery.TypeMeta, loaderJobQuery.ObjectMeta)
			switch 1 {
			case int(loaderJobQuery.Status.Succeeded):
				instance.Status.LoaderJobStatus.Phase = string(devv1alpha1.JobSucceeded)
			case int(loaderJobQuery.Status.Active):
				instance.Status.LoaderJobStatus.Phase = string(devv1alpha1.JobActive)
			case int(loaderJobQuery.Status.Failed):
				instance.Status.LoaderJobStatus.Phase = string(devv1alpha1.JobFailed)
			}
		}
	}

	return nil
}

func (r *DevspaceReconciler) reconcileCheckDevSuite(ctx context.Context, instance *devv1alpha1.Devspace) error {

	devSuiteQuery := &devv1alpha1.DevSuite{}
	err := r.Get(ctx, *instance.GetDevSuiteMetadata(), devSuiteQuery)
	if err != nil && errors.IsNotFound(err) {
		instance.Status.DevSuiteStatus = devv1alpha1.DevSuiteInstanceStatus{}
	} else if err != nil {
		return err
	} else {

		if instance.Spec.DevSuiteTemplate.IDEEnabled || instance.Spec.DevSuiteTemplate.VDIEnabled {

			if !reflect.DeepEqual(instance.Spec.DevSuiteTemplate, devSuiteQuery.Spec) {
				devSuiteQuery.Spec = instance.Spec.DevSuiteTemplate
				err = r.Update(ctx, devSuiteQuery)
				if err != nil {
					return err
				}
			}

			instance.Status.DevSuiteStatus.Resource.Created = true
			reference.SetReference(&instance.Status.DevSuiteStatus.Resource.Reference, devSuiteQuery.TypeMeta, devSuiteQuery.ObjectMeta)
			instance.Status.DevSuiteStatus.Status = devSuiteQuery.Status

		} else {

			err := r.Delete(ctx, devSuiteQuery)
			if err != nil {
				return err
			}

		}

	}

	return nil
}

func (r *DevspaceReconciler) reconcileCheckWorkspaceManager(ctx context.Context, instance *devv1alpha1.Devspace) error {

	workspaceManagerQuery := &devv1alpha1.WorkspaceManager{}
	err := r.Get(ctx, *instance.GetWorkspaceManagerMetadata(), workspaceManagerQuery)
	if err != nil && errors.IsNotFound(err) {
		instance.Status.WorkspaceManagerStatus = devv1alpha1.WorkspaceManagerInstanceStatus{}
	} else if err != nil {
		return err
	} else {

		instance.Status.WorkspaceManagerStatus.Resource.Created = true
		reference.SetReference(&instance.Status.WorkspaceManagerStatus.Resource.Reference, workspaceManagerQuery.TypeMeta, workspaceManagerQuery.ObjectMeta)
		instance.Status.WorkspaceManagerStatus.Status = workspaceManagerQuery.Status

		if !reflect.DeepEqual(instance.Spec.WorkspaceManagerTemplate.Workspaces, workspaceManagerQuery.Spec.Workspaces) {
			workspaceManagerQuery.Spec = instance.Spec.WorkspaceManagerTemplate
			workspaceManagerQuery.Spec.UpdateNeeded = true
			err = r.Update(ctx, workspaceManagerQuery)
			if err != nil {
				return err
			}

			// set phase configuring
			instance.Status.WorkspaceManagerStatus.Resource.Created = true
			instance.Status.WorkspaceManagerStatus.Status = devv1alpha1.WorkspaceManagerStatus{}
			instance.Status.WorkspaceManagerStatus.Status.Phase = devv1alpha1.WorkspaceManagerPhaseConfiguringWorkspaces
		}

	}

	return nil
}
