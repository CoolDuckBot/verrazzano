// Copyright (c) 2021, 2023, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/verrazzano/verrazzano/application-operator/apis/app/v1alpha1"
	scheme "github.com/verrazzano/verrazzano/application-operator/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MetricsTemplatesGetter has a method to return a MetricsTemplateInterface.
// A group's client should implement this interface.
type MetricsTemplatesGetter interface {
	MetricsTemplates(namespace string) MetricsTemplateInterface
}

// MetricsTemplateInterface has methods to work with MetricsTemplate resources.
type MetricsTemplateInterface interface {
	Create(ctx context.Context, metricsTemplate *v1alpha1.MetricsTemplate, opts v1.CreateOptions) (*v1alpha1.MetricsTemplate, error)
	Update(ctx context.Context, metricsTemplate *v1alpha1.MetricsTemplate, opts v1.UpdateOptions) (*v1alpha1.MetricsTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.MetricsTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MetricsTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MetricsTemplate, err error)
	MetricsTemplateExpansion
}

// metricsTemplates implements MetricsTemplateInterface
type metricsTemplates struct {
	client rest.Interface
	ns     string
}

// newMetricsTemplates returns a MetricsTemplates
func newMetricsTemplates(c *AppV1alpha1Client, namespace string) *metricsTemplates {
	return &metricsTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the metricsTemplate, and returns the corresponding metricsTemplate object, and an error if there is any.
func (c *metricsTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MetricsTemplate, err error) {
	result = &v1alpha1.MetricsTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("metricstemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MetricsTemplates that match those selectors.
func (c *metricsTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MetricsTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MetricsTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("metricstemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested metricsTemplates.
func (c *metricsTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("metricstemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a metricsTemplate and creates it.  Returns the server's representation of the metricsTemplate, and an error, if there is any.
func (c *metricsTemplates) Create(ctx context.Context, metricsTemplate *v1alpha1.MetricsTemplate, opts v1.CreateOptions) (result *v1alpha1.MetricsTemplate, err error) {
	result = &v1alpha1.MetricsTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("metricstemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(metricsTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a metricsTemplate and updates it. Returns the server's representation of the metricsTemplate, and an error, if there is any.
func (c *metricsTemplates) Update(ctx context.Context, metricsTemplate *v1alpha1.MetricsTemplate, opts v1.UpdateOptions) (result *v1alpha1.MetricsTemplate, err error) {
	result = &v1alpha1.MetricsTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("metricstemplates").
		Name(metricsTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(metricsTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the metricsTemplate and deletes it. Returns an error if one occurs.
func (c *metricsTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("metricstemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *metricsTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("metricstemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched metricsTemplate.
func (c *metricsTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MetricsTemplate, err error) {
	result = &v1alpha1.MetricsTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("metricstemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}