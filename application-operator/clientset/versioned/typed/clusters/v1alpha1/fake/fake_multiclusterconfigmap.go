// Copyright (c) 2021, 2022, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/verrazzano/verrazzano/application-operator/apis/clusters/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMultiClusterConfigMaps implements MultiClusterConfigMapInterface
type FakeMultiClusterConfigMaps struct {
	Fake *FakeClustersV1alpha1
	ns   string
}

var multiclusterconfigmapsResource = schema.GroupVersionResource{Group: "clusters.verrazzano.io", Version: "v1alpha1", Resource: "multiclusterconfigmaps"}

var multiclusterconfigmapsKind = schema.GroupVersionKind{Group: "clusters.verrazzano.io", Version: "v1alpha1", Kind: "MultiClusterConfigMap"}

// Get takes name of the multiClusterConfigMap, and returns the corresponding multiClusterConfigMap object, and an error if there is any.
func (c *FakeMultiClusterConfigMaps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MultiClusterConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(multiclusterconfigmapsResource, c.ns, name), &v1alpha1.MultiClusterConfigMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterConfigMap), err
}

// List takes label and field selectors, and returns the list of MultiClusterConfigMaps that match those selectors.
func (c *FakeMultiClusterConfigMaps) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MultiClusterConfigMapList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(multiclusterconfigmapsResource, multiclusterconfigmapsKind, c.ns, opts), &v1alpha1.MultiClusterConfigMapList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MultiClusterConfigMapList{ListMeta: obj.(*v1alpha1.MultiClusterConfigMapList).ListMeta}
	for _, item := range obj.(*v1alpha1.MultiClusterConfigMapList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested multiClusterConfigMaps.
func (c *FakeMultiClusterConfigMaps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(multiclusterconfigmapsResource, c.ns, opts))

}

// Create takes the representation of a multiClusterConfigMap and creates it.  Returns the server's representation of the multiClusterConfigMap, and an error, if there is any.
func (c *FakeMultiClusterConfigMaps) Create(ctx context.Context, multiClusterConfigMap *v1alpha1.MultiClusterConfigMap, opts v1.CreateOptions) (result *v1alpha1.MultiClusterConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(multiclusterconfigmapsResource, c.ns, multiClusterConfigMap), &v1alpha1.MultiClusterConfigMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterConfigMap), err
}

// Update takes the representation of a multiClusterConfigMap and updates it. Returns the server's representation of the multiClusterConfigMap, and an error, if there is any.
func (c *FakeMultiClusterConfigMaps) Update(ctx context.Context, multiClusterConfigMap *v1alpha1.MultiClusterConfigMap, opts v1.UpdateOptions) (result *v1alpha1.MultiClusterConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(multiclusterconfigmapsResource, c.ns, multiClusterConfigMap), &v1alpha1.MultiClusterConfigMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterConfigMap), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMultiClusterConfigMaps) UpdateStatus(ctx context.Context, multiClusterConfigMap *v1alpha1.MultiClusterConfigMap, opts v1.UpdateOptions) (*v1alpha1.MultiClusterConfigMap, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(multiclusterconfigmapsResource, "status", c.ns, multiClusterConfigMap), &v1alpha1.MultiClusterConfigMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterConfigMap), err
}

// Delete takes name of the multiClusterConfigMap and deletes it. Returns an error if one occurs.
func (c *FakeMultiClusterConfigMaps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(multiclusterconfigmapsResource, c.ns, name, opts), &v1alpha1.MultiClusterConfigMap{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMultiClusterConfigMaps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(multiclusterconfigmapsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MultiClusterConfigMapList{})
	return err
}

// Patch applies the patch and returns the patched multiClusterConfigMap.
func (c *FakeMultiClusterConfigMaps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MultiClusterConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(multiclusterconfigmapsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MultiClusterConfigMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterConfigMap), err
}
