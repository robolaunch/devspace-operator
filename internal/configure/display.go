package configure

import (
	"github.com/robolaunch/devspace-operator/internal"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

func InjectPodDisplayConfiguration(pod *corev1.Pod, devSpaceVDI devv1alpha1.DevSpaceVDI) *corev1.Pod {

	configurePod(pod, devSpaceVDI)
	for k, container := range pod.Spec.Containers {
		configureContainer(&container, devSpaceVDI)
		pod.Spec.Containers[k] = container
	}

	return pod
}

func configurePod(pod *corev1.Pod, devSpaceVDI devv1alpha1.DevSpaceVDI) {
	volume := GetVolumeX11Unix(&devSpaceVDI)
	pod.Spec.Volumes = append(pod.Spec.Volumes, volume)
}

func configureContainer(container *corev1.Container, devSpaceVDI devv1alpha1.DevSpaceVDI) {
	volume := GetVolumeX11Unix(&devSpaceVDI)
	volumeMount := GetVolumeMount(internal.X11_UNIX_PATH, volume)
	environmentVariables := []corev1.EnvVar{
		internal.Env("DISPLAY", ":0"),
	}

	container.VolumeMounts = append(container.VolumeMounts, volumeMount)
	container.Env = append(container.Env, environmentVariables...)
}
