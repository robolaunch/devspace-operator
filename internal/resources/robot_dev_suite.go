package resources

import (
	"github.com/robolaunch/devspace-operator/internal"
	robotv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func GetDevSpaceVDI(robotDevSuite *robotv1alpha1.RobotDevSuite, devSpaceVDINamespacedName *types.NamespacedName) *robotv1alpha1.DevSpaceVDI {

	devSpaceVDI := robotv1alpha1.DevSpaceVDI{
		ObjectMeta: metav1.ObjectMeta{
			Name:      devSpaceVDINamespacedName.Name,
			Namespace: devSpaceVDINamespacedName.Namespace,
			Labels:    robotDevSuite.Labels,
		},
		Spec: robotDevSuite.Spec.DevSpaceVDITemplate,
	}

	return &devSpaceVDI
}

func GetDevSpaceIDE(robotDevSuite *robotv1alpha1.RobotDevSuite, devSpaceIDENamespacedName *types.NamespacedName) *robotv1alpha1.DevSpaceIDE {

	devSpaceIDE := robotv1alpha1.DevSpaceIDE{
		ObjectMeta: metav1.ObjectMeta{
			Name:      devSpaceIDENamespacedName.Name,
			Namespace: devSpaceIDENamespacedName.Namespace,
			Labels:    robotDevSuite.Labels,
		},
		Spec: robotDevSuite.Spec.DevSpaceIDETemplate,
	}

	if robotDevSuite.Spec.VDIEnabled {
		devSpaceIDE.Labels[internal.TARGET_VDI_LABEL_KEY] = robotDevSuite.GetDevSpaceVDIMetadata().Name
	}

	return &devSpaceIDE
}
