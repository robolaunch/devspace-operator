package resources

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/robolaunch/devspace-operator/internal"
	"github.com/robolaunch/devspace-operator/internal/configure"
	"github.com/robolaunch/devspace-operator/internal/label"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func GetClonerJob(workspaceManager *devv1alpha1.WorkspaceManager, jobNamespacedName *types.NamespacedName, devspace *devv1alpha1.DevSpace) *batchv1.Job {

	var clonerCmdBuilder strings.Builder
	for wsKey, ws := range workspaceManager.Spec.Workspaces {

		var cmdBuilder strings.Builder
		cmdBuilder.WriteString("mkdir -p " + filepath.Join(workspaceManager.Spec.WorkspacesPath, ws.Name, "src") + " && ")
		cmdBuilder.WriteString("cd " + filepath.Join(workspaceManager.Spec.WorkspacesPath, ws.Name, "src") + " && ")
		cmdBuilder.WriteString(GetCloneCommand(workspaceManager.Spec.Workspaces, wsKey))
		clonerCmdBuilder.WriteString(cmdBuilder.String())

	}

	clonerCmdBuilder.WriteString("echo \"DONE\"")

	clonerContainer := corev1.Container{
		Name:    "cloner",
		Image:   "ubuntu:focal",
		Command: internal.Bash(clonerCmdBuilder.String()),
		VolumeMounts: []corev1.VolumeMount{
			configure.GetVolumeMount("", configure.GetVolumeVar(devspace)),
			configure.GetVolumeMount("", configure.GetVolumeUsr(devspace)),
			configure.GetVolumeMount("", configure.GetVolumeOpt(devspace)),
			configure.GetVolumeMount("", configure.GetVolumeEtc(devspace)),
			configure.GetVolumeMount(workspaceManager.Spec.WorkspacesPath, configure.GetVolumeWorkspace(devspace)),
		},
	}

	podSpec := &corev1.PodSpec{
		Containers: []corev1.Container{
			clonerContainer,
		},
		Volumes: []corev1.Volume{
			configure.GetVolumeVar(devspace),
			configure.GetVolumeUsr(devspace),
			configure.GetVolumeOpt(devspace),
			configure.GetVolumeEtc(devspace),
			configure.GetVolumeWorkspace(devspace),
		},
	}

	podSpec.RestartPolicy = corev1.RestartPolicyNever
	podSpec.NodeSelector = label.GetTenancyMap(devspace)

	job := batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      workspaceManager.GetClonerJobMetadata().Name,
			Namespace: workspaceManager.GetClonerJobMetadata().Namespace,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: *podSpec,
			},
		},
	}

	return &job
}

func GetCleanupJob(workspaceManager *devv1alpha1.WorkspaceManager, jobNamespacedName *types.NamespacedName, devspace *devv1alpha1.DevSpace) *batchv1.Job {

	var cmdBuilder strings.Builder
	cmdBuilder.WriteString("cd " + workspaceManager.Spec.WorkspacesPath + " && ")
	cmdBuilder.WriteString("GLOBIGNORE=old &&")
	cmdBuilder.WriteString("mkdir -p old/backup-" + strconv.Itoa(workspaceManager.Status.Version) + " &&")
	cmdBuilder.WriteString("mv * old/backup-" + strconv.Itoa(workspaceManager.Status.Version) + " || true &&")
	cmdBuilder.WriteString("rm -rf *")

	cleanupContainer := corev1.Container{
		Name:    "cleanup",
		Image:   "ubuntu:focal",
		Command: internal.Bash(cmdBuilder.String()),
		VolumeMounts: []corev1.VolumeMount{
			configure.GetVolumeMount("", configure.GetVolumeVar(devspace)),
			configure.GetVolumeMount("", configure.GetVolumeUsr(devspace)),
			configure.GetVolumeMount("", configure.GetVolumeOpt(devspace)),
			configure.GetVolumeMount("", configure.GetVolumeEtc(devspace)),
			configure.GetVolumeMount(workspaceManager.Spec.WorkspacesPath, configure.GetVolumeWorkspace(devspace)),
		},
	}

	podSpec := &corev1.PodSpec{
		Containers: []corev1.Container{
			cleanupContainer,
		},
		Volumes: []corev1.Volume{
			configure.GetVolumeVar(devspace),
			configure.GetVolumeUsr(devspace),
			configure.GetVolumeOpt(devspace),
			configure.GetVolumeEtc(devspace),
			configure.GetVolumeWorkspace(devspace),
		},
	}

	podSpec.RestartPolicy = corev1.RestartPolicyNever
	podSpec.NodeSelector = label.GetTenancyMap(devspace)

	job := batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      workspaceManager.GetCleanupJobMetadata().Name,
			Namespace: workspaceManager.GetCleanupJobMetadata().Namespace,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: *podSpec,
			},
		},
	}

	return &job
}

func GetCloneCommand(workspaces []devv1alpha1.Workspace, wsKey int) string {

	var cmdBuilder strings.Builder
	for key, repo := range workspaces[wsKey].Repositories {
		cmdBuilder.WriteString("git clone " + repo.URL + " -b " + repo.Branch + " " + key + " &&")
	}
	return cmdBuilder.String()
}
