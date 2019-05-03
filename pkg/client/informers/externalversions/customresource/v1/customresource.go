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

	customresourcev1 "github.com/andy2046/k8s-custom-resource-watch/pkg/apis/customresource/v1"
	versioned "github.com/andy2046/k8s-custom-resource-watch/pkg/client/clientset/versioned"
	internalinterfaces "github.com/andy2046/k8s-custom-resource-watch/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/andy2046/k8s-custom-resource-watch/pkg/client/listers/customresource/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CustomResourceInformer provides access to a shared informer and lister for
// CustomResources.
type CustomResourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CustomResourceLister
}

type customResourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCustomResourceInformer constructs a new informer for CustomResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCustomResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCustomResourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCustomResourceInformer constructs a new informer for CustomResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCustomResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NokubeV1().CustomResources(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NokubeV1().CustomResources(namespace).Watch(options)
			},
		},
		&customresourcev1.CustomResource{},
		resyncPeriod,
		indexers,
	)
}

func (f *customResourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCustomResourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *customResourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&customresourcev1.CustomResource{}, f.defaultInformer)
}

func (f *customResourceInformer) Lister() v1.CustomResourceLister {
	return v1.NewCustomResourceLister(f.Informer().GetIndexer())
}
