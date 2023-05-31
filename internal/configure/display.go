package configure

import (
	"github.com/robolaunch/devspace-operator/internal"
	robotv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

func InjectPodDisplayConfiguration(pod *corev1.Pod, robotVDI robotv1alpha1.RobotVDI) *corev1.Pod {

	configurePod(pod, robotVDI)
	for k, container := range pod.Spec.Containers {
		configureContainer(&container, robotVDI)
		pod.Spec.Containers[k] = container
	}

	return pod
}

func configurePod(pod *corev1.Pod, robotVDI robotv1alpha1.RobotVDI) {
	volume := GetVolumeX11Unix(&robotVDI)
	pod.Spec.Volumes = append(pod.Spec.Volumes, volume)
}

func configureContainer(container *corev1.Container, robotVDI robotv1alpha1.RobotVDI) {
	volume := GetVolumeX11Unix(&robotVDI)
	volumeMount := GetVolumeMount(internal.X11_UNIX_PATH, volume)
	environmentVariables := []corev1.EnvVar{
		internal.Env("DISPLAY", ":0"),
	}

	container.VolumeMounts = append(container.VolumeMounts, volumeMount)
	container.Env = append(container.Env, environmentVariables...)
}
