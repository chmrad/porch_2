// Copyright 2023-2024 The kpt and Nephio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/nephio-project/porch/api/generated/clientset/versioned"
	internalinterfaces "github.com/nephio-project/porch/api/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/nephio-project/porch/api/generated/listers/porch/v1alpha1"
	porchv1alpha1 "github.com/nephio-project/porch/api/porch/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PorchPackageInformer provides access to a shared informer and lister for
// PorchPackages.
type PorchPackageInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PorchPackageLister
}

type porchPackageInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPorchPackageInformer constructs a new informer for PorchPackage type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPorchPackageInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPorchPackageInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPorchPackageInformer constructs a new informer for PorchPackage type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPorchPackageInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PorchV1alpha1().PorchPackages(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PorchV1alpha1().PorchPackages(namespace).Watch(context.TODO(), options)
			},
		},
		&porchv1alpha1.PorchPackage{},
		resyncPeriod,
		indexers,
	)
}

func (f *porchPackageInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPorchPackageInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *porchPackageInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&porchv1alpha1.PorchPackage{}, f.defaultInformer)
}

func (f *porchPackageInformer) Lister() v1alpha1.PorchPackageLister {
	return v1alpha1.NewPorchPackageLister(f.Informer().GetIndexer())
}
