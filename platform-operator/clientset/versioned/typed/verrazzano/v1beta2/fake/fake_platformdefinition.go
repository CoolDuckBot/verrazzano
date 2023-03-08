// Copyright (c) 2020, 2023, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta2 "github.com/verrazzano/verrazzano/platform-operator/apis/verrazzano/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePlatformDefinitions implements PlatformDefinitionInterface
type FakePlatformDefinitions struct {
	Fake *FakeVerrazzanoV1beta2
	ns   string
}

var platformdefinitionsResource = schema.GroupVersionResource{Group: "verrazzano", Version: "v1beta2", Resource: "platformdefinitions"}

var platformdefinitionsKind = schema.GroupVersionKind{Group: "verrazzano", Version: "v1beta2", Kind: "PlatformDefinition"}

// Get takes name of the platformDefinition, and returns the corresponding platformDefinition object, and an error if there is any.
func (c *FakePlatformDefinitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.PlatformDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(platformdefinitionsResource, c.ns, name), &v1beta2.PlatformDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.PlatformDefinition), err
}

// List takes label and field selectors, and returns the list of PlatformDefinitions that match those selectors.
func (c *FakePlatformDefinitions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.PlatformDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(platformdefinitionsResource, platformdefinitionsKind, c.ns, opts), &v1beta2.PlatformDefinitionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.PlatformDefinitionList{ListMeta: obj.(*v1beta2.PlatformDefinitionList).ListMeta}
	for _, item := range obj.(*v1beta2.PlatformDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested platformDefinitions.
func (c *FakePlatformDefinitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(platformdefinitionsResource, c.ns, opts))

}

// Create takes the representation of a platformDefinition and creates it.  Returns the server's representation of the platformDefinition, and an error, if there is any.
func (c *FakePlatformDefinitions) Create(ctx context.Context, platformDefinition *v1beta2.PlatformDefinition, opts v1.CreateOptions) (result *v1beta2.PlatformDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(platformdefinitionsResource, c.ns, platformDefinition), &v1beta2.PlatformDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.PlatformDefinition), err
}

// Update takes the representation of a platformDefinition and updates it. Returns the server's representation of the platformDefinition, and an error, if there is any.
func (c *FakePlatformDefinitions) Update(ctx context.Context, platformDefinition *v1beta2.PlatformDefinition, opts v1.UpdateOptions) (result *v1beta2.PlatformDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(platformdefinitionsResource, c.ns, platformDefinition), &v1beta2.PlatformDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.PlatformDefinition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePlatformDefinitions) UpdateStatus(ctx context.Context, platformDefinition *v1beta2.PlatformDefinition, opts v1.UpdateOptions) (*v1beta2.PlatformDefinition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(platformdefinitionsResource, "status", c.ns, platformDefinition), &v1beta2.PlatformDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.PlatformDefinition), err
}

// Delete takes name of the platformDefinition and deletes it. Returns an error if one occurs.
func (c *FakePlatformDefinitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(platformdefinitionsResource, c.ns, name, opts), &v1beta2.PlatformDefinition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePlatformDefinitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(platformdefinitionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.PlatformDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched platformDefinition.
func (c *FakePlatformDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.PlatformDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(platformdefinitionsResource, c.ns, name, pt, data, subresources...), &v1beta2.PlatformDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.PlatformDefinition), err
}