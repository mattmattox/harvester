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

// RancherUserNotificationsGetter has a method to return a RancherUserNotificationInterface.
// A group's client should implement this interface.
type RancherUserNotificationsGetter interface {
	RancherUserNotifications() RancherUserNotificationInterface
}

// RancherUserNotificationInterface has methods to work with RancherUserNotification resources.
type RancherUserNotificationInterface interface {
	Create(ctx context.Context, rancherUserNotification *v3.RancherUserNotification, opts v1.CreateOptions) (*v3.RancherUserNotification, error)
	Update(ctx context.Context, rancherUserNotification *v3.RancherUserNotification, opts v1.UpdateOptions) (*v3.RancherUserNotification, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.RancherUserNotification, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.RancherUserNotificationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.RancherUserNotification, err error)
	RancherUserNotificationExpansion
}

// rancherUserNotifications implements RancherUserNotificationInterface
type rancherUserNotifications struct {
	client rest.Interface
}

// newRancherUserNotifications returns a RancherUserNotifications
func newRancherUserNotifications(c *ManagementV3Client) *rancherUserNotifications {
	return &rancherUserNotifications{
		client: c.RESTClient(),
	}
}

// Get takes name of the rancherUserNotification, and returns the corresponding rancherUserNotification object, and an error if there is any.
func (c *rancherUserNotifications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.RancherUserNotification, err error) {
	result = &v3.RancherUserNotification{}
	err = c.client.Get().
		Resource("rancherusernotifications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RancherUserNotifications that match those selectors.
func (c *rancherUserNotifications) List(ctx context.Context, opts v1.ListOptions) (result *v3.RancherUserNotificationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.RancherUserNotificationList{}
	err = c.client.Get().
		Resource("rancherusernotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested rancherUserNotifications.
func (c *rancherUserNotifications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("rancherusernotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a rancherUserNotification and creates it.  Returns the server's representation of the rancherUserNotification, and an error, if there is any.
func (c *rancherUserNotifications) Create(ctx context.Context, rancherUserNotification *v3.RancherUserNotification, opts v1.CreateOptions) (result *v3.RancherUserNotification, err error) {
	result = &v3.RancherUserNotification{}
	err = c.client.Post().
		Resource("rancherusernotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(rancherUserNotification).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a rancherUserNotification and updates it. Returns the server's representation of the rancherUserNotification, and an error, if there is any.
func (c *rancherUserNotifications) Update(ctx context.Context, rancherUserNotification *v3.RancherUserNotification, opts v1.UpdateOptions) (result *v3.RancherUserNotification, err error) {
	result = &v3.RancherUserNotification{}
	err = c.client.Put().
		Resource("rancherusernotifications").
		Name(rancherUserNotification.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(rancherUserNotification).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the rancherUserNotification and deletes it. Returns an error if one occurs.
func (c *rancherUserNotifications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("rancherusernotifications").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *rancherUserNotifications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("rancherusernotifications").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched rancherUserNotification.
func (c *rancherUserNotifications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.RancherUserNotification, err error) {
	result = &v3.RancherUserNotification{}
	err = c.client.Patch(pt).
		Resource("rancherusernotifications").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
