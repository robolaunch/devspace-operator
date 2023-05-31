package robot_dev_suite

import (
	"context"
	"reflect"

	"github.com/robolaunch/devspace-operator/internal/reference"
	robotv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
)

func (r *DevSuiteReconciler) reconcileCheckDevSpaceVDI(ctx context.Context, instance *robotv1alpha1.DevSuite) error {

	devSpaceVDIQuery := &robotv1alpha1.DevSpaceVDI{}
	err := r.Get(ctx, *instance.GetDevSpaceVDIMetadata(), devSpaceVDIQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.DevSpaceVDIStatus = robotv1alpha1.OwnedRobotServiceStatus{}
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

func (r *DevSuiteReconciler) reconcileCheckDevSpaceIDE(ctx context.Context, instance *robotv1alpha1.DevSuite) error {

	devSpaceIDEQuery := &robotv1alpha1.DevSpaceIDE{}
	err := r.Get(ctx, *instance.GetDevSpaceIDEMetadata(), devSpaceIDEQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.DevSpaceIDEStatus = robotv1alpha1.OwnedRobotServiceStatus{}
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
