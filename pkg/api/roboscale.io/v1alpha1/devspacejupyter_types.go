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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DevSpaceJupyterSpec defines the desired state of DevSpaceJupyter
type DevSpaceJupyterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of DevSpaceJupyter. Edit devspacejupyter_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// DevSpaceJupyterStatus defines the observed state of DevSpaceJupyter
type DevSpaceJupyterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DevSpaceJupyter is the Schema for the devspacejupyters API
type DevSpaceJupyter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DevSpaceJupyterSpec   `json:"spec,omitempty"`
	Status DevSpaceJupyterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DevSpaceJupyterList contains a list of DevSpaceJupyter
type DevSpaceJupyterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevSpaceJupyter `json:"items"`
}
