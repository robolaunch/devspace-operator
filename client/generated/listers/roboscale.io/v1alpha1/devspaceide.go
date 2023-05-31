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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/robolaunch/robot-operator/pkg/api/roboscale.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DevSpaceIDELister helps list DevSpaceIDEs.
// All objects returned here must be treated as read-only.
type DevSpaceIDELister interface {
	// List lists all DevSpaceIDEs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DevSpaceIDE, err error)
	// DevSpaceIDEs returns an object that can list and get DevSpaceIDEs.
	DevSpaceIDEs(namespace string) DevSpaceIDENamespaceLister
	DevSpaceIDEListerExpansion
}

// devSpaceIDELister implements the DevSpaceIDELister interface.
type devSpaceIDELister struct {
	indexer cache.Indexer
}

// NewDevSpaceIDELister returns a new DevSpaceIDELister.
func NewDevSpaceIDELister(indexer cache.Indexer) DevSpaceIDELister {
	return &devSpaceIDELister{indexer: indexer}
}

// List lists all DevSpaceIDEs in the indexer.
func (s *devSpaceIDELister) List(selector labels.Selector) (ret []*v1alpha1.DevSpaceIDE, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DevSpaceIDE))
	})
	return ret, err
}

// DevSpaceIDEs returns an object that can list and get DevSpaceIDEs.
func (s *devSpaceIDELister) DevSpaceIDEs(namespace string) DevSpaceIDENamespaceLister {
	return devSpaceIDENamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DevSpaceIDENamespaceLister helps list and get DevSpaceIDEs.
// All objects returned here must be treated as read-only.
type DevSpaceIDENamespaceLister interface {
	// List lists all DevSpaceIDEs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DevSpaceIDE, err error)
	// Get retrieves the DevSpaceIDE from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DevSpaceIDE, error)
	DevSpaceIDENamespaceListerExpansion
}

// devSpaceIDENamespaceLister implements the DevSpaceIDENamespaceLister
// interface.
type devSpaceIDENamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DevSpaceIDEs in the indexer for a given namespace.
func (s devSpaceIDENamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DevSpaceIDE, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DevSpaceIDE))
	})
	return ret, err
}

// Get retrieves the DevSpaceIDE from the indexer for a given namespace and name.
func (s devSpaceIDENamespaceLister) Get(name string) (*v1alpha1.DevSpaceIDE, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("devspaceide"), name)
	}
	return obj.(*v1alpha1.DevSpaceIDE), nil
}
