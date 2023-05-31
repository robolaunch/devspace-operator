package resources

import (
	"github.com/robolaunch/devspace-operator/internal"
	robotv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func GetDevSpaceVDI(devSuite *robotv1alpha1.DevSuite, devSpaceVDINamespacedName *types.NamespacedName) *robotv1alpha1.DevSpaceVDI {

	devSpaceVDI := robotv1alpha1.DevSpaceVDI{
		ObjectMeta: metav1.ObjectMeta{
			Name:      devSpaceVDINamespacedName.Name,
			Namespace: devSpaceVDINamespacedName.Namespace,
			Labels:    devSuite.Labels,
		},
		Spec: devSuite.Spec.DevSpaceVDITemplate,
	}

	return &devSpaceVDI
}

func GetDevSpaceIDE(devSuite *robotv1alpha1.DevSuite, devSpaceIDENamespacedName *types.NamespacedName) *robotv1alpha1.DevSpaceIDE {

	devSpaceIDE := robotv1alpha1.DevSpaceIDE{
		ObjectMeta: metav1.ObjectMeta{
			Name:      devSpaceIDENamespacedName.Name,
			Namespace: devSpaceIDENamespacedName.Namespace,
			Labels:    devSuite.Labels,
		},
		Spec: devSuite.Spec.DevSpaceIDETemplate,
	}

	if devSuite.Spec.VDIEnabled {
		devSpaceIDE.Labels[internal.TARGET_VDI_LABEL_KEY] = devSuite.GetDevSpaceVDIMetadata().Name
	}

	return &devSpaceIDE
}
