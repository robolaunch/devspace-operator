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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/robolaunch/devspace-operator/client/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Robots returns a RobotInformer.
	Robots() RobotInformer
	// RobotDevSuites returns a RobotDevSuiteInformer.
	RobotDevSuites() RobotDevSuiteInformer
	// DevSpaceIDEs returns a DevSpaceIDEInformer.
	DevSpaceIDEs() DevSpaceIDEInformer
	// DevSpaceVDIs returns a DevSpaceVDIInformer.
	DevSpaceVDIs() DevSpaceVDIInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Robots returns a RobotInformer.
func (v *version) Robots() RobotInformer {
	return &robotInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RobotDevSuites returns a RobotDevSuiteInformer.
func (v *version) RobotDevSuites() RobotDevSuiteInformer {
	return &robotDevSuiteInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DevSpaceIDEs returns a DevSpaceIDEInformer.
func (v *version) DevSpaceIDEs() DevSpaceIDEInformer {
	return &devSpaceIDEInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DevSpaceVDIs returns a DevSpaceVDIInformer.
func (v *version) DevSpaceVDIs() DevSpaceVDIInformer {
	return &devSpaceVDIInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
