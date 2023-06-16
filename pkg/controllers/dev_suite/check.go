package dev_suite

import (
	"context"
	"reflect"

	"github.com/robolaunch/devspace-operator/internal/reference"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
)

func (r *DevSuiteReconciler) reconcileCheckDevSpaceVDI(ctx context.Context, instance *devv1alpha1.DevSuite) error {

	devSpaceVDIQuery := &devv1alpha1.DevSpaceVDI{}
	err := r.Get(ctx, *instance.GetDevSpaceVDIMetadata(), devSpaceVDIQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.DevSpaceVDIStatus = devv1alpha1.OwnedDevSpaceServiceStatus{}
		} else {
			return err
		}
	} else {

		if instance.Spec.VDIEnabled {

			if !reflect.DeepEqual(instance.Spec.DevSpaceVDITemplate, devSpaceVDIQuery.Spec) {
				devSpaceVDIQuery.Spec = instance.Spec.DevSpaceVDITemplate
				err = r.Update(ctx, devSpaceVDIQuery)
				if err != nil {
					return err
				}
			}

			instance.Status.DevSpaceVDIStatus.Resource.Created = true
			reference.SetReference(&instance.Status.DevSpaceVDIStatus.Resource.Reference, devSpaceVDIQuery.TypeMeta, devSpaceVDIQuery.ObjectMeta)
			instance.Status.DevSpaceVDIStatus.Resource.Phase = string(devSpaceVDIQuery.Status.Phase)
			instance.Status.DevSpaceVDIStatus.Connection = devSpaceVDIQuery.Status.ServiceTCPStatus.URL

		} else {

			err := r.Delete(ctx, devSpaceVDIQuery)
			if err != nil {
				return err
			}

		}

	}

	return nil
}

func (r *DevSuiteReconciler) reconcileCheckDevSpaceIDE(ctx context.Context, instance *devv1alpha1.DevSuite) error {

	devSpaceIDEQuery := &devv1alpha1.DevSpaceIDE{}
	err := r.Get(ctx, *instance.GetDevSpaceIDEMetadata(), devSpaceIDEQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.DevSpaceIDEStatus = devv1alpha1.OwnedDevSpaceServiceStatus{}
		} else {
			return err
		}
	} else {

		if instance.Spec.IDEEnabled {

			if !reflect.DeepEqual(instance.Spec.DevSpaceIDETemplate, devSpaceIDEQuery.Spec) {
				devSpaceIDEQuery.Spec = instance.Spec.DevSpaceIDETemplate
				err = r.Update(ctx, devSpaceIDEQuery)
				if err != nil {
					return err
				}
			}

			instance.Status.DevSpaceIDEStatus.Resource.Created = true
			reference.SetReference(&instance.Status.DevSpaceIDEStatus.Resource.Reference, devSpaceIDEQuery.TypeMeta, devSpaceIDEQuery.ObjectMeta)
			instance.Status.DevSpaceIDEStatus.Resource.Phase = string(devSpaceIDEQuery.Status.Phase)
			instance.Status.DevSpaceIDEStatus.Connection = devSpaceIDEQuery.Status.ServiceStatus.URL

		} else {

			err := r.Delete(ctx, devSpaceIDEQuery)
			if err != nil {
				return err
			}

		}

	}

	return nil
}

func (r *DevSuiteReconciler) reconcileCheckDevSpaceJupyter(ctx context.Context, instance *devv1alpha1.DevSuite) error {

	devSpaceJupyterQuery := &devv1alpha1.DevSpaceJupyter{}
	err := r.Get(ctx, *instance.GetDevSpaceJupyterMetadata(), devSpaceJupyterQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.DevSpaceJupyterStatus = devv1alpha1.OwnedDevSpaceServiceStatus{}
		} else {
			return err
		}
	} else {

		if instance.Spec.IDEEnabled {

			if !reflect.DeepEqual(instance.Spec.DevSpaceJupyterTemplate, devSpaceJupyterQuery.Spec) {
				devSpaceJupyterQuery.Spec = instance.Spec.DevSpaceJupyterTemplate
				err = r.Update(ctx, devSpaceJupyterQuery)
				if err != nil {
					return err
				}
			}

			instance.Status.DevSpaceJupyterStatus.Resource.Created = true
			reference.SetReference(&instance.Status.DevSpaceJupyterStatus.Resource.Reference, devSpaceJupyterQuery.TypeMeta, devSpaceJupyterQuery.ObjectMeta)
			instance.Status.DevSpaceJupyterStatus.Resource.Phase = string(devSpaceJupyterQuery.Status.Phase)
			instance.Status.DevSpaceJupyterStatus.Connection = devSpaceJupyterQuery.Status.ServiceStatus.URL

		} else {

			err := r.Delete(ctx, devSpaceJupyterQuery)
			if err != nil {
				return err
			}

		}

	}

	return nil
}
