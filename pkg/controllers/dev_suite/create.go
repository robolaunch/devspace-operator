package dev_suite

import (
	"context"

	"github.com/robolaunch/devspace-operator/internal/resources"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *DevSuiteReconciler) reconcileCreateDevSpaceVDI(ctx context.Context, instance *devv1alpha1.DevSuite) error {

	devSpaceVDI := resources.GetDevSpaceVDI(instance, instance.GetDevSpaceVDIMetadata())

	err := ctrl.SetControllerReference(instance, devSpaceVDI, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, devSpaceVDI)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: DevSpace VDI is created.")

	return nil
}

func (r *DevSuiteReconciler) reconcileCreateDevSpaceIDE(ctx context.Context, instance *devv1alpha1.DevSuite) error {

	devSpaceIDE := resources.GetDevSpaceIDE(instance, instance.GetDevSpaceIDEMetadata())

	err := ctrl.SetControllerReference(instance, devSpaceIDE, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, devSpaceIDE)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: DevSpace IDE is created.")

	return nil
}

func (r *DevSuiteReconciler) reconcileCreateDevSpaceJupyter(ctx context.Context, instance *devv1alpha1.DevSuite) error {

	devSpaceJupyter := resources.GetDevSpaceJupyter(instance, instance.GetDevSpaceJupyterMetadata())

	err := ctrl.SetControllerReference(instance, devSpaceJupyter, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, devSpaceJupyter)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info("STATUS: DevSpace Jupyter is created.")

	return nil
}
