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

package v1

import (
	"time"

	v1 "github.com/xigang/kube-crd-controller/pkg/apis/example.com/v1"
	scheme "github.com/xigang/kube-crd-controller/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MyCustomResourcesGetter has a method to return a MyCustomResourceInterface.
// A group's client should implement this interface.
type MyCustomResourcesGetter interface {
	MyCustomResources(namespace string) MyCustomResourceInterface
}

// MyCustomResourceInterface has methods to work with MyCustomResource resources.
type MyCustomResourceInterface interface {
	Create(*v1.MyCustomResource) (*v1.MyCustomResource, error)
	Update(*v1.MyCustomResource) (*v1.MyCustomResource, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.MyCustomResource, error)
	List(opts metav1.ListOptions) (*v1.MyCustomResourceList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MyCustomResource, err error)
	MyCustomResourceExpansion
}

// myCustomResources implements MyCustomResourceInterface
type myCustomResources struct {
	client rest.Interface
	ns     string
}

// newMyCustomResources returns a MyCustomResources
func newMyCustomResources(c *ExampleV1Client, namespace string) *myCustomResources {
	return &myCustomResources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the myCustomResource, and returns the corresponding myCustomResource object, and an error if there is any.
func (c *myCustomResources) Get(name string, options metav1.GetOptions) (result *v1.MyCustomResource, err error) {
	result = &v1.MyCustomResource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mycustomresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MyCustomResources that match those selectors.
func (c *myCustomResources) List(opts metav1.ListOptions) (result *v1.MyCustomResourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MyCustomResourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mycustomresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested myCustomResources.
func (c *myCustomResources) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mycustomresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a myCustomResource and creates it.  Returns the server's representation of the myCustomResource, and an error, if there is any.
func (c *myCustomResources) Create(myCustomResource *v1.MyCustomResource) (result *v1.MyCustomResource, err error) {
	result = &v1.MyCustomResource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mycustomresources").
		Body(myCustomResource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a myCustomResource and updates it. Returns the server's representation of the myCustomResource, and an error, if there is any.
func (c *myCustomResources) Update(myCustomResource *v1.MyCustomResource) (result *v1.MyCustomResource, err error) {
	result = &v1.MyCustomResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mycustomresources").
		Name(myCustomResource.Name).
		Body(myCustomResource).
		Do().
		Into(result)
	return
}

// Delete takes name of the myCustomResource and deletes it. Returns an error if one occurs.
func (c *myCustomResources) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mycustomresources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *myCustomResources) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mycustomresources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched myCustomResource.
func (c *myCustomResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MyCustomResource, err error) {
	result = &v1.MyCustomResource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mycustomresources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
