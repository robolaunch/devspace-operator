package devspace

import (
	"context"
	"sort"

	"github.com/robolaunch/devspace-operator/internal"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *DevSpaceReconciler) reconcileAttachDevObject(ctx context.Context, instance *devv1alpha1.DevSpace) error {

	// Get attached dev objects for this devspace
	requirements := []labels.Requirement{}
	newReq, err := labels.NewRequirement(internal.TARGET_DEVSPACE_LABEL_KEY, selection.In, []string{instance.Name})
	if err != nil {
		return err
	}
	requirements = append(requirements, *newReq)

	devspaceSelector := labels.NewSelector().Add(requirements...)

	devSuiteList := devv1alpha1.DevSuiteList{}
	err = r.List(ctx, &devSuiteList, &client.ListOptions{Namespace: instance.Namespace, LabelSelector: devspaceSelector})
	if err != nil {
		return err
	}

	if len(devSuiteList.Items) == 0 {
		instance.Status.AttachedDevObjects = []devv1alpha1.AttachedDevObject{}
		return nil
	}

	// Sort attached dev objects for this devspace according to their creation timestamps
	sort.SliceStable(devSuiteList.Items[:], func(i, j int) bool {
		return devSuiteList.Items[i].CreationTimestamp.String() < devSuiteList.Items[j].CreationTimestamp.String()
	})

	instance.Status.AttachedDevObjects = []devv1alpha1.AttachedDevObject{}

	for _, rds := range devSuiteList.Items {
		instance.Status.AttachedDevObjects = append(instance.Status.AttachedDevObjects, devv1alpha1.AttachedDevObject{
			Reference: corev1.ObjectReference{
				Kind:            rds.Kind,
				Namespace:       rds.Namespace,
				Name:            rds.Name,
				UID:             rds.UID,
				APIVersion:      rds.APIVersion,
				ResourceVersion: rds.ResourceVersion,
			},
		})
	}

	return nil
}
