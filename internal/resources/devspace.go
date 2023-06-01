package resources

import (
	"path/filepath"
	"strconv"
	"strings"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/robolaunch/devspace-operator/internal"
	"github.com/robolaunch/devspace-operator/internal/configure"
	"github.com/robolaunch/devspace-operator/internal/label"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
)

func GetPersistentVolumeClaim(devspace *devv1alpha1.DevSpace, pvcNamespacedName *types.NamespacedName) *corev1.PersistentVolumeClaim {

	pvc := corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pvcNamespacedName.Name,
			Namespace: pvcNamespacedName.Namespace,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			StorageClassName: &devspace.Spec.Storage.StorageClassConfig.Name,
			AccessModes: []corev1.PersistentVolumeAccessMode{
				devspace.Spec.Storage.StorageClassConfig.AccessMode,
			},
			Resources: corev1.ResourceRequirements{
				Limits: corev1.ResourceList{
					corev1.ResourceName(corev1.ResourceStorage): resource.MustParse(getClaimStorage(pvcNamespacedName, devspace.Spec.Storage.Amount)),
				},
				Requests: corev1.ResourceList{
					corev1.ResourceName(corev1.ResourceStorage): resource.MustParse(getClaimStorage(pvcNamespacedName, devspace.Spec.Storage.Amount)),
				},
			},
		},
	}

	return &pvc
}

func getClaimStorage(pvc *types.NamespacedName, totalStorage int) string {
	storageInt := 0

	if strings.Contains(pvc.Name, "pvc-var") {
		storageInt = totalStorage / 20
	} else if strings.Contains(pvc.Name, "pvc-opt") {
		storageInt = 3 * totalStorage / 10
	} else if strings.Contains(pvc.Name, "pvc-usr") {
		storageInt = totalStorage * 5 / 10
	} else if strings.Contains(pvc.Name, "pvc-etc") {
		storageInt = totalStorage / 20
	} else if strings.Contains(pvc.Name, "pvc-display") {
		storageInt = 100
	} else if strings.Contains(pvc.Name, "pvc-workspace") {
		storageInt = totalStorage / 10
	} else {
		storageInt = 0
	}
	return strconv.Itoa(storageInt) + "M"

}

func GetLoaderJob(devspace *devv1alpha1.DevSpace, jobNamespacedName *types.NamespacedName, hasGPU bool) *batchv1.Job {

	var copierCmdBuilder strings.Builder
	copierCmdBuilder.WriteString("yes | cp -rf /var /ros/;")
	copierCmdBuilder.WriteString(" yes | cp -rf /usr /ros/;")
	copierCmdBuilder.WriteString(" yes | cp -rf /opt /ros/;")
	copierCmdBuilder.WriteString(" yes | cp -rf /etc /ros/;")
	copierCmdBuilder.WriteString(" echo \"DONE\"")

	var preparerCmdBuilder strings.Builder
	preparerCmdBuilder.WriteString("apt-get update")
	preparerCmdBuilder.WriteString(" && apt-get dist-upgrade -y")
	preparerCmdBuilder.WriteString(" && apt-get update")

	copierContainer := corev1.Container{
		Name:            "copier",
		Image:           devspace.Status.Image,
		Command:         internal.Bash(copierCmdBuilder.String()),
		ImagePullPolicy: corev1.PullAlways,
		VolumeMounts: []corev1.VolumeMount{
			configure.GetVolumeMount("/ros/", configure.GetVolumeVar(devspace)),
			configure.GetVolumeMount("/ros/", configure.GetVolumeUsr(devspace)),
			configure.GetVolumeMount("/ros/", configure.GetVolumeOpt(devspace)),
			configure.GetVolumeMount("/ros/", configure.GetVolumeEtc(devspace)),
		},
	}

	preparerContainer := corev1.Container{
		Name:    "preparer",
		Image:   "ubuntu:focal",
		Command: internal.Bash(preparerCmdBuilder.String()),
		VolumeMounts: []corev1.VolumeMount{
			configure.GetVolumeMount("", configure.GetVolumeVar(devspace)),
			configure.GetVolumeMount("", configure.GetVolumeUsr(devspace)),
			configure.GetVolumeMount("", configure.GetVolumeOpt(devspace)),
			configure.GetVolumeMount("", configure.GetVolumeEtc(devspace)),
		},
	}

	podSpec := &corev1.PodSpec{
		InitContainers: []corev1.Container{
			copierContainer,
		},
		Containers: []corev1.Container{
			preparerContainer,
		},
		Volumes: []corev1.Volume{
			configure.GetVolumeVar(devspace),
			configure.GetVolumeUsr(devspace),
			configure.GetVolumeOpt(devspace),
			configure.GetVolumeEtc(devspace),
			configure.GetVolumeWorkspace(devspace),
		},
	}

	if hasGPU {

		var driverInstallerCmdBuilder strings.Builder

		// run /etc/vdi/install-driver.sh
		driverInstallerCmdBuilder.WriteString(filepath.Join("/etc", "vdi", "install-driver.sh"))

		driverInstaller := corev1.Container{
			Name:            "driver-installer",
			Image:           devspace.Status.Image,
			Command:         internal.Bash(driverInstallerCmdBuilder.String()),
			ImagePullPolicy: corev1.PullAlways,
			Env: []corev1.EnvVar{
				internal.Env("NVIDIA_DRIVER_VERSION", "agnostic"),
				internal.Env("RESOLUTION", devspace.Spec.DevSuiteTemplate.DevSpaceVDITemplate.Resolution),
			},
			VolumeMounts: []corev1.VolumeMount{
				configure.GetVolumeMount("", configure.GetVolumeVar(devspace)),
				configure.GetVolumeMount("", configure.GetVolumeUsr(devspace)),
				configure.GetVolumeMount("", configure.GetVolumeOpt(devspace)),
				configure.GetVolumeMount("", configure.GetVolumeEtc(devspace)),
			},
		}

		podSpec.InitContainers = append(podSpec.InitContainers, driverInstaller)

	}

	podSpec.RestartPolicy = corev1.RestartPolicyNever
	podSpec.NodeSelector = label.GetTenancyMap(devspace)

	job := batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      devspace.GetLoaderJobMetadata().Name,
			Namespace: devspace.GetLoaderJobMetadata().Namespace,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: *podSpec,
			},
		},
	}

	return &job
}

func GetDevSuite(devspace *devv1alpha1.DevSpace, rdsNamespacedName *types.NamespacedName) *devv1alpha1.DevSuite {

	labels := devspace.Labels
	labels[internal.TARGET_DEVSPACE_LABEL_KEY] = devspace.Name
	labels[internal.DEVSPACE_DEV_SUITE_OWNED] = "true"

	devSuite := devv1alpha1.DevSuite{
		ObjectMeta: metav1.ObjectMeta{
			Name:      rdsNamespacedName.Name,
			Namespace: rdsNamespacedName.Namespace,
			Labels:    devspace.Labels,
		},
		Spec: devspace.Spec.DevSuiteTemplate,
	}

	return &devSuite

}

func GetWorkspaceManager(devspace *devv1alpha1.DevSpace, wsmNamespacedName *types.NamespacedName) *devv1alpha1.WorkspaceManager {

	labels := devspace.Labels
	labels[internal.TARGET_DEVSPACE_LABEL_KEY] = devspace.Name

	workspaceManager := devv1alpha1.WorkspaceManager{
		ObjectMeta: metav1.ObjectMeta{
			Name:      wsmNamespacedName.Name,
			Namespace: wsmNamespacedName.Namespace,
			Labels:    devspace.Labels,
		},
		Spec: devspace.Spec.WorkspaceManagerTemplate,
	}

	return &workspaceManager

}

func GetCloneCommand(workspaces []devv1alpha1.Workspace, wsKey int) string {

	var cmdBuilder strings.Builder
	for key, repo := range workspaces[wsKey].Repositories {
		cmdBuilder.WriteString("git clone " + repo.URL + " -b " + repo.Branch + " " + key + " &&")
	}
	return cmdBuilder.String()
}
