package robot

import (
	"context"

	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *DevspaceReconciler) reconcileCheckDeletion(ctx context.Context, instance *devv1alpha1.Devspace) error {

	robotFinalizer := "dev.roboscale.io/finalizer"

	if instance.DeletionTimestamp.IsZero() {

		if !controllerutil.ContainsFinalizer(instance, robotFinalizer) {
			controllerutil.AddFinalizer(instance, robotFinalizer)
			if err := r.Update(ctx, instance); err != nil {
				return err
			}
		}

	} else {

		if controllerutil.ContainsFinalizer(instance, robotFinalizer) {

			err := r.waitForLoaderJobDeletion(ctx, instance)
			if err != nil {
				return err
			}

			err = r.waitForPersistentVolumeClaimDeletion(ctx, instance, instance.GetPVCVarMetadata())
			if err != nil {
				return err
			}

			err = r.waitForPersistentVolumeClaimDeletion(ctx, instance, instance.GetPVCEtcMetadata())
			if err != nil {
				return err
			}

			err = r.waitForPersistentVolumeClaimDeletion(ctx, instance, instance.GetPVCOptMetadata())
			if err != nil {
				return err
			}

			err = r.waitForPersistentVolumeClaimDeletion(ctx, instance, instance.GetPVCUsrMetadata())
			if err != nil {
				return err
			}

			err = r.waitForPersistentVolumeClaimDeletion(ctx, instance, instance.GetPVCWorkspaceMetadata())
			if err != nil {
				return err
			}

			controllerutil.RemoveFinalizer(instance, robotFinalizer)
			if err := r.Update(ctx, instance); err != nil {
				return err
			}
		}

		return errors.NewNotFound(schema.GroupResource{
			Group:    instance.GetObjectKind().GroupVersionKind().Group,
			Resource: instance.GetObjectKind().GroupVersionKind().Kind,
		}, instance.Name)
	}

	return nil
}

func (r *DevspaceReconciler) waitForLoaderJobDeletion(ctx context.Context, instance *devv1alpha1.Devspace) error {

	instance.Status.Phase = devv1alpha1.DevspacePhaseDeletingLoaderJob
	err := r.reconcileUpdateInstanceStatus(ctx, instance)
	if err != nil {
		return err
	}

	loaderJobQuery := &batchv1.Job{}
	err = r.Get(ctx, *instance.GetLoaderJobMetadata(), loaderJobQuery)
	if err != nil && errors.IsNotFound(err) {
		return nil
	} else if err != nil {
		return err
	} else {
		logger.Info("FINALIZER: Loader job is being deleted.")
		propagationPolicy := metav1.DeletePropagationBackground
		err := r.Delete(ctx, loaderJobQuery, &client.DeleteOptions{
			PropagationPolicy: &propagationPolicy,
		})
		if err != nil {
			return err
		}
	}

	loaderJobQuery = &batchv1.Job{}
	err = r.Get(ctx, *instance.GetLoaderJobMetadata(), loaderJobQuery)
	if err != nil && errors.IsNotFound(err) {
		return nil
	} else if err != nil {
		return err
	} else {

		resourceInterface := r.DynamicClient.Resource(schema.GroupVersionResource{
			Group:    loaderJobQuery.GroupVersionKind().Group,
			Version:  loaderJobQuery.GroupVersionKind().Version,
			Resource: "jobs",
		})
		jobWatcher, err := resourceInterface.Watch(ctx, metav1.ListOptions{
			FieldSelector: "metadata.name=" + instance.GetLoaderJobMetadata().Name,
		})
		if err != nil {
			return err
		}

		defer jobWatcher.Stop()

		jobDeleted := false
		for {
			if !jobDeleted {
				select {
				case event := <-jobWatcher.ResultChan():

					if event.Type == watch.Deleted {
						logger.Info("FINALIZER: Loader job is deleted gracefully.")
						jobDeleted = true
					}
				}
			} else {
				break
			}

		}
	}
	return nil
}

func (r *DevspaceReconciler) waitForPersistentVolumeClaimDeletion(ctx context.Context, instance *devv1alpha1.Devspace, pvcNamespacedName *types.NamespacedName) error {

	pvcQuery := &corev1.PersistentVolumeClaim{}
	err := r.Get(ctx, *pvcNamespacedName, pvcQuery)
	if err != nil && errors.IsNotFound(err) {
		return nil
	} else if err != nil {
		return err
	} else {
		logger.Info("FINALIZER: PVC " + pvcNamespacedName.Name + " is being deleted.")
		err := r.Delete(ctx, pvcQuery)
		if err != nil {
			return err
		}

		instance.Status.Phase = devv1alpha1.DevspacePhaseDeletingVolumes
		err = r.reconcileUpdateInstanceStatus(ctx, instance)
		if err != nil {
			return err
		}

		resourceInterface := r.DynamicClient.Resource(schema.GroupVersionResource{
			Group:    pvcQuery.GroupVersionKind().Group,
			Version:  pvcQuery.GroupVersionKind().Version,
			Resource: "persistentvolumeclaims",
		})
		pvcWatcher, err := resourceInterface.Watch(ctx, metav1.ListOptions{
			FieldSelector: "metadata.name=" + pvcNamespacedName.Name,
		})
		if err != nil {
			return err
		}

		defer pvcWatcher.Stop()

		pvcDeleted := false
		for {
			if !pvcDeleted {
				select {
				case event := <-pvcWatcher.ResultChan():

					if event.Type == watch.Deleted {
						logger.Info("FINALIZER: PVC " + pvcNamespacedName.Name + " is deleted gracefully.")
						pvcDeleted = true
					}
				}
			} else {
				break
			}

		}
	}
	return nil
}
