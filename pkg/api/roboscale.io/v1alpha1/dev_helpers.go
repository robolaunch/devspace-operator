package v1alpha1

import (
	"github.com/robolaunch/robot-operator/internal"
	"k8s.io/apimachinery/pkg/types"
)

// ********************************
// RobotDevSuite helpers
// ********************************

func (robotDevSuite *RobotDevSuite) GetRobotVDIMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: robotDevSuite.Namespace,
		Name:      robotDevSuite.Name + internal.ROBOT_VDI_POSTFIX,
	}
}

func (robotDevSuite *RobotDevSuite) GetDevSpaceIDEMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: robotDevSuite.Namespace,
		Name:      robotDevSuite.Name + internal.ROBOT_IDE_POSTFIX,
	}
}

// ********************************
// DevSpaceIDE helpers
// ********************************

func (devspaceide *DevSpaceIDE) GetDevSpaceIDEPodMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: devspaceide.Namespace,
		Name:      devspaceide.Name + internal.POD_IDE_POSTFIX,
	}
}

func (devspaceide *DevSpaceIDE) GetDevSpaceIDEServiceMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: devspaceide.Namespace,
		Name:      devspaceide.Name + internal.SVC_IDE_POSTFIX,
	}
}

func (devspaceide *DevSpaceIDE) GetDevSpaceIDEIngressMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: devspaceide.Namespace,
		Name:      devspaceide.Name + internal.INGRESS_IDE_POSTFIX,
	}
}

// ********************************
// RobotVDI helpers
// ********************************

func (robotvdi *RobotVDI) GetRobotVDIPVCMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: robotvdi.Namespace,
		Name:      robotvdi.Name + internal.PVC_VDI_POSTFIX,
	}
}

func (robotvdi *RobotVDI) GetRobotVDIPodMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: robotvdi.Namespace,
		Name:      robotvdi.Name + internal.POD_VDI_POSTFIX,
	}
}

func (robotvdi *RobotVDI) GetRobotVDIServiceTCPMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: robotvdi.Namespace,
		Name:      robotvdi.Name + internal.SVC_TCP_VDI_POSTFIX,
	}
}

func (robotvdi *RobotVDI) GetRobotVDIServiceUDPMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: robotvdi.Namespace,
		Name:      robotvdi.Name + internal.SVC_UDP_VDI_POSTFIX,
	}
}

func (robotvdi *RobotVDI) GetRobotVDIIngressMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: robotvdi.Namespace,
		Name:      robotvdi.Name + internal.INGRESS_VDI_POSTFIX,
	}
}
