package resources

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robolaunch/devspace-operator/internal"
	"github.com/robolaunch/devspace-operator/internal/configure"
	"github.com/robolaunch/devspace-operator/internal/label"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func getDevSpaceIDESelector(devSpaceIDE devv1alpha1.DevSpaceIDE) map[string]string {
	return map[string]string{
		"devSpaceIDE": devSpaceIDE.Name,
	}
}

var devspaceIDEContainerPort int = 9000

func GetDevSpaceIDEPod(devSpaceIDE *devv1alpha1.DevSpaceIDE, podNamespacedName *types.NamespacedName, devspace devv1alpha1.DevSpace, devSpaceVDI devv1alpha1.DevSpaceVDI, node corev1.Node) *corev1.Pod {

	// discovery server

	var cmdBuilder strings.Builder
	cmdBuilder.WriteString("code-server " + devspace.Spec.WorkspaceManagerTemplate.WorkspacesPath + " --bind-addr 0.0.0.0:" + strconv.Itoa(devspaceIDEContainerPort) + " --auth none")

	labels := getDevSpaceIDESelector(*devSpaceIDE)
	for k, v := range devSpaceIDE.Labels {
		labels[k] = v
	}

	pod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podNamespacedName.Name,
			Namespace: podNamespacedName.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "code-server",
					Image:   devspace.Status.Image,
					Command: internal.BashWithUser(cmdBuilder.String(), "robolaunch"),
					Env: []corev1.EnvVar{
						internal.Env("DEVSPACE_NAMESPACE", devspace.Namespace),
						internal.Env("DEVSPACE_NAME", devspace.Name),
						internal.Env("TERM", "xterm-256color"),
					},
					VolumeMounts: []corev1.VolumeMount{
						configure.GetVolumeMount("", configure.GetVolumeVar(&devspace)),
						configure.GetVolumeMount("", configure.GetVolumeUsr(&devspace)),
						configure.GetVolumeMount("", configure.GetVolumeOpt(&devspace)),
						configure.GetVolumeMount("", configure.GetVolumeEtc(&devspace)),
						configure.GetVolumeMount(devspace.Spec.WorkspaceManagerTemplate.WorkspacesPath, configure.GetVolumeWorkspace(&devspace)),
					},
					Ports: []corev1.ContainerPort{
						{
							Name:          "code-server",
							ContainerPort: int32(devspaceIDEContainerPort),
						},
					},
					Resources: corev1.ResourceRequirements{
						Limits: getResourceLimits(devSpaceIDE.Spec.Resources),
					},
					SecurityContext: &corev1.SecurityContext{
						Privileged: &devSpaceIDE.Spec.Privileged,
					},
				},
			},
			Volumes: []corev1.Volume{
				configure.GetVolumeVar(&devspace),
				configure.GetVolumeUsr(&devspace),
				configure.GetVolumeOpt(&devspace),
				configure.GetVolumeEtc(&devspace),
				configure.GetVolumeWorkspace(&devspace),
			},
			RestartPolicy: corev1.RestartPolicyNever,
		},
	}

	configure.SchedulePod(&pod, label.GetTenancyMap(devSpaceIDE))
	configure.InjectGenericEnvironmentVariables(&pod, devspace)
	configure.InjectRuntimeClass(&pod, devspace, node)
	if devSpaceIDE.Spec.Display && label.GetTargetDevSpaceVDI(devSpaceIDE) != "" {
		// TODO: Add control for validating DevSpace VDI
		configure.InjectPodDisplayConfiguration(&pod, devSpaceVDI)
	}

	return &pod
}

func GetDevSpaceIDEService(devSpaceIDE *devv1alpha1.DevSpaceIDE, svcNamespacedName *types.NamespacedName) *corev1.Service {

	serviceSpec := corev1.ServiceSpec{
		Type:     devSpaceIDE.Spec.ServiceType,
		Selector: getDevSpaceIDESelector(*devSpaceIDE),
		Ports: []corev1.ServicePort{
			{
				Port: int32(devspaceIDEContainerPort),
				TargetPort: intstr.IntOrString{
					IntVal: int32(devspaceIDEContainerPort),
				},
				Protocol: corev1.ProtocolTCP,
				Name:     "cloud-ide",
			},
		},
	}

	service := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      devSpaceIDE.GetDevSpaceIDEServiceMetadata().Name,
			Namespace: devSpaceIDE.GetDevSpaceIDEServiceMetadata().Namespace,
		},
		Spec: serviceSpec,
	}

	return &service
}

func GetDevSpaceIDEIngress(devSpaceIDE *devv1alpha1.DevSpaceIDE, ingressNamespacedName *types.NamespacedName, devspace devv1alpha1.DevSpace) *networkingv1.Ingress {

	tenancy := label.GetTenancy(&devspace)

	rootDNSConfig := devspace.Spec.RootDNSConfig
	secretName := devspace.Spec.TLSSecretReference.Name

	annotations := map[string]string{
		internal.INGRESS_AUTH_URL_KEY:                fmt.Sprintf(internal.INGRESS_AUTH_URL_VAL, tenancy.Organization, rootDNSConfig.Host),
		internal.INGRESS_AUTH_SIGNIN_KEY:             fmt.Sprintf(internal.INGRESS_AUTH_SIGNIN_VAL, tenancy.Organization, rootDNSConfig.Host),
		internal.INGRESS_AUTH_RESPONSE_HEADERS_KEY:   internal.INGRESS_AUTH_RESPONSE_HEADERS_VAL,
		internal.INGRESS_CONFIGURATION_SNIPPET_KEY:   internal.INGRESS_IDE_CONFIGURATION_SNIPPET_VAL,
		internal.INGRESS_CERT_MANAGER_KEY:            internal.INGRESS_CERT_MANAGER_VAL,
		internal.INGRESS_NGINX_PROXY_BUFFER_SIZE_KEY: internal.INGRESS_NGINX_PROXY_BUFFER_SIZE_VAL,
		internal.INGRESS_NGINX_REWRITE_TARGET_KEY:    internal.INGRESS_NGINX_REWRITE_TARGET_VAL,
	}

	pathTypePrefix := networkingv1.PathTypePrefix
	ingressClassNameNginx := "nginx"

	ingressSpec := networkingv1.IngressSpec{
		TLS: []networkingv1.IngressTLS{
			{
				Hosts: []string{
					tenancy.Organization + "." + rootDNSConfig.Host,
				},
				SecretName: secretName,
			},
		},
		Rules: []networkingv1.IngressRule{
			{
				Host: tenancy.Organization + "." + rootDNSConfig.Host,
				IngressRuleValue: networkingv1.IngressRuleValue{
					HTTP: &networkingv1.HTTPIngressRuleValue{
						Paths: []networkingv1.HTTPIngressPath{
							{
								Path:     devv1alpha1.GetDevSpaceServicePath(devspace, "/ide") + "(/|$)(.*)",
								PathType: &pathTypePrefix,
								Backend: networkingv1.IngressBackend{
									Service: &networkingv1.IngressServiceBackend{
										Name: devSpaceIDE.GetDevSpaceIDEServiceMetadata().Name,
										Port: networkingv1.ServiceBackendPort{
											Number: int32(devspaceIDEContainerPort),
										},
									},
								},
							},
						},
					},
				},
			},
		},
		IngressClassName: &ingressClassNameNginx,
	}

	ingress := &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:        ingressNamespacedName.Name,
			Namespace:   ingressNamespacedName.Namespace,
			Annotations: annotations,
		},
		Spec: ingressSpec,
	}

	return ingress
}
