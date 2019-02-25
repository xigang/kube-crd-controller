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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	examplecomv1 "github.com/xigang/kube-crd-controller/pkg/apis/example.com/v1"
	versioned "github.com/xigang/kube-crd-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/xigang/kube-crd-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/xigang/kube-crd-controller/pkg/client/listers/example.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MyCustomResourceInformer provides access to a shared informer and lister for
// MyCustomResources.
type MyCustomResourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.MyCustomResourceLister
}

type myCustomResourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMyCustomResourceInformer constructs a new informer for MyCustomResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMyCustomResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMyCustomResourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMyCustomResourceInformer constructs a new informer for MyCustomResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMyCustomResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1().MyCustomResources(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1().MyCustomResources(namespace).Watch(options)
			},
		},
		&examplecomv1.MyCustomResource{},
		resyncPeriod,
		indexers,
	)
}

func (f *myCustomResourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMyCustomResourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *myCustomResourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&examplecomv1.MyCustomResource{}, f.defaultInformer)
}

func (f *myCustomResourceInformer) Lister() v1.MyCustomResourceLister {
	return v1.NewMyCustomResourceLister(f.Informer().GetIndexer())
}
