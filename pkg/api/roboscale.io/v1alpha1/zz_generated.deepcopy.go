//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachedDevObject) DeepCopyInto(out *AttachedDevObject) {
	*out = *in
	out.Reference = in.Reference
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachedDevObject.
func (in *AttachedDevObject) DeepCopy() *AttachedDevObject {
	if in == nil {
		return nil
	}
	out := new(AttachedDevObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpace) DeepCopyInto(out *DevSpace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpace.
func (in *DevSpace) DeepCopy() *DevSpace {
	if in == nil {
		return nil
	}
	out := new(DevSpace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevSpace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceIDE) DeepCopyInto(out *DevSpaceIDE) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceIDE.
func (in *DevSpaceIDE) DeepCopy() *DevSpaceIDE {
	if in == nil {
		return nil
	}
	out := new(DevSpaceIDE)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevSpaceIDE) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceIDEList) DeepCopyInto(out *DevSpaceIDEList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DevSpaceIDE, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceIDEList.
func (in *DevSpaceIDEList) DeepCopy() *DevSpaceIDEList {
	if in == nil {
		return nil
	}
	out := new(DevSpaceIDEList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevSpaceIDEList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceIDESpec) DeepCopyInto(out *DevSpaceIDESpec) {
	*out = *in
	out.Resources = in.Resources
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceIDESpec.
func (in *DevSpaceIDESpec) DeepCopy() *DevSpaceIDESpec {
	if in == nil {
		return nil
	}
	out := new(DevSpaceIDESpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceIDEStatus) DeepCopyInto(out *DevSpaceIDEStatus) {
	*out = *in
	out.PodStatus = in.PodStatus
	out.ServiceStatus = in.ServiceStatus
	out.IngressStatus = in.IngressStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceIDEStatus.
func (in *DevSpaceIDEStatus) DeepCopy() *DevSpaceIDEStatus {
	if in == nil {
		return nil
	}
	out := new(DevSpaceIDEStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceImage) DeepCopyInto(out *DevSpaceImage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceImage.
func (in *DevSpaceImage) DeepCopy() *DevSpaceImage {
	if in == nil {
		return nil
	}
	out := new(DevSpaceImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceJupyter) DeepCopyInto(out *DevSpaceJupyter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceJupyter.
func (in *DevSpaceJupyter) DeepCopy() *DevSpaceJupyter {
	if in == nil {
		return nil
	}
	out := new(DevSpaceJupyter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevSpaceJupyter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceJupyterList) DeepCopyInto(out *DevSpaceJupyterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DevSpaceJupyter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceJupyterList.
func (in *DevSpaceJupyterList) DeepCopy() *DevSpaceJupyterList {
	if in == nil {
		return nil
	}
	out := new(DevSpaceJupyterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevSpaceJupyterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceJupyterSpec) DeepCopyInto(out *DevSpaceJupyterSpec) {
	*out = *in
	out.Resources = in.Resources
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceJupyterSpec.
func (in *DevSpaceJupyterSpec) DeepCopy() *DevSpaceJupyterSpec {
	if in == nil {
		return nil
	}
	out := new(DevSpaceJupyterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceJupyterStatus) DeepCopyInto(out *DevSpaceJupyterStatus) {
	*out = *in
	out.PodStatus = in.PodStatus
	out.ServiceStatus = in.ServiceStatus
	out.IngressStatus = in.IngressStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceJupyterStatus.
func (in *DevSpaceJupyterStatus) DeepCopy() *DevSpaceJupyterStatus {
	if in == nil {
		return nil
	}
	out := new(DevSpaceJupyterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceList) DeepCopyInto(out *DevSpaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DevSpace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceList.
func (in *DevSpaceList) DeepCopy() *DevSpaceList {
	if in == nil {
		return nil
	}
	out := new(DevSpaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevSpaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceSpec) DeepCopyInto(out *DevSpaceSpec) {
	*out = *in
	out.Environment = in.Environment
	out.Storage = in.Storage
	out.DevSuiteTemplate = in.DevSuiteTemplate
	in.WorkspaceManagerTemplate.DeepCopyInto(&out.WorkspaceManagerTemplate)
	out.RootDNSConfig = in.RootDNSConfig
	out.TLSSecretReference = in.TLSSecretReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceSpec.
func (in *DevSpaceSpec) DeepCopy() *DevSpaceSpec {
	if in == nil {
		return nil
	}
	out := new(DevSpaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceStatus) DeepCopyInto(out *DevSpaceStatus) {
	*out = *in
	out.VolumeStatuses = in.VolumeStatuses
	out.LoaderJobStatus = in.LoaderJobStatus
	out.WorkspaceManagerStatus = in.WorkspaceManagerStatus
	out.DevSuiteStatus = in.DevSuiteStatus
	if in.AttachedDevObjects != nil {
		in, out := &in.AttachedDevObjects, &out.AttachedDevObjects
		*out = make([]AttachedDevObject, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceStatus.
func (in *DevSpaceStatus) DeepCopy() *DevSpaceStatus {
	if in == nil {
		return nil
	}
	out := new(DevSpaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceVDI) DeepCopyInto(out *DevSpaceVDI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceVDI.
func (in *DevSpaceVDI) DeepCopy() *DevSpaceVDI {
	if in == nil {
		return nil
	}
	out := new(DevSpaceVDI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevSpaceVDI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceVDIList) DeepCopyInto(out *DevSpaceVDIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DevSpaceVDI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceVDIList.
func (in *DevSpaceVDIList) DeepCopy() *DevSpaceVDIList {
	if in == nil {
		return nil
	}
	out := new(DevSpaceVDIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevSpaceVDIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceVDISpec) DeepCopyInto(out *DevSpaceVDISpec) {
	*out = *in
	out.Resources = in.Resources
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceVDISpec.
func (in *DevSpaceVDISpec) DeepCopy() *DevSpaceVDISpec {
	if in == nil {
		return nil
	}
	out := new(DevSpaceVDISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSpaceVDIStatus) DeepCopyInto(out *DevSpaceVDIStatus) {
	*out = *in
	out.PodStatus = in.PodStatus
	out.ServiceTCPStatus = in.ServiceTCPStatus
	out.ServiceUDPStatus = in.ServiceUDPStatus
	out.IngressStatus = in.IngressStatus
	out.PVCStatus = in.PVCStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSpaceVDIStatus.
func (in *DevSpaceVDIStatus) DeepCopy() *DevSpaceVDIStatus {
	if in == nil {
		return nil
	}
	out := new(DevSpaceVDIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSuite) DeepCopyInto(out *DevSuite) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSuite.
func (in *DevSuite) DeepCopy() *DevSuite {
	if in == nil {
		return nil
	}
	out := new(DevSuite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevSuite) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSuiteInstanceStatus) DeepCopyInto(out *DevSuiteInstanceStatus) {
	*out = *in
	out.Resource = in.Resource
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSuiteInstanceStatus.
func (in *DevSuiteInstanceStatus) DeepCopy() *DevSuiteInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(DevSuiteInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSuiteList) DeepCopyInto(out *DevSuiteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DevSuite, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSuiteList.
func (in *DevSuiteList) DeepCopy() *DevSuiteList {
	if in == nil {
		return nil
	}
	out := new(DevSuiteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevSuiteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSuiteSpec) DeepCopyInto(out *DevSuiteSpec) {
	*out = *in
	out.DevSpaceVDITemplate = in.DevSpaceVDITemplate
	out.DevSpaceIDETemplate = in.DevSpaceIDETemplate
	out.DevSpaceJupyterTemplate = in.DevSpaceJupyterTemplate
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSuiteSpec.
func (in *DevSuiteSpec) DeepCopy() *DevSuiteSpec {
	if in == nil {
		return nil
	}
	out := new(DevSuiteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevSuiteStatus) DeepCopyInto(out *DevSuiteStatus) {
	*out = *in
	out.DevSpaceVDIStatus = in.DevSpaceVDIStatus
	out.DevSpaceIDEStatus = in.DevSpaceIDEStatus
	out.DevSpaceJupyterStatus = in.DevSpaceJupyterStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevSuiteStatus.
func (in *DevSuiteStatus) DeepCopy() *DevSuiteStatus {
	if in == nil {
		return nil
	}
	out := new(DevSuiteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	out.Application = in.Application
	out.DevSpaceImage = in.DevSpaceImage
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment.
func (in *Environment) DeepCopy() *Environment {
	if in == nil {
		return nil
	}
	out := new(Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OwnedDevSpaceServiceStatus) DeepCopyInto(out *OwnedDevSpaceServiceStatus) {
	*out = *in
	out.Resource = in.Resource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnedDevSpaceServiceStatus.
func (in *OwnedDevSpaceServiceStatus) DeepCopy() *OwnedDevSpaceServiceStatus {
	if in == nil {
		return nil
	}
	out := new(OwnedDevSpaceServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OwnedPodStatus) DeepCopyInto(out *OwnedPodStatus) {
	*out = *in
	out.Resource = in.Resource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnedPodStatus.
func (in *OwnedPodStatus) DeepCopy() *OwnedPodStatus {
	if in == nil {
		return nil
	}
	out := new(OwnedPodStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OwnedResourceStatus) DeepCopyInto(out *OwnedResourceStatus) {
	*out = *in
	out.Reference = in.Reference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnedResourceStatus.
func (in *OwnedResourceStatus) DeepCopy() *OwnedResourceStatus {
	if in == nil {
		return nil
	}
	out := new(OwnedResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OwnedServiceStatus) DeepCopyInto(out *OwnedServiceStatus) {
	*out = *in
	out.Resource = in.Resource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnedServiceStatus.
func (in *OwnedServiceStatus) DeepCopy() *OwnedServiceStatus {
	if in == nil {
		return nil
	}
	out := new(OwnedServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootDNSConfig) DeepCopyInto(out *RootDNSConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootDNSConfig.
func (in *RootDNSConfig) DeepCopy() *RootDNSConfig {
	if in == nil {
		return nil
	}
	out := new(RootDNSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	out.StorageClassConfig = in.StorageClassConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClassConfig) DeepCopyInto(out *StorageClassConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClassConfig.
func (in *StorageClassConfig) DeepCopy() *StorageClassConfig {
	if in == nil {
		return nil
	}
	out := new(StorageClassConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSSecretReference) DeepCopyInto(out *TLSSecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSSecretReference.
func (in *TLSSecretReference) DeepCopy() *TLSSecretReference {
	if in == nil {
		return nil
	}
	out := new(TLSSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeStatuses) DeepCopyInto(out *VolumeStatuses) {
	*out = *in
	out.Var = in.Var
	out.Etc = in.Etc
	out.Usr = in.Usr
	out.Opt = in.Opt
	out.Workspace = in.Workspace
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeStatuses.
func (in *VolumeStatuses) DeepCopy() *VolumeStatuses {
	if in == nil {
		return nil
	}
	out := new(VolumeStatuses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workspace) DeepCopyInto(out *Workspace) {
	*out = *in
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make(map[string]Repository, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspace.
func (in *Workspace) DeepCopy() *Workspace {
	if in == nil {
		return nil
	}
	out := new(Workspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceManager) DeepCopyInto(out *WorkspaceManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceManager.
func (in *WorkspaceManager) DeepCopy() *WorkspaceManager {
	if in == nil {
		return nil
	}
	out := new(WorkspaceManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceManagerInstanceStatus) DeepCopyInto(out *WorkspaceManagerInstanceStatus) {
	*out = *in
	out.Resource = in.Resource
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceManagerInstanceStatus.
func (in *WorkspaceManagerInstanceStatus) DeepCopy() *WorkspaceManagerInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceManagerInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceManagerList) DeepCopyInto(out *WorkspaceManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkspaceManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceManagerList.
func (in *WorkspaceManagerList) DeepCopy() *WorkspaceManagerList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceManagerSpec) DeepCopyInto(out *WorkspaceManagerSpec) {
	*out = *in
	if in.Workspaces != nil {
		in, out := &in.Workspaces, &out.Workspaces
		*out = make([]Workspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceManagerSpec.
func (in *WorkspaceManagerSpec) DeepCopy() *WorkspaceManagerSpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceManagerStatus) DeepCopyInto(out *WorkspaceManagerStatus) {
	*out = *in
	out.ClonerJobStatus = in.ClonerJobStatus
	out.CleanupJobStatus = in.CleanupJobStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceManagerStatus.
func (in *WorkspaceManagerStatus) DeepCopy() *WorkspaceManagerStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceManagerStatus)
	in.DeepCopyInto(out)
	return out
}
