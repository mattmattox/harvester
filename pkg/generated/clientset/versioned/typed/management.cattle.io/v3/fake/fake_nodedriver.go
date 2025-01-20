/*
Copyright 2025 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodeDrivers implements NodeDriverInterface
type FakeNodeDrivers struct {
	Fake *FakeManagementV3
}

var nodedriversResource = v3.SchemeGroupVersion.WithResource("nodedrivers")

var nodedriversKind = v3.SchemeGroupVersion.WithKind("NodeDriver")

// Get takes name of the nodeDriver, and returns the corresponding nodeDriver object, and an error if there is any.
func (c *FakeNodeDrivers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.NodeDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(nodedriversResource, name), &v3.NodeDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.NodeDriver), err
}

// List takes label and field selectors, and returns the list of NodeDrivers that match those selectors.
func (c *FakeNodeDrivers) List(ctx context.Context, opts v1.ListOptions) (result *v3.NodeDriverList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(nodedriversResource, nodedriversKind, opts), &v3.NodeDriverList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.NodeDriverList{ListMeta: obj.(*v3.NodeDriverList).ListMeta}
	for _, item := range obj.(*v3.NodeDriverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeDrivers.
func (c *FakeNodeDrivers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(nodedriversResource, opts))
}

// Create takes the representation of a nodeDriver and creates it.  Returns the server's representation of the nodeDriver, and an error, if there is any.
func (c *FakeNodeDrivers) Create(ctx context.Context, nodeDriver *v3.NodeDriver, opts v1.CreateOptions) (result *v3.NodeDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nodedriversResource, nodeDriver), &v3.NodeDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.NodeDriver), err
}

// Update takes the representation of a nodeDriver and updates it. Returns the server's representation of the nodeDriver, and an error, if there is any.
func (c *FakeNodeDrivers) Update(ctx context.Context, nodeDriver *v3.NodeDriver, opts v1.UpdateOptions) (result *v3.NodeDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(nodedriversResource, nodeDriver), &v3.NodeDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.NodeDriver), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodeDrivers) UpdateStatus(ctx context.Context, nodeDriver *v3.NodeDriver, opts v1.UpdateOptions) (*v3.NodeDriver, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(nodedriversResource, "status", nodeDriver), &v3.NodeDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.NodeDriver), err
}

// Delete takes name of the nodeDriver and deletes it. Returns an error if one occurs.
func (c *FakeNodeDrivers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(nodedriversResource, name, opts), &v3.NodeDriver{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeDrivers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(nodedriversResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.NodeDriverList{})
	return err
}

// Patch applies the patch and returns the patched nodeDriver.
func (c *FakeNodeDrivers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.NodeDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodedriversResource, name, pt, data, subresources...), &v3.NodeDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.NodeDriver), err
}
