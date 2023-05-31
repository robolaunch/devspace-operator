package v1alpha1

import (
	"errors"

	"github.com/robolaunch/devspace-operator/internal"
	"k8s.io/apimachinery/pkg/types"
)

// ********************************
// Devspace helpers
// ********************************

func (robot *Devspace) GetPVCVarMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      robot.Name + internal.PVC_VAR_POSTFIX,
		Namespace: robot.Namespace,
	}
}

func (robot *Devspace) GetPVCOptMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      robot.Name + internal.PVC_OPT_POSTFIX,
		Namespace: robot.Namespace,
	}
}

func (robot *Devspace) GetPVCUsrMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      robot.Name + internal.PVC_USR_POSTFIX,
		Namespace: robot.Namespace,
	}
}

func (robot *Devspace) GetPVCEtcMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      robot.Name + internal.PVC_ETC_POSTFIX,
		Namespace: robot.Namespace,
	}
}

func (robot *Devspace) GetPVCWorkspaceMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      robot.Name + internal.PVC_WORKSPACE_POSTFIX,
		Namespace: robot.Namespace,
	}
}

func (robot *Devspace) GetLoaderJobMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      robot.Name + internal.JOB_LOADER_POSTFIX,
		Namespace: robot.Namespace,
	}
}

func (robot *Devspace) GetDevSuiteMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      robot.Name + internal.ROBOT_DEV_SUITE_POSTFIX,
		Namespace: robot.Namespace,
	}
}

func (robot *Devspace) GetWorkspaceManagerMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      robot.Name + internal.WORKSPACE_MANAGER_POSTFIX,
		Namespace: robot.Namespace,
	}
}

func (robot *Devspace) GetWorkspaceByName(name string) (Workspace, error) {

	for _, ws := range robot.Spec.WorkspaceManagerTemplate.Workspaces {
		if ws.Name == name {
			return ws, nil
		}
	}

	return Workspace{}, errors.New("workspace not found")
}
