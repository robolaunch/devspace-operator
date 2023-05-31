package label

import (
	"github.com/robolaunch/devspace-operator/internal"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetTargetDevSpace(obj metav1.Object) string {
	labels := obj.GetLabels()

	if targetDevSpace, ok := labels[internal.TARGET_DEVSPACE_LABEL_KEY]; ok {
		return targetDevSpace
	}

	return ""
}

func GetTargetDevSpaceVDI(obj metav1.Object) string {
	labels := obj.GetLabels()

	if targetDevSpaceVDI, ok := labels[internal.TARGET_VDI_LABEL_KEY]; ok {
		return targetDevSpaceVDI
	}

	return ""
}

func GetDevSuiteOwned(obj metav1.Object) string {
	labels := obj.GetLabels()

	if devSuiteOwned, ok := labels[internal.DEVSPACE_DEV_SUITE_OWNED]; ok {
		return devSuiteOwned
	}

	return ""
}
