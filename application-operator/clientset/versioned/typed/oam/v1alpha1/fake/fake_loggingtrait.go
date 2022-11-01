// Copyright (c) 2021, 2022, Oracle and/or its affiliates.
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

// FakeLoggingTraits implements LoggingTraitInterface
type FakeLoggingTraits struct {
	Fake *FakeOamV1alpha1
	ns   string
}

var loggingtraitsResource = schema.GroupVersionResource{Group: "oam.verrazzano.io", Version: "v1alpha1", Resource: "loggingtraits"}

var loggingtraitsKind = schema.GroupVersionKind{Group: "oam.verrazzano.io", Version: "v1alpha1", Kind: "LoggingTrait"}

// Get takes name of the loggingTrait, and returns the corresponding loggingTrait object, and an error if there is any.
func (c *FakeLoggingTraits) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LoggingTrait, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(loggingtraitsResource, c.ns, name), &v1alpha1.LoggingTrait{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoggingTrait), err
}

// List takes label and field selectors, and returns the list of LoggingTraits that match those selectors.
func (c *FakeLoggingTraits) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LoggingTraitList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(loggingtraitsResource, loggingtraitsKind, c.ns, opts), &v1alpha1.LoggingTraitList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LoggingTraitList{ListMeta: obj.(*v1alpha1.LoggingTraitList).ListMeta}
	for _, item := range obj.(*v1alpha1.LoggingTraitList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested loggingTraits.
func (c *FakeLoggingTraits) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(loggingtraitsResource, c.ns, opts))

}

// Create takes the representation of a loggingTrait and creates it.  Returns the server's representation of the loggingTrait, and an error, if there is any.
func (c *FakeLoggingTraits) Create(ctx context.Context, loggingTrait *v1alpha1.LoggingTrait, opts v1.CreateOptions) (result *v1alpha1.LoggingTrait, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(loggingtraitsResource, c.ns, loggingTrait), &v1alpha1.LoggingTrait{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoggingTrait), err
}

// Update takes the representation of a loggingTrait and updates it. Returns the server's representation of the loggingTrait, and an error, if there is any.
func (c *FakeLoggingTraits) Update(ctx context.Context, loggingTrait *v1alpha1.LoggingTrait, opts v1.UpdateOptions) (result *v1alpha1.LoggingTrait, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(loggingtraitsResource, c.ns, loggingTrait), &v1alpha1.LoggingTrait{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoggingTrait), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLoggingTraits) UpdateStatus(ctx context.Context, loggingTrait *v1alpha1.LoggingTrait, opts v1.UpdateOptions) (*v1alpha1.LoggingTrait, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(loggingtraitsResource, "status", c.ns, loggingTrait), &v1alpha1.LoggingTrait{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoggingTrait), err
}

// Delete takes name of the loggingTrait and deletes it. Returns an error if one occurs.
func (c *FakeLoggingTraits) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(loggingtraitsResource, c.ns, name, opts), &v1alpha1.LoggingTrait{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLoggingTraits) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(loggingtraitsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LoggingTraitList{})
	return err
}

// Patch applies the patch and returns the patched loggingTrait.
func (c *FakeLoggingTraits) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LoggingTrait, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(loggingtraitsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LoggingTrait{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoggingTrait), err
}
