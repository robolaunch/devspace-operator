package v1alpha1

import (
	"github.com/robolaunch/devspace-operator/internal"
	"k8s.io/apimachinery/pkg/types"
)

// ********************************
// WorkspaceManager helpers
// ********************************

func (workspacemanager *WorkspaceManager) GetClonerJobMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      workspacemanager.Name + internal.JOB_CLONER_POSTFIX,
		Namespace: workspacemanager.Namespace,
	}
}

func (workspacemanager *WorkspaceManager) GetCleanupJobMetadata() *types.NamespacedName {
	return &types.NamespacedName{
		Name:      workspacemanager.Name + internal.JOB_CLEANUP_POSTFIX,
		Namespace: workspacemanager.Namespace,
	}
}
