package robot_dev_suite

import (
	"context"
	"time"

	robotv1alpha1 "github.com/robolaunch/robot-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *RobotDevSuiteReconciler) reconcileDeleteRobotVDI(ctx context.Context, instance *robotv1alpha1.RobotDevSuite) error {

	robotVDIQuery := &robotv1alpha1.RobotVDI{}
	err := r.Get(ctx, *instance.GetRobotVDIMetadata(), robotVDIQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.RobotVDIStatus = robotv1alpha1.OwnedRobotServiceStatus{}
		} else {
			return err
		}
	} else {

		propagationPolicy := v1.DeletePropagationForeground
		err := r.Delete(ctx, robotVDIQuery, &client.DeleteOptions{
			PropagationPolicy: &propagationPolicy,
		})
		if err != nil {
			return err
		}

		// watch until it's deleted
		deleted := false
		for !deleted {
			robotVDIQuery := &robotv1alpha1.RobotVDI{}
			err := r.Get(ctx, *instance.GetRobotVDIMetadata(), robotVDIQuery)
			if err != nil && errors.IsNotFound(err) {
				deleted = true
			}
			time.Sleep(time.Second * 1)
		}

		instance.Status.RobotVDIStatus = robotv1alpha1.OwnedRobotServiceStatus{}
	}

	return nil
}

func (r *RobotDevSuiteReconciler) reconcileDeleteDevSpaceIDE(ctx context.Context, instance *robotv1alpha1.RobotDevSuite) error {

	devSpaceIDEQuery := &robotv1alpha1.DevSpaceIDE{}
	err := r.Get(ctx, *instance.GetDevSpaceIDEMetadata(), devSpaceIDEQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.DevSpaceIDEStatus = robotv1alpha1.OwnedRobotServiceStatus{}
		} else {
			return err
		}
	} else {

		propagationPolicy := v1.DeletePropagationForeground
		err := r.Delete(ctx, devSpaceIDEQuery, &client.DeleteOptions{
			PropagationPolicy: &propagationPolicy,
		})
		if err != nil {
			return err
		}

		// watch until it's deleted
		deleted := false
		for !deleted {
			devSpaceIDEQuery := &robotv1alpha1.DevSpaceIDE{}
			err := r.Get(ctx, *instance.GetDevSpaceIDEMetadata(), devSpaceIDEQuery)
			if err != nil && errors.IsNotFound(err) {
				deleted = true
			}
			time.Sleep(time.Second * 1)
		}

		instance.Status.DevSpaceIDEStatus = robotv1alpha1.OwnedRobotServiceStatus{}
	}

	return nil
}
