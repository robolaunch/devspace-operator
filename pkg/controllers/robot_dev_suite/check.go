package robot_dev_suite

import (
	"context"
	"reflect"

	"github.com/robolaunch/robot-operator/internal/reference"
	robotv1alpha1 "github.com/robolaunch/robot-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
)

func (r *RobotDevSuiteReconciler) reconcileCheckRobotVDI(ctx context.Context, instance *robotv1alpha1.RobotDevSuite) error {

	robotVDIQuery := &robotv1alpha1.RobotVDI{}
	err := r.Get(ctx, *instance.GetRobotVDIMetadata(), robotVDIQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.RobotVDIStatus = robotv1alpha1.OwnedRobotServiceStatus{}
		} else {
			return err
		}
	} else {

		if instance.Spec.VDIEnabled {

			if !reflect.DeepEqual(instance.Spec.RobotVDITemplate, robotVDIQuery.Spec) {
				robotVDIQuery.Spec = instance.Spec.RobotVDITemplate
				err = r.Update(ctx, robotVDIQuery)
				if err != nil {
					return err
				}
			}

			instance.Status.RobotVDIStatus.Resource.Created = true
			reference.SetReference(&instance.Status.RobotVDIStatus.Resource.Reference, robotVDIQuery.TypeMeta, robotVDIQuery.ObjectMeta)
			instance.Status.RobotVDIStatus.Resource.Phase = string(robotVDIQuery.Status.Phase)
			instance.Status.RobotVDIStatus.Connection = robotVDIQuery.Status.ServiceTCPStatus.URL

		} else {

			err := r.Delete(ctx, robotVDIQuery)
			if err != nil {
				return err
			}

		}

	}

	return nil
}

func (r *RobotDevSuiteReconciler) reconcileCheckDevSpaceIDE(ctx context.Context, instance *robotv1alpha1.RobotDevSuite) error {

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
