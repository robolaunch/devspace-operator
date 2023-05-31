/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	SchemeBuilder.Register(&DevSuite{}, &DevSuiteList{})
	SchemeBuilder.Register(&DevSpaceIDE{}, &DevSpaceIDEList{})
	SchemeBuilder.Register(&DevSpaceVDI{}, &DevSpaceVDIList{})
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DevSuite is a custom resource that creates dynamically configured
// development environments for devspaces.
type DevSuite struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the DevSuite.
	Spec DevSuiteSpec `json:"spec,omitempty"`
	// Most recently observed status of the DevSuite.
	Status DevSuiteStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DevSuiteList contains a list of DevSuite.
type DevSuiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevSuite `json:"items"`
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DevSpaceIDE creates and manages Cloud IDE resources and workloads.
type DevSpaceIDE struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the DevSpaceIDE.
	Spec DevSpaceIDESpec `json:"spec,omitempty"`
	// Most recently observed status of the DevSpaceIDE.
	Status DevSpaceIDEStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DevSpaceIDEList contains a list of DevSpaceIDE.
type DevSpaceIDEList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevSpaceIDE `json:"items"`
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DevSpaceVDI creates and manages Cloud VDI resources and workloads.
type DevSpaceVDI struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the DevSpaceVDI.
	Spec DevSpaceVDISpec `json:"spec,omitempty"`
	// Most recently observed status of the DevSpaceVDI.
	Status DevSpaceVDIStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DevSpaceVDIList contains a list of DevSpaceVDI.
type DevSpaceVDIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevSpaceVDI `json:"items"`
}

// ********************************
// DevSuite types
// ********************************

// DevSuiteSpec defines the desired state of DevSuite.
type DevSuiteSpec struct {
	// If `true`, a Cloud VDI will be provisioned inside development suite.
	VDIEnabled bool `json:"vdiEnabled,omitempty"`
	// Configurational parameters of DevSpaceVDI. Applied if `.spec.vdiEnabled` is set to `true`.
	DevSpaceVDITemplate DevSpaceVDISpec `json:"devSpaceVDITemplate,omitempty"`
	// If `true`, a Cloud IDE will be provisioned inside development suite.
	IDEEnabled bool `json:"ideEnabled,omitempty"`
	// Configurational parameters of DevSpaceIDE. Applied if `.spec.ideEnabled` is set to `true`.
	DevSpaceIDETemplate DevSpaceIDESpec `json:"devSpaceIDETemplate,omitempty"`
}

// DevSuiteStatus defines the observed state of DevSuite.
type DevSuiteStatus struct {
	// Phase of DevSuite.
	Phase DevSuitePhase `json:"phase,omitempty"`
	// Status of DevSpaceVDI.
	DevSpaceVDIStatus OwnedDevSpaceServiceStatus `json:"devSpaceVDIStatus,omitempty"`
	// Status of DevSpaceIDE.
	DevSpaceIDEStatus OwnedDevSpaceServiceStatus `json:"devSpaceIDEStatus,omitempty"`
	// [*alpha*] Indicates if DevSuite is attached to a DevSpace and actively provisioned it's resources.
	Active bool `json:"active,omitempty"`
}

// ********************************
// DevSpaceIDE types
// ********************************

// DevSpaceIDESpec defines the desired state of DevSpaceIDE.
type DevSpaceIDESpec struct {
	// Resource limitations of Cloud IDE.
	Resources Resources `json:"resources,omitempty"`
	// Service type of Cloud IDE. `ClusterIP` and `NodePort` is supported.
	// +kubebuilder:validation:Enum=ClusterIP;NodePort
	// +kubebuilder:default="NodePort"
	ServiceType corev1.ServiceType `json:"serviceType,omitempty"`
	// If `true`, containers of DevSpaceIDE will be privileged containers.
	// It can be used in physical instances where it's necessary to access
	// I/O devices on the host machine.
	// Not recommended to activate this field on cloud instances.
	Privileged bool `json:"privileged,omitempty"`
	// Cloud IDE connects an X11 socket if it's set to `true` and a target DevSpaceVDI resource is set in labels with key `robolaunch.io/target-vdi`.
	// Applications that requires GUI can be executed such as rViz.
	Display bool `json:"display,omitempty"`
	// [*alpha*] DevSpaceIDE will create an Ingress resource if `true`.
	Ingress bool `json:"ingress,omitempty"`
}

// DevSpaceIDEStatus defines the observed state of DevSpaceIDE.
type DevSpaceIDEStatus struct {
	// Phase of DevSpaceIDE.
	Phase DevSpaceIDEPhase `json:"phase,omitempty"`
	// Status of Cloud IDE pod.
	PodStatus OwnedPodStatus `json:"podStatus,omitempty"`
	// Status of Cloud IDE service.
	ServiceStatus OwnedServiceStatus `json:"serviceStatus,omitempty"`
	// Status of Cloud IDE Ingress.
	IngressStatus OwnedResourceStatus `json:"ingressStatus,omitempty"`
}

// ********************************
// DevSpaceVDI types
// ********************************

// VDI resource limits.
type Resources struct {
	// GPU core number that will be allocated.
	GPUCore int `json:"gpuCore,omitempty"`
	// CPU resource limit.
	// +kubebuilder:validation:Pattern=`^([0-9])+(m)$`
	CPU string `json:"cpu,omitempty"`
	// Memory resource limit.
	// +kubebuilder:validation:Pattern=`^([0-9])+(Mi|Gi)$`
	Memory string `json:"memory,omitempty"`
}

// DevSpaceVDISpec defines the desired state of DevSpaceVDI.
type DevSpaceVDISpec struct {
	// Resource limitations of Cloud IDE.
	Resources Resources `json:"resources,omitempty"`
	// Service type of Cloud IDE. `ClusterIP` and `NodePort` is supported.
	// +kubebuilder:validation:Enum=ClusterIP;NodePort
	// +kubebuilder:default="NodePort"
	ServiceType corev1.ServiceType `json:"serviceType,omitempty"`
	// If `true`, containers of DevSpaceIDE will be privileged containers.
	// It can be used in physical instances where it's necessary to access
	// I/O devices on the host machine.
	// Not recommended to activate this field on cloud instances.
	Privileged bool `json:"privileged,omitempty"`
	// NAT1TO1 option for Cloud VDI.
	NAT1TO1 string `json:"nat1to1,omitempty"`
	// UDP port range to used in WebRTC connections.
	// +kubebuilder:validation:Pattern=`^([0-9])+-([0-9])+$`
	// +kubebuilder:validation:Required
	WebRTCPortRange string `json:"webrtcPortRange,omitempty"`
	// VDI screen resolution options. Default is `2048x1152`.
	// +kubebuilder:validation:Enum="2048x1152";"1920x1080";"1600x1200"
	// +kubebuilder:default="2048x1152"
	Resolution string `json:"resolution,omitempty"`
	// [*alpha*] DevSpaceIDE will create an Ingress resource if `true`.
	Ingress bool `json:"ingress,omitempty"`
}

// DevSpaceVDIStatus defines the observed state of DevSpaceVDI.
type DevSpaceVDIStatus struct {
	// Phase of DevSpaceVDI.
	Phase DevSpaceVDIPhase `json:"phase,omitempty"`
	// Status of Cloud VDI pod.
	PodStatus OwnedPodStatus `json:"podStatus,omitempty"`
	// Status of Cloud VDI TCP service.
	ServiceTCPStatus OwnedServiceStatus `json:"serviceTCPStatus,omitempty"`
	// Status of Cloud VDI UDP service.
	ServiceUDPStatus OwnedResourceStatus `json:"serviceUDPStatus,omitempty"`
	// Status of Cloud VDI Ingress.
	IngressStatus OwnedResourceStatus `json:"ingressStatus,omitempty"`
	// Status of Cloud VDI persistent volume claim.
	// This PVC dynamically provisions a volume that is a shared
	// between DevSpaceVDI workloads and other workloads that requests
	// display.
	PVCStatus OwnedResourceStatus `json:"pvcStatus,omitempty"`
}
