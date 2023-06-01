package resources

import (
	"github.com/robolaunch/devspace-operator/internal"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func GetDevSpaceVDI(devSuite *devv1alpha1.DevSuite, devSpaceVDINamespacedName *types.NamespacedName) *devv1alpha1.DevSpaceVDI {

	devSpaceVDI := devv1alpha1.DevSpaceVDI{
		ObjectMeta: metav1.ObjectMeta{
			Name:      devSpaceVDINamespacedName.Name,
			Namespace: devSpaceVDINamespacedName.Namespace,
			Labels:    devSuite.Labels,
		},
		Spec: devSuite.Spec.DevSpaceVDITemplate,
	}

	return &devSpaceVDI
}

func GetDevSpaceIDE(devSuite *devv1alpha1.DevSuite, devSpaceIDENamespacedName *types.NamespacedName) *devv1alpha1.DevSpaceIDE {

	devSpaceIDE := devv1alpha1.DevSpaceIDE{
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
