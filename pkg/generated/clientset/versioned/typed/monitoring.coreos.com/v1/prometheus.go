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
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PrometheusesGetter has a method to return a PrometheusInterface.
// A group's client should implement this interface.
type PrometheusesGetter interface {
	Prometheuses(namespace string) PrometheusInterface
}

// PrometheusInterface has methods to work with Prometheus resources.
type PrometheusInterface interface {
	Create(ctx context.Context, prometheus *v1.Prometheus, opts metav1.CreateOptions) (*v1.Prometheus, error)
	Update(ctx context.Context, prometheus *v1.Prometheus, opts metav1.UpdateOptions) (*v1.Prometheus, error)
	UpdateStatus(ctx context.Context, prometheus *v1.Prometheus, opts metav1.UpdateOptions) (*v1.Prometheus, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Prometheus, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PrometheusList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Prometheus, err error)
	PrometheusExpansion
}

// prometheuses implements PrometheusInterface
type prometheuses struct {
	client rest.Interface
	ns     string
}

// newPrometheuses returns a Prometheuses
func newPrometheuses(c *MonitoringV1Client, namespace string) *prometheuses {
	return &prometheuses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the prometheus, and returns the corresponding prometheus object, and an error if there is any.
func (c *prometheuses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Prometheus, err error) {
	result = &v1.Prometheus{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("prometheuses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Prometheuses that match those selectors.
func (c *prometheuses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PrometheusList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PrometheusList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("prometheuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested prometheuses.
func (c *prometheuses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("prometheuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a prometheus and creates it.  Returns the server's representation of the prometheus, and an error, if there is any.
func (c *prometheuses) Create(ctx context.Context, prometheus *v1.Prometheus, opts metav1.CreateOptions) (result *v1.Prometheus, err error) {
	result = &v1.Prometheus{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("prometheuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(prometheus).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a prometheus and updates it. Returns the server's representation of the prometheus, and an error, if there is any.
func (c *prometheuses) Update(ctx context.Context, prometheus *v1.Prometheus, opts metav1.UpdateOptions) (result *v1.Prometheus, err error) {
	result = &v1.Prometheus{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("prometheuses").
		Name(prometheus.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(prometheus).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *prometheuses) UpdateStatus(ctx context.Context, prometheus *v1.Prometheus, opts metav1.UpdateOptions) (result *v1.Prometheus, err error) {
	result = &v1.Prometheus{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("prometheuses").
		Name(prometheus.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(prometheus).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the prometheus and deletes it. Returns an error if one occurs.
func (c *prometheuses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("prometheuses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *prometheuses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("prometheuses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched prometheus.
func (c *prometheuses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Prometheus, err error) {
	result = &v1.Prometheus{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("prometheuses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
