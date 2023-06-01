package v1alpha1

import (
	"github.com/robolaunch/devspace-operator/internal/label"
	corev1 "k8s.io/api/core/v1"
)

// Generic status for any owned resource.
type OwnedResourceStatus struct {
	// Shows if the owned resource is created.
	Created bool `json:"created"`
	// Reference to the owned resource.
	Reference corev1.ObjectReference `json:"reference,omitempty"`
	// Phase of the owned resource.
	Phase string `json:"phase,omitempty"`
}

type OwnedDevSpaceServiceStatus struct {
	// Generic status for any owned resource.
	Resource OwnedResourceStatus `json:"resource,omitempty"`
	// Address of the devspace service that can be reached from outside.
	Connection string `json:"connection,omitempty"`
}

type OwnedServiceStatus struct {
	// Generic status for any owned resource.
	Resource OwnedResourceStatus `json:"resource,omitempty"`
	// Connection URL.
	URL string `json:"url,omitempty"`
}

type OwnedPodStatus struct {
	// Generic status for any owned resource.
	Resource OwnedResourceStatus `json:"resource,omitempty"`
	// IP of the pod.
	IP string `json:"ip,omitempty"`
}

type DevSuiteInstanceStatus struct {
	// Generic status for any owned resource.
	Resource OwnedResourceStatus `json:"resource,omitempty"`
	// Status of the DevSuite instance.
	Status DevSuiteStatus `json:"status,omitempty"`
}

type WorkspaceManagerInstanceStatus struct {
	// Generic status for any owned resource.
	Resource OwnedResourceStatus `json:"resource,omitempty"`
	// Status of the WorkspaceManager instance.
	Status WorkspaceManagerStatus `json:"status,omitempty"`
}

func GetDevSpaceServiceDNS(devspace DevSpace, prefix, postfix string) string {
	tenancy := label.GetTenancy(&devspace)
	connectionStr := tenancy.Organization + "." + devspace.Spec.RootDNSConfig.Host + GetDevSpaceServicePath(devspace, postfix)

	if prefix != "" {
		connectionStr = prefix + connectionStr
	}

	return connectionStr
}

func GetDevSpaceServicePath(devspace DevSpace, postfix string) string {
	tenancy := label.GetTenancy(&devspace)
	connectionStr := "/" + tenancy.Team +
		"/" + tenancy.Region +
		"/" + tenancy.CloudInstance +
		"/" + devspace.Namespace +
		"/" + devspace.Name

	if postfix != "" {
		connectionStr = connectionStr + postfix
	}

	return connectionStr
}
