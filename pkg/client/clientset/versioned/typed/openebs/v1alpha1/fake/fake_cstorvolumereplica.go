/*
Copyright 2018 The OpenEBS Authors

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
package fake

import (
	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCStorVolumeReplicas implements CStorVolumeReplicaInterface
type FakeCStorVolumeReplicas struct {
	Fake *FakeOpenebsV1alpha1
}

var cstorvolumereplicasResource = schema.GroupVersionResource{Group: "openebs.io", Version: "v1alpha1", Resource: "cstorvolumereplicas"}

var cstorvolumereplicasKind = schema.GroupVersionKind{Group: "openebs.io", Version: "v1alpha1", Kind: "CStorVolumeReplica"}

// Get takes name of the cStorVolumeReplica, and returns the corresponding cStorVolumeReplica object, and an error if there is any.
func (c *FakeCStorVolumeReplicas) Get(name string, options v1.GetOptions) (result *v1alpha1.CStorVolumeReplica, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(cstorvolumereplicasResource, name), &v1alpha1.CStorVolumeReplica{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CStorVolumeReplica), err
}

// List takes label and field selectors, and returns the list of CStorVolumeReplicas that match those selectors.
func (c *FakeCStorVolumeReplicas) List(opts v1.ListOptions) (result *v1alpha1.CStorVolumeReplicaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(cstorvolumereplicasResource, cstorvolumereplicasKind, opts), &v1alpha1.CStorVolumeReplicaList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CStorVolumeReplicaList{}
	for _, item := range obj.(*v1alpha1.CStorVolumeReplicaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cStorVolumeReplicas.
func (c *FakeCStorVolumeReplicas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(cstorvolumereplicasResource, opts))
}

// Create takes the representation of a cStorVolumeReplica and creates it.  Returns the server's representation of the cStorVolumeReplica, and an error, if there is any.
func (c *FakeCStorVolumeReplicas) Create(cStorVolumeReplica *v1alpha1.CStorVolumeReplica) (result *v1alpha1.CStorVolumeReplica, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(cstorvolumereplicasResource, cStorVolumeReplica), &v1alpha1.CStorVolumeReplica{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CStorVolumeReplica), err
}

// Update takes the representation of a cStorVolumeReplica and updates it. Returns the server's representation of the cStorVolumeReplica, and an error, if there is any.
func (c *FakeCStorVolumeReplicas) Update(cStorVolumeReplica *v1alpha1.CStorVolumeReplica) (result *v1alpha1.CStorVolumeReplica, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(cstorvolumereplicasResource, cStorVolumeReplica), &v1alpha1.CStorVolumeReplica{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CStorVolumeReplica), err
}

// Delete takes name of the cStorVolumeReplica and deletes it. Returns an error if one occurs.
func (c *FakeCStorVolumeReplicas) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(cstorvolumereplicasResource, name), &v1alpha1.CStorVolumeReplica{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCStorVolumeReplicas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(cstorvolumereplicasResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CStorVolumeReplicaList{})
	return err
}

// Patch applies the patch and returns the patched cStorVolumeReplica.
func (c *FakeCStorVolumeReplicas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CStorVolumeReplica, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(cstorvolumereplicasResource, name, data, subresources...), &v1alpha1.CStorVolumeReplica{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CStorVolumeReplica), err
}
