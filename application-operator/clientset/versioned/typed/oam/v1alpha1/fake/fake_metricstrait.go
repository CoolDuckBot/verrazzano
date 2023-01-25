// Copyright (c) 2021, 2023, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/verrazzano/verrazzano/application-operator/apis/oam/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMetricsTraits implements MetricsTraitInterface
type FakeMetricsTraits struct {
	Fake *FakeOamV1alpha1
	ns   string
}

var metricstraitsResource = schema.GroupVersionResource{Group: "oam.verrazzano.io", Version: "v1alpha1", Resource: "metricstraits"}

var metricstraitsKind = schema.GroupVersionKind{Group: "oam.verrazzano.io", Version: "v1alpha1", Kind: "MetricsTrait"}

// Get takes name of the metricsTrait, and returns the corresponding metricsTrait object, and an error if there is any.
func (c *FakeMetricsTraits) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MetricsTrait, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(metricstraitsResource, c.ns, name), &v1alpha1.MetricsTrait{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricsTrait), err
}

// List takes label and field selectors, and returns the list of MetricsTraits that match those selectors.
func (c *FakeMetricsTraits) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MetricsTraitList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(metricstraitsResource, metricstraitsKind, c.ns, opts), &v1alpha1.MetricsTraitList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MetricsTraitList{ListMeta: obj.(*v1alpha1.MetricsTraitList).ListMeta}
	for _, item := range obj.(*v1alpha1.MetricsTraitList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested metricsTraits.
func (c *FakeMetricsTraits) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(metricstraitsResource, c.ns, opts))

}

// Create takes the representation of a metricsTrait and creates it.  Returns the server's representation of the metricsTrait, and an error, if there is any.
func (c *FakeMetricsTraits) Create(ctx context.Context, metricsTrait *v1alpha1.MetricsTrait, opts v1.CreateOptions) (result *v1alpha1.MetricsTrait, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(metricstraitsResource, c.ns, metricsTrait), &v1alpha1.MetricsTrait{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricsTrait), err
}

// Update takes the representation of a metricsTrait and updates it. Returns the server's representation of the metricsTrait, and an error, if there is any.
func (c *FakeMetricsTraits) Update(ctx context.Context, metricsTrait *v1alpha1.MetricsTrait, opts v1.UpdateOptions) (result *v1alpha1.MetricsTrait, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(metricstraitsResource, c.ns, metricsTrait), &v1alpha1.MetricsTrait{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricsTrait), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMetricsTraits) UpdateStatus(ctx context.Context, metricsTrait *v1alpha1.MetricsTrait, opts v1.UpdateOptions) (*v1alpha1.MetricsTrait, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(metricstraitsResource, "status", c.ns, metricsTrait), &v1alpha1.MetricsTrait{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricsTrait), err
}

// Delete takes name of the metricsTrait and deletes it. Returns an error if one occurs.
func (c *FakeMetricsTraits) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(metricstraitsResource, c.ns, name, opts), &v1alpha1.MetricsTrait{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMetricsTraits) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(metricstraitsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MetricsTraitList{})
	return err
}

// Patch applies the patch and returns the patched metricsTrait.
func (c *FakeMetricsTraits) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MetricsTrait, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(metricstraitsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MetricsTrait{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricsTrait), err
}
