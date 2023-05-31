package resources

import (
	"github.com/robolaunch/devspace-operator/internal"
	robotv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func GetRobotVDI(robotDevSuite *robotv1alpha1.RobotDevSuite, robotVDINamespacedName *types.NamespacedName) *robotv1alpha1.RobotVDI {

	robotVDI := robotv1alpha1.RobotVDI{
		ObjectMeta: metav1.ObjectMeta{
			Name:      robotVDINamespacedName.Name,
			Namespace: robotVDINamespacedName.Namespace,
			Labels:    robotDevSuite.Labels,
		},
		Spec: robotDevSuite.Spec.RobotVDITemplate,
	}

	return &robotVDI
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
		devSpaceIDE.Labels[internal.TARGET_VDI_LABEL_KEY] = robotDevSuite.GetRobotVDIMetadata().Name
	}

	return &devSpaceIDE
}
