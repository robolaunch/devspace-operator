package v1alpha1

import (
	"errors"

	"github.com/robolaunch/devspace-operator/internal"
	"k8s.io/apimachinery/pkg/types"
)

// ********************************
// DevSpace helpers
// ********************************

func (devspace *DevSpace) GetPVCVarMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      devspace.Name + internal.PVC_VAR_POSTFIX,
		Namespace: devspace.Namespace,
	}
}

func (devspace *DevSpace) GetPVCOptMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      devspace.Name + internal.PVC_OPT_POSTFIX,
		Namespace: devspace.Namespace,
	}
}

func (devspace *DevSpace) GetPVCUsrMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      devspace.Name + internal.PVC_USR_POSTFIX,
		Namespace: devspace.Namespace,
	}
}

func (devspace *DevSpace) GetPVCEtcMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      devspace.Name + internal.PVC_ETC_POSTFIX,
		Namespace: devspace.Namespace,
	}
}

func (devspace *DevSpace) GetPVCWorkspaceMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      devspace.Name + internal.PVC_WORKSPACE_POSTFIX,
		Namespace: devspace.Namespace,
	}
}

func (devspace *DevSpace) GetLoaderJobMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      devspace.Name + internal.JOB_LOADER_POSTFIX,
		Namespace: devspace.Namespace,
	}
}

func (devspace *DevSpace) GetDevSpaceVDIMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: devspace.Namespace,
		Name:      devspace.Name + internal.DEVSPACE_VDI_POSTFIX,
	}
}

func (devspace *DevSpace) GetDevSpaceIDEMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: devspace.Namespace,
		Name:      devspace.Name + internal.DEVSPACE_IDE_POSTFIX,
	}
}

func (devspace *DevSpace) GetWorkspaceManagerMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      devspace.Name + internal.WORKSPACE_MANAGER_POSTFIX,
		Namespace: devspace.Namespace,
	}
}

func (devspace *DevSpace) GetWorkspaceByName(name string) (Workspace, error) {

	for _, ws := range devspace.Spec.WorkspaceManagerTemplate.Workspaces {
		if ws.Name == name {
			return ws, nil
		}
	}

	return Workspace{}, errors.New("workspace not found")
}
