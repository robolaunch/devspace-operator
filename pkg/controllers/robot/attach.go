package robot

import (
	"context"
	"sort"

	"github.com/robolaunch/robot-operator/internal"
	robotv1alpha1 "github.com/robolaunch/robot-operator/pkg/api/roboscale.io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *RobotReconciler) reconcileAttachDevObject(ctx context.Context, instance *robotv1alpha1.Robot) error {

	// Get attached dev objects for this robot
	requirements := []labels.Requirement{}
	newReq, err := labels.NewRequirement(internal.TARGET_ROBOT_LABEL_KEY, selection.In, []string{instance.Name})
	if err != nil {
		return err
	}
	requirements = append(requirements, *newReq)

	robotSelector := labels.NewSelector().Add(requirements...)

	robotDevSuiteList := robotv1alpha1.RobotDevSuiteList{}
	err = r.List(ctx, &robotDevSuiteList, &client.ListOptions{Namespace: instance.Namespace, LabelSelector: robotSelector})
	if err != nil {
		return err
	}

	if len(robotDevSuiteList.Items) == 0 {
		instance.Status.AttachedDevObjects = []robotv1alpha1.AttachedDevObject{}
		return nil
	}

	// Sort attached dev objects for this robot according to their creation timestamps
	sort.SliceStable(robotDevSuiteList.Items[:], func(i, j int) bool {
		return robotDevSuiteList.Items[i].CreationTimestamp.String() < robotDevSuiteList.Items[j].CreationTimestamp.String()
	})

	instance.Status.AttachedDevObjects = []robotv1alpha1.AttachedDevObject{}

	for _, rds := range robotDevSuiteList.Items {
		instance.Status.AttachedDevObjects = append(instance.Status.AttachedDevObjects, robotv1alpha1.AttachedDevObject{
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
