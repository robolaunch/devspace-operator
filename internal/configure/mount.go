package configure

import (
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

func GetVolumeVar(devspace *devv1alpha1.DevSpace) corev1.Volume {

	volume := corev1.Volume{
		Name: "var",
		VolumeSource: corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: devspace.GetPVCVarMetadata().Name,
			},
		},
	}

	return volume
}

func GetVolumeOpt(devspace *devv1alpha1.DevSpace) corev1.Volume {

	volume := corev1.Volume{
		Name: "opt",
		VolumeSource: corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: devspace.GetPVCOptMetadata().Name,
			},
		},
	}

	return volume
}

func GetVolumeUsr(devspace *devv1alpha1.DevSpace) corev1.Volume {

	volume := corev1.Volume{
		Name: "usr",
		VolumeSource: corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: devspace.GetPVCUsrMetadata().Name,
			},
		},
	}

	return volume
}

func GetVolumeEtc(devspace *devv1alpha1.DevSpace) corev1.Volume {

	volume := corev1.Volume{
		Name: "etc",
		VolumeSource: corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: devspace.GetPVCEtcMetadata().Name,
			},
		},
	}

	return volume
}

func GetVolumeX11Unix(devSpaceVDI *devv1alpha1.DevSpaceVDI) corev1.Volume {

	volume := corev1.Volume{
		Name: "x11-unix",
		VolumeSource: corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: devSpaceVDI.GetDevSpaceVDIPVCMetadata().Name,
			},
		},
	}

	return volume
}

func GetVolumeWorkspace(devspace *devv1alpha1.DevSpace) corev1.Volume {

	volume := corev1.Volume{
		Name: "workspace",
		VolumeSource: corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: devspace.GetPVCWorkspaceMetadata().Name,
			},
		},
	}

	return volume
}

func GetVolumeDshm() corev1.Volume {

	volume := corev1.Volume{
		Name: "dshm",
		VolumeSource: corev1.VolumeSource{
			EmptyDir: &corev1.EmptyDirVolumeSource{
				Medium: corev1.StorageMediumMemory,
			},
		},
	}

	return volume
}

func GetVolumeXglCache() corev1.Volume {

	volume := corev1.Volume{
		Name: "xgl-cache-vol",
		VolumeSource: corev1.VolumeSource{
			EmptyDir: &corev1.EmptyDirVolumeSource{},
		},
	}

	return volume
}

func GetVolumeMount(
	mountPrefix string,
	volume corev1.Volume,
) corev1.VolumeMount {
	mountPath := mountPrefix + volume.Name

	if volume.Name == "workspace" {
		mountPath = mountPrefix
	}
	if volume.Name == "config-volume" {
		mountPath = mountPrefix
	}
	if volume.Name == "x11-unix" {
		mountPath = mountPrefix
	}
	if volume.Name == "dshm" {
		mountPath = mountPrefix
	}
	if volume.Name == "xgl-cache-vol" {
		mountPath = mountPrefix
	}

	volumeMount := corev1.VolumeMount{
		Name:      volume.Name,
		MountPath: mountPath,
	}

	return volumeMount
}
