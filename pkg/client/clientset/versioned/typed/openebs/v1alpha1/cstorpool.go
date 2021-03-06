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
package v1alpha1

import (
	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	scheme "github.com/openebs/maya/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CStorPoolsGetter has a method to return a CStorPoolInterface.
// A group's client should implement this interface.
type CStorPoolsGetter interface {
	CStorPools() CStorPoolInterface
}

// CStorPoolInterface has methods to work with CStorPool resources.
type CStorPoolInterface interface {
	Create(*v1alpha1.CStorPool) (*v1alpha1.CStorPool, error)
	Update(*v1alpha1.CStorPool) (*v1alpha1.CStorPool, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CStorPool, error)
	List(opts v1.ListOptions) (*v1alpha1.CStorPoolList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CStorPool, err error)
	CStorPoolExpansion
}

// cStorPools implements CStorPoolInterface
type cStorPools struct {
	client rest.Interface
}

// newCStorPools returns a CStorPools
func newCStorPools(c *OpenebsV1alpha1Client) *cStorPools {
	return &cStorPools{
		client: c.RESTClient(),
	}
}

// Get takes name of the cStorPool, and returns the corresponding cStorPool object, and an error if there is any.
func (c *cStorPools) Get(name string, options v1.GetOptions) (result *v1alpha1.CStorPool, err error) {
	result = &v1alpha1.CStorPool{}
	err = c.client.Get().
		Resource("cstorpools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CStorPools that match those selectors.
func (c *cStorPools) List(opts v1.ListOptions) (result *v1alpha1.CStorPoolList, err error) {
	result = &v1alpha1.CStorPoolList{}
	err = c.client.Get().
		Resource("cstorpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cStorPools.
func (c *cStorPools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("cstorpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a cStorPool and creates it.  Returns the server's representation of the cStorPool, and an error, if there is any.
func (c *cStorPools) Create(cStorPool *v1alpha1.CStorPool) (result *v1alpha1.CStorPool, err error) {
	result = &v1alpha1.CStorPool{}
	err = c.client.Post().
		Resource("cstorpools").
		Body(cStorPool).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cStorPool and updates it. Returns the server's representation of the cStorPool, and an error, if there is any.
func (c *cStorPools) Update(cStorPool *v1alpha1.CStorPool) (result *v1alpha1.CStorPool, err error) {
	result = &v1alpha1.CStorPool{}
	err = c.client.Put().
		Resource("cstorpools").
		Name(cStorPool.Name).
		Body(cStorPool).
		Do().
		Into(result)
	return
}

// Delete takes name of the cStorPool and deletes it. Returns an error if one occurs.
func (c *cStorPools) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("cstorpools").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cStorPools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("cstorpools").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cStorPool.
func (c *cStorPools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CStorPool, err error) {
	result = &v1alpha1.CStorPool{}
	err = c.client.Patch(pt).
		Resource("cstorpools").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
