package robot_dev_suite

import (
	"context"
	"time"

	robotv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *DevSuiteReconciler) reconcileDeleteDevSpaceVDI(ctx context.Context, instance *robotv1alpha1.DevSuite) error {

	devSpaceVDIQuery := &robotv1alpha1.DevSpaceVDI{}
	err := r.Get(ctx, *instance.GetDevSpaceVDIMetadata(), devSpaceVDIQuery)
	if err != nil {
		if errors.IsNotFound(err) {
			instance.Status.DevSpaceVDIStatus = robotv1alpha1.OwnedRobotServiceStatus{}
		} else {
			return err
		}
	} else {

		propagationPolicy := v1.DeletePropagationForeground
		err := r.Delete(ctx, devSpaceVDIQuery, &client.DeleteOptions{
			PropagationPolicy: &propagationPolicy,
		})
		if err != nil {
			return err
		}

		// watch until it's deleted
		deleted := false
		for !deleted {
			devSpaceVDIQuery := &robotv1alpha1.DevSpaceVDI{}
			err := r.Get(ctx, *instance.GetDevSpaceVDIMetadata(), devSpaceVDIQuery)
			if err != nil && errors.IsNotFound(err) {
				deleted = true
			}
			time.Sleep(time.Second * 1)
		}

		instance.Status.DevSpaceVDIStatus = robotv1alpha1.OwnedRobotServiceStatus{}
	}

	return nil
}

func (r *DevSuiteReconciler) reconcileDeleteDevSpaceIDE(ctx context.Context, instance *robotv1alpha1.DevSuite) error {

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
