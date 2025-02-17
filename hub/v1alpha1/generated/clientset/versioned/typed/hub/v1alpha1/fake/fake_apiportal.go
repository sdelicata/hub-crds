/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/traefik/hub-crds/hub/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAPIPortals implements APIPortalInterface
type FakeAPIPortals struct {
	Fake *FakeHubV1alpha1
}

var apiportalsResource = schema.GroupVersionResource{Group: "hub.traefik.io", Version: "v1alpha1", Resource: "apiportals"}

var apiportalsKind = schema.GroupVersionKind{Group: "hub.traefik.io", Version: "v1alpha1", Kind: "APIPortal"}

// Get takes name of the aPIPortal, and returns the corresponding aPIPortal object, and an error if there is any.
func (c *FakeAPIPortals) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIPortal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(apiportalsResource, name), &v1alpha1.APIPortal{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIPortal), err
}

// List takes label and field selectors, and returns the list of APIPortals that match those selectors.
func (c *FakeAPIPortals) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIPortalList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(apiportalsResource, apiportalsKind, opts), &v1alpha1.APIPortalList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.APIPortalList{ListMeta: obj.(*v1alpha1.APIPortalList).ListMeta}
	for _, item := range obj.(*v1alpha1.APIPortalList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aPIPortals.
func (c *FakeAPIPortals) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(apiportalsResource, opts))
}

// Create takes the representation of a aPIPortal and creates it.  Returns the server's representation of the aPIPortal, and an error, if there is any.
func (c *FakeAPIPortals) Create(ctx context.Context, aPIPortal *v1alpha1.APIPortal, opts v1.CreateOptions) (result *v1alpha1.APIPortal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(apiportalsResource, aPIPortal), &v1alpha1.APIPortal{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIPortal), err
}

// Update takes the representation of a aPIPortal and updates it. Returns the server's representation of the aPIPortal, and an error, if there is any.
func (c *FakeAPIPortals) Update(ctx context.Context, aPIPortal *v1alpha1.APIPortal, opts v1.UpdateOptions) (result *v1alpha1.APIPortal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(apiportalsResource, aPIPortal), &v1alpha1.APIPortal{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIPortal), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAPIPortals) UpdateStatus(ctx context.Context, aPIPortal *v1alpha1.APIPortal, opts v1.UpdateOptions) (*v1alpha1.APIPortal, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(apiportalsResource, "status", aPIPortal), &v1alpha1.APIPortal{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIPortal), err
}

// Delete takes name of the aPIPortal and deletes it. Returns an error if one occurs.
func (c *FakeAPIPortals) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(apiportalsResource, name), &v1alpha1.APIPortal{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAPIPortals) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(apiportalsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.APIPortalList{})
	return err
}

// Patch applies the patch and returns the patched aPIPortal.
func (c *FakeAPIPortals) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIPortal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(apiportalsResource, name, pt, data, subresources...), &v1alpha1.APIPortal{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIPortal), err
}
