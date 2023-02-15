//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	apisv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/apis/v1alpha1"
	scopedclientset "github.com/kcp-dev/kcp/pkg/client/clientset/versioned"
	clientset "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/cluster"
	"github.com/kcp-dev/kcp/pkg/client/informers/externalversions/internalinterfaces"
	apisv1alpha1listers "github.com/kcp-dev/kcp/pkg/client/listers/apis/v1alpha1"
)

// APIExportClusterInformer provides access to a shared informer and lister for
// APIExports.
type APIExportClusterInformer interface {
	Cluster(logicalcluster.Name) APIExportInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() apisv1alpha1listers.APIExportClusterLister
}

type aPIExportClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAPIExportClusterInformer constructs a new informer for APIExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAPIExportClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredAPIExportClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAPIExportClusterInformer constructs a new informer for APIExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAPIExportClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApisV1alpha1().APIExports().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApisV1alpha1().APIExports().Watch(context.TODO(), options)
			},
		},
		&apisv1alpha1.APIExport{},
		resyncPeriod,
		indexers,
	)
}

func (f *aPIExportClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredAPIExportClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *aPIExportClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&apisv1alpha1.APIExport{}, f.defaultInformer)
}

func (f *aPIExportClusterInformer) Lister() apisv1alpha1listers.APIExportClusterLister {
	return apisv1alpha1listers.NewAPIExportClusterLister(f.Informer().GetIndexer())
}

// APIExportInformer provides access to a shared informer and lister for
// APIExports.
type APIExportInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() apisv1alpha1listers.APIExportLister
}

func (f *aPIExportClusterInformer) Cluster(clusterName logicalcluster.Name) APIExportInformer {
	return &aPIExportInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type aPIExportInformer struct {
	informer cache.SharedIndexInformer
	lister   apisv1alpha1listers.APIExportLister
}

func (f *aPIExportInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *aPIExportInformer) Lister() apisv1alpha1listers.APIExportLister {
	return f.lister
}

type aPIExportScopedInformer struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func (f *aPIExportScopedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisv1alpha1.APIExport{}, f.defaultInformer)
}

func (f *aPIExportScopedInformer) Lister() apisv1alpha1listers.APIExportLister {
	return apisv1alpha1listers.NewAPIExportLister(f.Informer().GetIndexer())
}

// NewAPIExportInformer constructs a new informer for APIExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAPIExportInformer(client scopedclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAPIExportInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAPIExportInformer constructs a new informer for APIExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAPIExportInformer(client scopedclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApisV1alpha1().APIExports().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApisV1alpha1().APIExports().Watch(context.TODO(), options)
			},
		},
		&apisv1alpha1.APIExport{},
		resyncPeriod,
		indexers,
	)
}

func (f *aPIExportScopedInformer) defaultInformer(client scopedclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAPIExportInformer(client, resyncPeriod, cache.Indexers{}, f.tweakListOptions)
}