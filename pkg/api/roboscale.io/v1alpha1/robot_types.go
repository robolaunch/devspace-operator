package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	SchemeBuilder.Register(&Robot{}, &RobotList{})
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Distributions",type=string,JSONPath=`.spec.distributions`
//+kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

// Robot is the custom resource that contains ROS 2 components (Workloads, Cloud VDI, Cloud IDE, ROS Bridge, Configurational Resources), robolaunch Robot instances can be decomposed and distributed to both cloud instances and physical instances using federation.
type Robot struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the Robot.
	Spec RobotSpec `json:"spec,omitempty"`
	// Most recently observed status of the Robot.
	Status RobotStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RobotList contains a list of Robot
type RobotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Robot `json:"items"`
}

// ********************************
// Robot types
// ********************************

// ROS 2 distribution selection. Currently supported distributions are Humble, Foxy, Galactic.
// +kubebuilder:validation:Enum=foxy;galactic;humble
type ROSDistro string

const (
	// ROS Melodic Morenia
	ROSDistroMelodic ROSDistro = "melodic"
	// ROS Noetic Ninjemys
	ROSDistroNoetic ROSDistro = "noetic"
	// ROS 2 Foxy Fitzroy
	ROSDistroFoxy ROSDistro = "foxy"
	// ROS 2 Galactic Geochelone
	ROSDistroGalactic ROSDistro = "galactic"
	// ROS 2 Humble Hawksbill
	ROSDistroHumble ROSDistro = "humble"
)

// Storage class configuration for a volume type.
type StorageClassConfig struct {
	// Storage class name.
	Name string `json:"name,omitempty"`
	// PVC access modes. Currently, only `ReadWriteOnce` is supported.
	AccessMode corev1.PersistentVolumeAccessMode `json:"accessMode,omitempty"`
}

// Robot's resource limitations.
type Storage struct {
	// Specifies how much storage will be allocated in total. Use MB as a unit of measurement. (eg. `10240` is equal to 10 GB)
	// +kubebuilder:default=10000
	Amount int `json:"amount,omitempty"`
	// Storage class selection for robot's volumes.
	StorageClassConfig StorageClassConfig `json:"storageClassConfig,omitempty"`
}

type TLSSecretReference struct {
	// [*alpha*] TLS secret object name.
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// [*alpha*] TLS secret object namespace.
	// +kubebuilder:validation:Required
	Namespace string `json:"namespace"`
}

type RootDNSConfig struct {
	// [*alpha*] Root DNS name..
	// +kubebuilder:validation:Required
	Host string `json:"host"`
}

// RobotSpec defines the desired state of Robot.
type RobotSpec struct {
	// ROS 2 distributions to be used. You can select multiple distributions if they are supported in the same underlying OS.
	// (eg. `foxy` and `galactic` are supported in Ubuntu Focal, so they can be used together but both cannot be used with `humble`)
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=2
	Distributions []ROSDistro `json:"distributions"`
	// Total storage amount to persist via Robot. Unit of measurement is MB. (eg. `10240` corresponds 10 GB)
	// This amount is being shared between different components.
	Storage Storage `json:"storage,omitempty"`
	// Robot development suite template
	RobotDevSuiteTemplate RobotDevSuiteSpec `json:"robotDevSuiteTemplate,omitempty"`
	// Workspace manager template to configure ROS 2 workspaces.
	WorkspaceManagerTemplate WorkspaceManagerSpec `json:"workspaceManagerTemplate,omitempty"`
	// [*alpha*] Switch to development mode if `true`.
	Development bool `json:"development,omitempty"`
	// [*alpha*] Root DNS configuration.
	RootDNSConfig RootDNSConfig `json:"rootDNSConfig,omitempty"`
	// [*alpha*] TLS secret reference.
	TLSSecretReference TLSSecretReference `json:"tlsSecretRef,omitempty"`
}

type VolumeStatuses struct {
	// Holds PVC status of the `/var` directory of underlying OS.
	Var OwnedResourceStatus `json:"varDir,omitempty"`
	// Holds PVC status of the `/etc` directory of underlying OS.
	Etc OwnedResourceStatus `json:"etcDir,omitempty"`
	// Holds PVC status of the `/usr` directory of underlying OS.
	Usr OwnedResourceStatus `json:"usrDir,omitempty"`
	// Holds PVC status of the `/opt` directory of underlying OS.
	Opt OwnedResourceStatus `json:"optDir,omitempty"`
	// Holds PVC status of the workspaces directory of underlying OS.
	Workspace OwnedResourceStatus `json:"workspaceDir,omitempty"`
}

type JobPhase string

const (
	JobActive    JobPhase = "Active"
	JobSucceeded JobPhase = "Succeeded"
	JobFailed    JobPhase = "Failed"
)

type AttachedDevObject struct {
	// Reference to the RobotDevSuite instance.
	Reference corev1.ObjectReference `json:"reference,omitempty"`
	// Status of attached RobotDevSuite.
	Status RobotDevSuiteStatus `json:"status,omitempty"`
}

// RobotStatus defines the observed state of Robot.
type RobotStatus struct {
	// Phase of Robot. It sums the general status of Robot.
	Phase RobotPhase `json:"phase,omitempty"`
	// Main image of Robot. It is derived either from the specifications or determined directly over labels.
	Image string `json:"image,omitempty"`
	// Node that Robot uses. It is selected via tenancy labels.
	NodeName string `json:"nodeName,omitempty"`
	// Robot persists some of the directories of underlying OS inside persistent volumes.
	// This field exposes persistent volume claims that dynamically provision PVs.
	VolumeStatuses VolumeStatuses `json:"volumeStatuses,omitempty"`
	// Status of loader job that configures environment.
	LoaderJobStatus OwnedResourceStatus `json:"loaderJobStatus,omitempty"`
	// Workspace manager instance status if exists.
	WorkspaceManagerStatus WorkspaceManagerInstanceStatus `json:"workspaceManagerStatus,omitempty"`
	// Robot development suite instance status.
	RobotDevSuiteStatus RobotDevSuiteInstanceStatus `json:"robotDevSuiteStatus,omitempty"`
	// [*alpha*] Attached dev object information.
	AttachedDevObjects []AttachedDevObject `json:"attachedDevObjects,omitempty"`
}
