package v1alpha1

import (
	"github.com/robolaunch/devspace-operator/internal"
	"k8s.io/apimachinery/pkg/types"
)

// ********************************
// RobotDevSuite helpers
// ********************************

func (robotDevSuite *RobotDevSuite) GetDevSpaceVDIMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: robotDevSuite.Namespace,
		Name:      robotDevSuite.Name + internal.DEVSPACE_VDI_POSTFIX,
	}
}

func (robotDevSuite *RobotDevSuite) GetDevSpaceIDEMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: robotDevSuite.Namespace,
		Name:      robotDevSuite.Name + internal.DEVSPACE_IDE_POSTFIX,
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
// DevSpaceVDI helpers
// ********************************

func (devspacevdi *DevSpaceVDI) GetDevSpaceVDIPVCMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: devspacevdi.Namespace,
		Name:      devspacevdi.Name + internal.PVC_VDI_POSTFIX,
	}
}

func (devspacevdi *DevSpaceVDI) GetDevSpaceVDIPodMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: devspacevdi.Namespace,
		Name:      devspacevdi.Name + internal.POD_VDI_POSTFIX,
	}
}

func (devspacevdi *DevSpaceVDI) GetDevSpaceVDIServiceTCPMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: devspacevdi.Namespace,
		Name:      devspacevdi.Name + internal.SVC_TCP_VDI_POSTFIX,
	}
}

func (devspacevdi *DevSpaceVDI) GetDevSpaceVDIServiceUDPMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: devspacevdi.Namespace,
		Name:      devspacevdi.Name + internal.SVC_UDP_VDI_POSTFIX,
	}
}

func (devspacevdi *DevSpaceVDI) GetDevSpaceVDIIngressMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: devspacevdi.Namespace,
		Name:      devspacevdi.Name + internal.INGRESS_VDI_POSTFIX,
	}
}
