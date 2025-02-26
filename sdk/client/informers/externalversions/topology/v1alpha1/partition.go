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

	topologyv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/topology/v1alpha1"
	scopedclientset "github.com/kcp-dev/kcp/sdk/client/clientset/versioned"
	clientset "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/cluster"
	"github.com/kcp-dev/kcp/sdk/client/informers/externalversions/internalinterfaces"
	topologyv1alpha1listers "github.com/kcp-dev/kcp/sdk/client/listers/topology/v1alpha1"
)

// PartitionClusterInformer provides access to a shared informer and lister for
// Partitions.
type PartitionClusterInformer interface {
	Cluster(logicalcluster.Name) PartitionInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() topologyv1alpha1listers.PartitionClusterLister
}

type partitionClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPartitionClusterInformer constructs a new informer for Partition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPartitionClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredPartitionClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPartitionClusterInformer constructs a new informer for Partition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPartitionClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TopologyV1alpha1().Partitions().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TopologyV1alpha1().Partitions().Watch(context.TODO(), options)
			},
		},
		&topologyv1alpha1.Partition{},
		resyncPeriod,
		indexers,
	)
}

func (f *partitionClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredPartitionClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *partitionClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&topologyv1alpha1.Partition{}, f.defaultInformer)
}

func (f *partitionClusterInformer) Lister() topologyv1alpha1listers.PartitionClusterLister {
	return topologyv1alpha1listers.NewPartitionClusterLister(f.Informer().GetIndexer())
}

// PartitionInformer provides access to a shared informer and lister for
// Partitions.
type PartitionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() topologyv1alpha1listers.PartitionLister
}

func (f *partitionClusterInformer) Cluster(clusterName logicalcluster.Name) PartitionInformer {
	return &partitionInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type partitionInformer struct {
	informer cache.SharedIndexInformer
	lister   topologyv1alpha1listers.PartitionLister
}

func (f *partitionInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *partitionInformer) Lister() topologyv1alpha1listers.PartitionLister {
	return f.lister
}

type partitionScopedInformer struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func (f *partitionScopedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&topologyv1alpha1.Partition{}, f.defaultInformer)
}

func (f *partitionScopedInformer) Lister() topologyv1alpha1listers.PartitionLister {
	return topologyv1alpha1listers.NewPartitionLister(f.Informer().GetIndexer())
}

// NewPartitionInformer constructs a new informer for Partition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPartitionInformer(client scopedclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPartitionInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPartitionInformer constructs a new informer for Partition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPartitionInformer(client scopedclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TopologyV1alpha1().Partitions().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TopologyV1alpha1().Partitions().Watch(context.TODO(), options)
			},
		},
		&topologyv1alpha1.Partition{},
		resyncPeriod,
		indexers,
	)
}

func (f *partitionScopedInformer) defaultInformer(client scopedclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPartitionInformer(client, resyncPeriod, cache.Indexers{}, f.tweakListOptions)
}
