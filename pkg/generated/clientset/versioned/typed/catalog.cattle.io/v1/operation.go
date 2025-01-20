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

package v1

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1 "github.com/rancher/rancher/pkg/apis/catalog.cattle.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OperationsGetter has a method to return a OperationInterface.
// A group's client should implement this interface.
type OperationsGetter interface {
	Operations(namespace string) OperationInterface
}

// OperationInterface has methods to work with Operation resources.
type OperationInterface interface {
	Create(ctx context.Context, operation *v1.Operation, opts metav1.CreateOptions) (*v1.Operation, error)
	Update(ctx context.Context, operation *v1.Operation, opts metav1.UpdateOptions) (*v1.Operation, error)
	UpdateStatus(ctx context.Context, operation *v1.Operation, opts metav1.UpdateOptions) (*v1.Operation, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Operation, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.OperationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Operation, err error)
	OperationExpansion
}

// operations implements OperationInterface
type operations struct {
	client rest.Interface
	ns     string
}

// newOperations returns a Operations
func newOperations(c *CatalogV1Client, namespace string) *operations {
	return &operations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the operation, and returns the corresponding operation object, and an error if there is any.
func (c *operations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Operation, err error) {
	result = &v1.Operation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("operations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Operations that match those selectors.
func (c *operations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.OperationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.OperationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("operations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested operations.
func (c *operations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("operations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a operation and creates it.  Returns the server's representation of the operation, and an error, if there is any.
func (c *operations) Create(ctx context.Context, operation *v1.Operation, opts metav1.CreateOptions) (result *v1.Operation, err error) {
	result = &v1.Operation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("operations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a operation and updates it. Returns the server's representation of the operation, and an error, if there is any.
func (c *operations) Update(ctx context.Context, operation *v1.Operation, opts metav1.UpdateOptions) (result *v1.Operation, err error) {
	result = &v1.Operation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("operations").
		Name(operation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operation).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *operations) UpdateStatus(ctx context.Context, operation *v1.Operation, opts metav1.UpdateOptions) (result *v1.Operation, err error) {
	result = &v1.Operation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("operations").
		Name(operation.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the operation and deletes it. Returns an error if one occurs.
func (c *operations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("operations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *operations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("operations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched operation.
func (c *operations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Operation, err error) {
	result = &v1.Operation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("operations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
