package dev_suite

import (
	"context"

	"github.com/robolaunch/devspace-operator/internal/label"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
)

func (r *DevSuiteReconciler) reconcileGetInstance(ctx context.Context, meta types.NamespacedName) (*devv1alpha1.DevSuite, error) {
	instance := &devv1alpha1.DevSuite{}
	err := r.Get(ctx, meta, instance)
	if err != nil {
		return &devv1alpha1.DevSuite{}, err
	}

	return instance, nil
}

func (r *DevSuiteReconciler) reconcileUpdateInstanceStatus(ctx context.Context, instance *devv1alpha1.DevSuite) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		instanceLV := &devv1alpha1.DevSuite{}
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

func (r *DevSuiteReconciler) reconcileGetTargetDevSpace(ctx context.Context, instance *devv1alpha1.DevSuite) (*devv1alpha1.DevSpace, error) {
	devspace := &devv1alpha1.DevSpace{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: instance.Namespace,
		Name:      label.GetTargetDevSpace(instance),
	}, devspace)
	if err != nil {
		return nil, err
	}

	return devspace, nil
}

func (r *DevSuiteReconciler) reconcileCheckTargetDevSpace(ctx context.Context, instance *devv1alpha1.DevSuite) error {

	if label.GetDevSuiteOwned(instance) == "true" {
		instance.Status.Active = true
	} else {
		devspace, err := r.reconcileGetTargetDevSpace(ctx, instance)
		if err != nil {
			return err
		}

		isActive := false
		for _, rds := range devspace.Status.AttachedDevObjects {
			if rds.Reference.Kind == instance.Kind && rds.Reference.Name == instance.Name {
				isActive = true
				break
			}
		}

		instance.Status.Active = isActive
	}

	return nil
}
