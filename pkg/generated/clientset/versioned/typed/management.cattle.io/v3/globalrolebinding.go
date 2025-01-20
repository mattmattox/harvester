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

package v3

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GlobalRoleBindingsGetter has a method to return a GlobalRoleBindingInterface.
// A group's client should implement this interface.
type GlobalRoleBindingsGetter interface {
	GlobalRoleBindings() GlobalRoleBindingInterface
}

// GlobalRoleBindingInterface has methods to work with GlobalRoleBinding resources.
type GlobalRoleBindingInterface interface {
	Create(ctx context.Context, globalRoleBinding *v3.GlobalRoleBinding, opts v1.CreateOptions) (*v3.GlobalRoleBinding, error)
	Update(ctx context.Context, globalRoleBinding *v3.GlobalRoleBinding, opts v1.UpdateOptions) (*v3.GlobalRoleBinding, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.GlobalRoleBinding, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.GlobalRoleBindingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.GlobalRoleBinding, err error)
	GlobalRoleBindingExpansion
}

// globalRoleBindings implements GlobalRoleBindingInterface
type globalRoleBindings struct {
	client rest.Interface
}

// newGlobalRoleBindings returns a GlobalRoleBindings
func newGlobalRoleBindings(c *ManagementV3Client) *globalRoleBindings {
	return &globalRoleBindings{
		client: c.RESTClient(),
	}
}

// Get takes name of the globalRoleBinding, and returns the corresponding globalRoleBinding object, and an error if there is any.
func (c *globalRoleBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.GlobalRoleBinding, err error) {
	result = &v3.GlobalRoleBinding{}
	err = c.client.Get().
		Resource("globalrolebindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlobalRoleBindings that match those selectors.
func (c *globalRoleBindings) List(ctx context.Context, opts v1.ListOptions) (result *v3.GlobalRoleBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.GlobalRoleBindingList{}
	err = c.client.Get().
		Resource("globalrolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested globalRoleBindings.
func (c *globalRoleBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("globalrolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a globalRoleBinding and creates it.  Returns the server's representation of the globalRoleBinding, and an error, if there is any.
func (c *globalRoleBindings) Create(ctx context.Context, globalRoleBinding *v3.GlobalRoleBinding, opts v1.CreateOptions) (result *v3.GlobalRoleBinding, err error) {
	result = &v3.GlobalRoleBinding{}
	err = c.client.Post().
		Resource("globalrolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalRoleBinding).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a globalRoleBinding and updates it. Returns the server's representation of the globalRoleBinding, and an error, if there is any.
func (c *globalRoleBindings) Update(ctx context.Context, globalRoleBinding *v3.GlobalRoleBinding, opts v1.UpdateOptions) (result *v3.GlobalRoleBinding, err error) {
	result = &v3.GlobalRoleBinding{}
	err = c.client.Put().
		Resource("globalrolebindings").
		Name(globalRoleBinding.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalRoleBinding).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the globalRoleBinding and deletes it. Returns an error if one occurs.
func (c *globalRoleBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("globalrolebindings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *globalRoleBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("globalrolebindings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched globalRoleBinding.
func (c *globalRoleBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.GlobalRoleBinding, err error) {
	result = &v3.GlobalRoleBinding{}
	err = c.client.Patch(pt).
		Resource("globalrolebindings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
