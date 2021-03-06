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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/xigang/kube-crd-controller/pkg/apis/example.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MyCustomResourceLister helps list MyCustomResources.
type MyCustomResourceLister interface {
	// List lists all MyCustomResources in the indexer.
	List(selector labels.Selector) (ret []*v1.MyCustomResource, err error)
	// MyCustomResources returns an object that can list and get MyCustomResources.
	MyCustomResources(namespace string) MyCustomResourceNamespaceLister
	MyCustomResourceListerExpansion
}

// myCustomResourceLister implements the MyCustomResourceLister interface.
type myCustomResourceLister struct {
	indexer cache.Indexer
}

// NewMyCustomResourceLister returns a new MyCustomResourceLister.
func NewMyCustomResourceLister(indexer cache.Indexer) MyCustomResourceLister {
	return &myCustomResourceLister{indexer: indexer}
}

// List lists all MyCustomResources in the indexer.
func (s *myCustomResourceLister) List(selector labels.Selector) (ret []*v1.MyCustomResource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MyCustomResource))
	})
	return ret, err
}

// MyCustomResources returns an object that can list and get MyCustomResources.
func (s *myCustomResourceLister) MyCustomResources(namespace string) MyCustomResourceNamespaceLister {
	return myCustomResourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MyCustomResourceNamespaceLister helps list and get MyCustomResources.
type MyCustomResourceNamespaceLister interface {
	// List lists all MyCustomResources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.MyCustomResource, err error)
	// Get retrieves the MyCustomResource from the indexer for a given namespace and name.
	Get(name string) (*v1.MyCustomResource, error)
	MyCustomResourceNamespaceListerExpansion
}

// myCustomResourceNamespaceLister implements the MyCustomResourceNamespaceLister
// interface.
type myCustomResourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MyCustomResources in the indexer for a given namespace.
func (s myCustomResourceNamespaceLister) List(selector labels.Selector) (ret []*v1.MyCustomResource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MyCustomResource))
	})
	return ret, err
}

// Get retrieves the MyCustomResource from the indexer for a given namespace and name.
func (s myCustomResourceNamespaceLister) Get(name string) (*v1.MyCustomResource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("mycustomresource"), name)
	}
	return obj.(*v1.MyCustomResource), nil
}
