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

// FakeDevSpaceVDIs implements DevSpaceVDIInterface
type FakeDevSpaceVDIs struct {
	Fake *FakeRoboscaleV1alpha1
	ns   string
}

var devspacevdisResource = schema.GroupVersionResource{Group: "roboscale.io", Version: "v1alpha1", Resource: "devspacevdis"}

var devspacevdisKind = schema.GroupVersionKind{Group: "roboscale.io", Version: "v1alpha1", Kind: "DevSpaceVDI"}

// Get takes name of the devSpaceVDI, and returns the corresponding devSpaceVDI object, and an error if there is any.
func (c *FakeDevSpaceVDIs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DevSpaceVDI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(devspacevdisResource, c.ns, name), &v1alpha1.DevSpaceVDI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevSpaceVDI), err
}

// List takes label and field selectors, and returns the list of DevSpaceVDIs that match those selectors.
func (c *FakeDevSpaceVDIs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DevSpaceVDIList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(devspacevdisResource, devspacevdisKind, c.ns, opts), &v1alpha1.DevSpaceVDIList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DevSpaceVDIList{ListMeta: obj.(*v1alpha1.DevSpaceVDIList).ListMeta}
	for _, item := range obj.(*v1alpha1.DevSpaceVDIList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested devSpaceVDIs.
func (c *FakeDevSpaceVDIs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(devspacevdisResource, c.ns, opts))

}

// Create takes the representation of a devSpaceVDI and creates it.  Returns the server's representation of the devSpaceVDI, and an error, if there is any.
func (c *FakeDevSpaceVDIs) Create(ctx context.Context, devSpaceVDI *v1alpha1.DevSpaceVDI, opts v1.CreateOptions) (result *v1alpha1.DevSpaceVDI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(devspacevdisResource, c.ns, devSpaceVDI), &v1alpha1.DevSpaceVDI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevSpaceVDI), err
}

// Update takes the representation of a devSpaceVDI and updates it. Returns the server's representation of the devSpaceVDI, and an error, if there is any.
func (c *FakeDevSpaceVDIs) Update(ctx context.Context, devSpaceVDI *v1alpha1.DevSpaceVDI, opts v1.UpdateOptions) (result *v1alpha1.DevSpaceVDI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(devspacevdisResource, c.ns, devSpaceVDI), &v1alpha1.DevSpaceVDI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevSpaceVDI), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDevSpaceVDIs) UpdateStatus(ctx context.Context, devSpaceVDI *v1alpha1.DevSpaceVDI, opts v1.UpdateOptions) (*v1alpha1.DevSpaceVDI, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(devspacevdisResource, "status", c.ns, devSpaceVDI), &v1alpha1.DevSpaceVDI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevSpaceVDI), err
}

// Delete takes name of the devSpaceVDI and deletes it. Returns an error if one occurs.
func (c *FakeDevSpaceVDIs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(devspacevdisResource, c.ns, name, opts), &v1alpha1.DevSpaceVDI{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDevSpaceVDIs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(devspacevdisResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DevSpaceVDIList{})
	return err
}

// Patch applies the patch and returns the patched devSpaceVDI.
func (c *FakeDevSpaceVDIs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DevSpaceVDI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(devspacevdisResource, c.ns, name, pt, data, subresources...), &v1alpha1.DevSpaceVDI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevSpaceVDI), err
}
