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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDevSpaceIDEs implements DevSpaceIDEInterface
type FakeDevSpaceIDEs struct {
	Fake *FakeRoboscaleV1alpha1
	ns   string
}

var devspaceidesResource = schema.GroupVersionResource{Group: "roboscale.io", Version: "v1alpha1", Resource: "devspaceides"}

var devspaceidesKind = schema.GroupVersionKind{Group: "roboscale.io", Version: "v1alpha1", Kind: "DevSpaceIDE"}

// Get takes name of the devSpaceIDE, and returns the corresponding devSpaceIDE object, and an error if there is any.
func (c *FakeDevSpaceIDEs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DevSpaceIDE, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(devspaceidesResource, c.ns, name), &v1alpha1.DevSpaceIDE{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevSpaceIDE), err
}

// List takes label and field selectors, and returns the list of DevSpaceIDEs that match those selectors.
func (c *FakeDevSpaceIDEs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DevSpaceIDEList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(devspaceidesResource, devspaceidesKind, c.ns, opts), &v1alpha1.DevSpaceIDEList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DevSpaceIDEList{ListMeta: obj.(*v1alpha1.DevSpaceIDEList).ListMeta}
	for _, item := range obj.(*v1alpha1.DevSpaceIDEList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested devSpaceIDEs.
func (c *FakeDevSpaceIDEs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(devspaceidesResource, c.ns, opts))

}

// Create takes the representation of a devSpaceIDE and creates it.  Returns the server's representation of the devSpaceIDE, and an error, if there is any.
func (c *FakeDevSpaceIDEs) Create(ctx context.Context, devSpaceIDE *v1alpha1.DevSpaceIDE, opts v1.CreateOptions) (result *v1alpha1.DevSpaceIDE, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(devspaceidesResource, c.ns, devSpaceIDE), &v1alpha1.DevSpaceIDE{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevSpaceIDE), err
}

// Update takes the representation of a devSpaceIDE and updates it. Returns the server's representation of the devSpaceIDE, and an error, if there is any.
func (c *FakeDevSpaceIDEs) Update(ctx context.Context, devSpaceIDE *v1alpha1.DevSpaceIDE, opts v1.UpdateOptions) (result *v1alpha1.DevSpaceIDE, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(devspaceidesResource, c.ns, devSpaceIDE), &v1alpha1.DevSpaceIDE{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevSpaceIDE), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDevSpaceIDEs) UpdateStatus(ctx context.Context, devSpaceIDE *v1alpha1.DevSpaceIDE, opts v1.UpdateOptions) (*v1alpha1.DevSpaceIDE, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(devspaceidesResource, "status", c.ns, devSpaceIDE), &v1alpha1.DevSpaceIDE{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevSpaceIDE), err
}

// Delete takes name of the devSpaceIDE and deletes it. Returns an error if one occurs.
func (c *FakeDevSpaceIDEs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(devspaceidesResource, c.ns, name, opts), &v1alpha1.DevSpaceIDE{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDevSpaceIDEs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(devspaceidesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DevSpaceIDEList{})
	return err
}

// Patch applies the patch and returns the patched devSpaceIDE.
func (c *FakeDevSpaceIDEs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DevSpaceIDE, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(devspaceidesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DevSpaceIDE{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevSpaceIDE), err
}
