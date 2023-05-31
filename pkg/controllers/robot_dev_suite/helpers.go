package robot_dev_suite

import (
	"context"

	"github.com/robolaunch/devspace-operator/internal/label"
	robotv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
)

func (r *RobotDevSuiteReconciler) reconcileGetInstance(ctx context.Context, meta types.NamespacedName) (*robotv1alpha1.RobotDevSuite, error) {
	instance := &robotv1alpha1.RobotDevSuite{}
	err := r.Get(ctx, meta, instance)
	if err != nil {
		return &robotv1alpha1.RobotDevSuite{}, err
	}

	return instance, nil
}

func (r *RobotDevSuiteReconciler) reconcileUpdateInstanceStatus(ctx context.Context, instance *robotv1alpha1.RobotDevSuite) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		instanceLV := &robotv1alpha1.RobotDevSuite{}
		err := r.Get(ctx, types.NamespacedName{
			Name:      instance.Name,
			Namespace: instance.Namespace,
		}, instanceLV)

		if err == nil {
			instance.ResourceVersion = instanceLV.ResourceVersion
		}

		err1 := r.Status().Update(ctx, instance)
		return err1
	})
}

func (r *RobotDevSuiteReconciler) reconcileGetTargetRobot(ctx context.Context, instance *robotv1alpha1.RobotDevSuite) (*robotv1alpha1.Robot, error) {
	robot := &robotv1alpha1.Robot{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: instance.Namespace,
		Name:      label.GetTargetRobot(instance),
	}, robot)
	if err != nil {
		return nil, err
	}

	return robot, nil
}

func (r *RobotDevSuiteReconciler) reconcileCheckTargetRobot(ctx context.Context, instance *robotv1alpha1.RobotDevSuite) error {

	if label.GetDevSuiteOwned(instance) == "true" {
		instance.Status.Active = true
	} else {
		robot, err := r.reconcileGetTargetRobot(ctx, instance)
		if err != nil {
			return err
		}

		isActive := false
		for _, rds := range robot.Status.AttachedDevObjects {
			if rds.Reference.Kind == instance.Kind && rds.Reference.Name == instance.Name {
				isActive = true
				break
			}
		}

		instance.Status.Active = isActive
	}

	return nil
}
