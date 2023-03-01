// Copyright (c) 2021, 2023, Oracle and/or its affiliates.
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

// FakeVerrazzanoProjects implements VerrazzanoProjectInterface
type FakeVerrazzanoProjects struct {
	Fake *FakeClustersV1alpha1
	ns   string
}

var verrazzanoprojectsResource = schema.GroupVersionResource{Group: "clusters.verrazzano.io", Version: "v1alpha1", Resource: "verrazzanoprojects"}

var verrazzanoprojectsKind = schema.GroupVersionKind{Group: "clusters.verrazzano.io", Version: "v1alpha1", Kind: "VerrazzanoProject"}

// Get takes name of the verrazzanoProject, and returns the corresponding verrazzanoProject object, and an error if there is any.
func (c *FakeVerrazzanoProjects) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VerrazzanoProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(verrazzanoprojectsResource, c.ns, name), &v1alpha1.VerrazzanoProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VerrazzanoProject), err
}

// List takes label and field selectors, and returns the list of VerrazzanoProjects that match those selectors.
func (c *FakeVerrazzanoProjects) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VerrazzanoProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(verrazzanoprojectsResource, verrazzanoprojectsKind, c.ns, opts), &v1alpha1.VerrazzanoProjectList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VerrazzanoProjectList{ListMeta: obj.(*v1alpha1.VerrazzanoProjectList).ListMeta}
	for _, item := range obj.(*v1alpha1.VerrazzanoProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested verrazzanoProjects.
func (c *FakeVerrazzanoProjects) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(verrazzanoprojectsResource, c.ns, opts))

}

// Create takes the representation of a verrazzanoProject and creates it.  Returns the server's representation of the verrazzanoProject, and an error, if there is any.
func (c *FakeVerrazzanoProjects) Create(ctx context.Context, verrazzanoProject *v1alpha1.VerrazzanoProject, opts v1.CreateOptions) (result *v1alpha1.VerrazzanoProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(verrazzanoprojectsResource, c.ns, verrazzanoProject), &v1alpha1.VerrazzanoProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VerrazzanoProject), err
}

// Update takes the representation of a verrazzanoProject and updates it. Returns the server's representation of the verrazzanoProject, and an error, if there is any.
func (c *FakeVerrazzanoProjects) Update(ctx context.Context, verrazzanoProject *v1alpha1.VerrazzanoProject, opts v1.UpdateOptions) (result *v1alpha1.VerrazzanoProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(verrazzanoprojectsResource, c.ns, verrazzanoProject), &v1alpha1.VerrazzanoProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VerrazzanoProject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVerrazzanoProjects) UpdateStatus(ctx context.Context, verrazzanoProject *v1alpha1.VerrazzanoProject, opts v1.UpdateOptions) (*v1alpha1.VerrazzanoProject, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(verrazzanoprojectsResource, "status", c.ns, verrazzanoProject), &v1alpha1.VerrazzanoProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VerrazzanoProject), err
}

// Delete takes name of the verrazzanoProject and deletes it. Returns an error if one occurs.
func (c *FakeVerrazzanoProjects) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(verrazzanoprojectsResource, c.ns, name, opts), &v1alpha1.VerrazzanoProject{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVerrazzanoProjects) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(verrazzanoprojectsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VerrazzanoProjectList{})
	return err
}

// Patch applies the patch and returns the patched verrazzanoProject.
func (c *FakeVerrazzanoProjects) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VerrazzanoProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(verrazzanoprojectsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VerrazzanoProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VerrazzanoProject), err
}
