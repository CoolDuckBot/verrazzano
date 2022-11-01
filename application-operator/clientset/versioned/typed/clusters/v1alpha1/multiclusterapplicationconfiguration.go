// Copyright (c) 2021, 2022, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/verrazzano/verrazzano/application-operator/apis/clusters/v1alpha1"
	scheme "github.com/verrazzano/verrazzano/application-operator/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MultiClusterApplicationConfigurationsGetter has a method to return a MultiClusterApplicationConfigurationInterface.
// A group's client should implement this interface.
type MultiClusterApplicationConfigurationsGetter interface {
	MultiClusterApplicationConfigurations(namespace string) MultiClusterApplicationConfigurationInterface
}

// MultiClusterApplicationConfigurationInterface has methods to work with MultiClusterApplicationConfiguration resources.
type MultiClusterApplicationConfigurationInterface interface {
	Create(ctx context.Context, multiClusterApplicationConfiguration *v1alpha1.MultiClusterApplicationConfiguration, opts v1.CreateOptions) (*v1alpha1.MultiClusterApplicationConfiguration, error)
	Update(ctx context.Context, multiClusterApplicationConfiguration *v1alpha1.MultiClusterApplicationConfiguration, opts v1.UpdateOptions) (*v1alpha1.MultiClusterApplicationConfiguration, error)
	UpdateStatus(ctx context.Context, multiClusterApplicationConfiguration *v1alpha1.MultiClusterApplicationConfiguration, opts v1.UpdateOptions) (*v1alpha1.MultiClusterApplicationConfiguration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.MultiClusterApplicationConfiguration, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MultiClusterApplicationConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MultiClusterApplicationConfiguration, err error)
	MultiClusterApplicationConfigurationExpansion
}

// multiClusterApplicationConfigurations implements MultiClusterApplicationConfigurationInterface
type multiClusterApplicationConfigurations struct {
	client rest.Interface
	ns     string
}

// newMultiClusterApplicationConfigurations returns a MultiClusterApplicationConfigurations
func newMultiClusterApplicationConfigurations(c *ClustersV1alpha1Client, namespace string) *multiClusterApplicationConfigurations {
	return &multiClusterApplicationConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the multiClusterApplicationConfiguration, and returns the corresponding multiClusterApplicationConfiguration object, and an error if there is any.
func (c *multiClusterApplicationConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MultiClusterApplicationConfiguration, err error) {
	result = &v1alpha1.MultiClusterApplicationConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("multiclusterapplicationconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MultiClusterApplicationConfigurations that match those selectors.
func (c *multiClusterApplicationConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MultiClusterApplicationConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MultiClusterApplicationConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("multiclusterapplicationconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested multiClusterApplicationConfigurations.
func (c *multiClusterApplicationConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("multiclusterapplicationconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a multiClusterApplicationConfiguration and creates it.  Returns the server's representation of the multiClusterApplicationConfiguration, and an error, if there is any.
func (c *multiClusterApplicationConfigurations) Create(ctx context.Context, multiClusterApplicationConfiguration *v1alpha1.MultiClusterApplicationConfiguration, opts v1.CreateOptions) (result *v1alpha1.MultiClusterApplicationConfiguration, err error) {
	result = &v1alpha1.MultiClusterApplicationConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("multiclusterapplicationconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(multiClusterApplicationConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a multiClusterApplicationConfiguration and updates it. Returns the server's representation of the multiClusterApplicationConfiguration, and an error, if there is any.
func (c *multiClusterApplicationConfigurations) Update(ctx context.Context, multiClusterApplicationConfiguration *v1alpha1.MultiClusterApplicationConfiguration, opts v1.UpdateOptions) (result *v1alpha1.MultiClusterApplicationConfiguration, err error) {
	result = &v1alpha1.MultiClusterApplicationConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("multiclusterapplicationconfigurations").
		Name(multiClusterApplicationConfiguration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(multiClusterApplicationConfiguration).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *multiClusterApplicationConfigurations) UpdateStatus(ctx context.Context, multiClusterApplicationConfiguration *v1alpha1.MultiClusterApplicationConfiguration, opts v1.UpdateOptions) (result *v1alpha1.MultiClusterApplicationConfiguration, err error) {
	result = &v1alpha1.MultiClusterApplicationConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("multiclusterapplicationconfigurations").
		Name(multiClusterApplicationConfiguration.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(multiClusterApplicationConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the multiClusterApplicationConfiguration and deletes it. Returns an error if one occurs.
func (c *multiClusterApplicationConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("multiclusterapplicationconfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *multiClusterApplicationConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("multiclusterapplicationconfigurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched multiClusterApplicationConfiguration.
func (c *multiClusterApplicationConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MultiClusterApplicationConfiguration, err error) {
	result = &v1alpha1.MultiClusterApplicationConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("multiclusterapplicationconfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
