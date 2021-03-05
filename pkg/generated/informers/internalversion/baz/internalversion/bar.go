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

package internalversion

import (
	"context"
	time "time"

	baz "github.com/Marcos30004347/k8s-custom-API-Server/pkg/apis/baz"
	clientsetinternalversion "github.com/Marcos30004347/k8s-custom-API-Server/pkg/generated/clientset/internalversion"
	internalinterfaces "github.com/Marcos30004347/k8s-custom-API-Server/pkg/generated/informers/internalversion/internalinterfaces"
	internalversion "github.com/Marcos30004347/k8s-custom-API-Server/pkg/generated/listers/baz/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BarInformer provides access to a shared informer and lister for
// Bars.
type BarInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.BarLister
}

type barInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewBarInformer constructs a new informer for Bar type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBarInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBarInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredBarInformer constructs a new informer for Bar type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBarInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Baz().Bars().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Baz().Bars().Watch(context.TODO(), options)
			},
		},
		&baz.Bar{},
		resyncPeriod,
		indexers,
	)
}

func (f *barInformer) defaultInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBarInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *barInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&baz.Bar{}, f.defaultInformer)
}

func (f *barInformer) Lister() internalversion.BarLister {
	return internalversion.NewBarLister(f.Informer().GetIndexer())
}
