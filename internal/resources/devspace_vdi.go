package resources

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/robolaunch/devspace-operator/internal"
	"github.com/robolaunch/devspace-operator/internal/configure"
	"github.com/robolaunch/devspace-operator/internal/label"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func getDevSpaceVDISelector(devSpaceVDI devv1alpha1.DevSpaceVDI) map[string]string {
	return map[string]string{
		"devSpaceVDI": devSpaceVDI.Name,
	}
}

func GetDevSpaceVDIPVC(devSpaceVDI *devv1alpha1.DevSpaceVDI, pvcNamespacedName *types.NamespacedName, devspace devv1alpha1.DevSpace) *corev1.PersistentVolumeClaim {

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
					corev1.ResourceName(corev1.ResourceStorage): resource.MustParse("100"),
				},
				Requests: corev1.ResourceList{
					corev1.ResourceName(corev1.ResourceStorage): resource.MustParse("100"),
				},
			},
		},
	}

	return &pvc
}

func GetDevSpaceVDIPod(devSpaceVDI *devv1alpha1.DevSpaceVDI, podNamespacedName *types.NamespacedName, devspace devv1alpha1.DevSpace, node corev1.Node) *corev1.Pod {

	// add tcp port
	ports := []corev1.ContainerPort{
		{
			Name:          "http",
			ContainerPort: 8055,
			Protocol:      corev1.ProtocolTCP,
		},
	}

	// add udp ports
	rangeLimits := strings.Split(devSpaceVDI.Spec.WebRTCPortRange, "-")
	rangeStart, _ := strconv.Atoi(rangeLimits[0])
	rangeEnd, _ := strconv.Atoi(rangeLimits[1])

	counter := 0
	for i := rangeStart; i <= rangeEnd; i++ {
		counter++
		ports = append(ports, corev1.ContainerPort{
			Name:          "webrtc-" + strconv.Itoa(counter),
			ContainerPort: int32(i),
			Protocol:      corev1.ProtocolUDP,
		})
	}

	icelite := "true"
	if devSpaceVDI.Spec.NAT1TO1 != "" {
		icelite = "false"
	}

	var cmdBuilder strings.Builder
	cmdBuilder.WriteString(filepath.Join("/etc", "vdi", "generate-xorg.sh") + " && ")
	cmdBuilder.WriteString("supervisord -c " + filepath.Join("/etc", "vdi", "supervisord.conf"))

	labels := getDevSpaceVDISelector(*devSpaceVDI)
	for k, v := range devSpaceVDI.Labels {
		labels[k] = v
	}

	vdiPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podNamespacedName.Name,
			Namespace: podNamespacedName.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "vdi",
					Image:   devspace.Status.Image,
					Command: internal.Bash(cmdBuilder.String()),
					Env: []corev1.EnvVar{
						internal.Env("VIDEO_PORT", "DFP"),
						internal.Env("NEKO_BIND", ":8055"),
						internal.Env("NEKO_EPR", devSpaceVDI.Spec.WebRTCPortRange),
						internal.Env("NEKO_ICELITE", icelite),
						internal.Env("NEKO_NAT1TO1", devSpaceVDI.Spec.NAT1TO1),
						internal.Env("RESOLUTION", devSpaceVDI.Spec.Resolution),
					},
					Stdin: true,
					TTY:   true,
					Ports: ports,
					VolumeMounts: []corev1.VolumeMount{
						configure.GetVolumeMount("", configure.GetVolumeVar(&devspace)),
						configure.GetVolumeMount("", configure.GetVolumeUsr(&devspace)),
						configure.GetVolumeMount("", configure.GetVolumeOpt(&devspace)),
						configure.GetVolumeMount("", configure.GetVolumeEtc(&devspace)),
						configure.GetVolumeMount(devspace.Spec.WorkspaceManagerTemplate.WorkspacesPath, configure.GetVolumeWorkspace(&devspace)),
						configure.GetVolumeMount("/dev/shm", configure.GetVolumeDshm()),
						configure.GetVolumeMount("/cache", configure.GetVolumeXglCache()),
					},
					Resources: corev1.ResourceRequirements{
						Limits: getResourceLimits(devSpaceVDI.Spec.Resources),
					},
					ImagePullPolicy:          corev1.PullAlways,
					TerminationMessagePolicy: corev1.TerminationMessageReadFile,
					SecurityContext: &corev1.SecurityContext{
						Privileged: &devSpaceVDI.Spec.Privileged,
					},
				},
			},
			Volumes: []corev1.Volume{
				configure.GetVolumeVar(&devspace),
				configure.GetVolumeUsr(&devspace),
				configure.GetVolumeOpt(&devspace),
				configure.GetVolumeEtc(&devspace),
				configure.GetVolumeWorkspace(&devspace),
				configure.GetVolumeDshm(),
				configure.GetVolumeXglCache(),
			},
			RestartPolicy: corev1.RestartPolicyNever,
		},
	}

	configure.SchedulePod(vdiPod, label.GetTenancyMap(devSpaceVDI))
	configure.InjectGenericEnvironmentVariables(vdiPod, devspace)
	configure.InjectPodDisplayConfiguration(vdiPod, *devSpaceVDI)
	configure.InjectRuntimeClass(vdiPod, devspace, node)

	return vdiPod
}

func GetDevSpaceVDIServiceTCP(devSpaceVDI *devv1alpha1.DevSpaceVDI, svcNamespacedName *types.NamespacedName) *corev1.Service {

	ports := []corev1.ServicePort{
		{
			Port: 8055,
			TargetPort: intstr.IntOrString{
				IntVal: int32(8055),
			},
			Protocol: corev1.ProtocolTCP,
			Name:     "http",
		},
	}

	serviceSpec := corev1.ServiceSpec{
		Type:     devSpaceVDI.Spec.ServiceType,
		Selector: getDevSpaceVDISelector(*devSpaceVDI),
		Ports:    ports,
	}

	service := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      svcNamespacedName.Name,
			Namespace: svcNamespacedName.Namespace,
		},
		Spec: serviceSpec,
	}

	return &service
}

func GetDevSpaceVDIServiceUDP(devSpaceVDI *devv1alpha1.DevSpaceVDI, svcNamespacedName *types.NamespacedName) *corev1.Service {

	ports := []corev1.ServicePort{}

	// add udp ports
	rangeLimits := strings.Split(devSpaceVDI.Spec.WebRTCPortRange, "-")
	rangeStart, _ := strconv.Atoi(rangeLimits[0])
	rangeEnd, _ := strconv.Atoi(rangeLimits[1])

	counter := 0
	for i := rangeStart; i <= rangeEnd; i++ {
		counter++
		ports = append(ports, corev1.ServicePort{
			Name: "webrtc-" + strconv.Itoa(counter),
			Port: int32(i),
			TargetPort: intstr.IntOrString{
				IntVal: int32(i),
			},
			NodePort: int32(i),
			Protocol: corev1.ProtocolUDP,
		})
	}

	serviceSpec := corev1.ServiceSpec{
		Type:                  corev1.ServiceTypeNodePort,
		ExternalTrafficPolicy: corev1.ServiceExternalTrafficPolicyTypeLocal,
		Selector:              getDevSpaceVDISelector(*devSpaceVDI),
		Ports:                 ports,
	}

	service := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      svcNamespacedName.Name,
			Namespace: svcNamespacedName.Namespace,
		},
		Spec: serviceSpec,
	}

	return &service
}

func GetDevSpaceVDIIngress(devSpaceVDI *devv1alpha1.DevSpaceVDI, ingressNamespacedName *types.NamespacedName, devspace devv1alpha1.DevSpace) *networkingv1.Ingress {

	tenancy := label.GetTenancy(&devspace)

	rootDNSConfig := devspace.Spec.RootDNSConfig
	secretName := devspace.Spec.TLSSecretReference.Name

	annotations := map[string]string{
		internal.INGRESS_AUTH_URL_KEY:                   fmt.Sprintf(internal.INGRESS_AUTH_URL_VAL, tenancy.Organization, rootDNSConfig.Host),
		internal.INGRESS_AUTH_SIGNIN_KEY:                fmt.Sprintf(internal.INGRESS_AUTH_SIGNIN_VAL, tenancy.Organization, rootDNSConfig.Host),
		internal.INGRESS_AUTH_RESPONSE_HEADERS_KEY:      internal.INGRESS_AUTH_RESPONSE_HEADERS_VAL,
		internal.INGRESS_CONFIGURATION_SNIPPET_KEY:      internal.INGRESS_VDI_CONFIGURATION_SNIPPET_VAL,
		internal.INGRESS_CERT_MANAGER_KEY:               internal.INGRESS_CERT_MANAGER_VAL,
		internal.INGRESS_NGINX_PROXY_BUFFER_SIZE_KEY:    internal.INGRESS_NGINX_PROXY_BUFFER_SIZE_VAL,
		internal.INGRESS_NGINX_PROXY_BUFFERS_NUMBER_KEY: internal.INGRESS_VDI_NGINX_PROXY_BUFFERS_NUMBER_VAL,
		internal.INGRESS_NGINX_REWRITE_TARGET_KEY:       internal.INGRESS_NGINX_REWRITE_TARGET_VAL,
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
								Path:     devv1alpha1.GetDevSpaceServicePath(devspace, "/vdi") + "(/|$)(.*)",
								PathType: &pathTypePrefix,
								Backend: networkingv1.IngressBackend{
									Service: &networkingv1.IngressServiceBackend{
										Name: devSpaceVDI.GetDevSpaceVDIServiceTCPMetadata().Name,
										Port: networkingv1.ServiceBackendPort{
											Number: int32(8055),
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

func getResourceLimits(resources devv1alpha1.Resources) corev1.ResourceList {
	resourceLimits := corev1.ResourceList{}
	if resources.GPUCore != 0 {
		resourceLimits["nvidia.com/gpu"] = resource.MustParse(strconv.Itoa(resources.GPUCore))
	}
	if resources.CPU != "" {
		resourceLimits["cpu"] = resource.MustParse(resources.CPU)
	}
	if resources.Memory != "" {
		resourceLimits["memory"] = resource.MustParse(resources.Memory)
	}

	return resourceLimits
}
