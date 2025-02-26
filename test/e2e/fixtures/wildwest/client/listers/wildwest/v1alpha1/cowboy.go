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
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	wildwestv1alpha1 "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/apis/wildwest/v1alpha1"
)

// CowboyClusterLister can list Cowboys across all workspaces, or scope down to a CowboyLister for one workspace.
// All objects returned here must be treated as read-only.
type CowboyClusterLister interface {
	// List lists all Cowboys in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*wildwestv1alpha1.Cowboy, err error)
	// Cluster returns a lister that can list and get Cowboys in one workspace.
	Cluster(clusterName logicalcluster.Name) CowboyLister
	CowboyClusterListerExpansion
}

type cowboyClusterLister struct {
	indexer cache.Indexer
}

// NewCowboyClusterLister returns a new CowboyClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewCowboyClusterLister(indexer cache.Indexer) *cowboyClusterLister {
	return &cowboyClusterLister{indexer: indexer}
}

// List lists all Cowboys in the indexer across all workspaces.
func (s *cowboyClusterLister) List(selector labels.Selector) (ret []*wildwestv1alpha1.Cowboy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*wildwestv1alpha1.Cowboy))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get Cowboys.
func (s *cowboyClusterLister) Cluster(clusterName logicalcluster.Name) CowboyLister {
	return &cowboyLister{indexer: s.indexer, clusterName: clusterName}
}

// CowboyLister can list Cowboys across all namespaces, or scope down to a CowboyNamespaceLister for one namespace.
// All objects returned here must be treated as read-only.
type CowboyLister interface {
	// List lists all Cowboys in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*wildwestv1alpha1.Cowboy, err error)
	// Cowboys returns a lister that can list and get Cowboys in one workspace and namespace.
	Cowboys(namespace string) CowboyNamespaceLister
	CowboyListerExpansion
}

// cowboyLister can list all Cowboys inside a workspace or scope down to a CowboyLister for one namespace.
type cowboyLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all Cowboys in the indexer for a workspace.
func (s *cowboyLister) List(selector labels.Selector) (ret []*wildwestv1alpha1.Cowboy, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*wildwestv1alpha1.Cowboy))
	})
	return ret, err
}

// Cowboys returns an object that can list and get Cowboys in one namespace.
func (s *cowboyLister) Cowboys(namespace string) CowboyNamespaceLister {
	return &cowboyNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// cowboyNamespaceLister helps list and get Cowboys.
// All objects returned here must be treated as read-only.
type CowboyNamespaceLister interface {
	// List lists all Cowboys in the workspace and namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*wildwestv1alpha1.Cowboy, err error)
	// Get retrieves the Cowboy from the indexer for a given workspace, namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*wildwestv1alpha1.Cowboy, error)
	CowboyNamespaceListerExpansion
}

// cowboyNamespaceLister helps list and get Cowboys.
// All objects returned here must be treated as read-only.
type cowboyNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// List lists all Cowboys in the indexer for a given workspace and namespace.
func (s *cowboyNamespaceLister) List(selector labels.Selector) (ret []*wildwestv1alpha1.Cowboy, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*wildwestv1alpha1.Cowboy))
	})
	return ret, err
}

// Get retrieves the Cowboy from the indexer for a given workspace, namespace and name.
func (s *cowboyNamespaceLister) Get(name string) (*wildwestv1alpha1.Cowboy, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(wildwestv1alpha1.Resource("cowboys"), name)
	}
	return obj.(*wildwestv1alpha1.Cowboy), nil
}

// NewCowboyLister returns a new CowboyLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
// - has the cache.NamespaceIndex as an index
func NewCowboyLister(indexer cache.Indexer) *cowboyScopedLister {
	return &cowboyScopedLister{indexer: indexer}
}

// cowboyScopedLister can list all Cowboys inside a workspace or scope down to a CowboyLister for one namespace.
type cowboyScopedLister struct {
	indexer cache.Indexer
}

// List lists all Cowboys in the indexer for a workspace.
func (s *cowboyScopedLister) List(selector labels.Selector) (ret []*wildwestv1alpha1.Cowboy, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*wildwestv1alpha1.Cowboy))
	})
	return ret, err
}

// Cowboys returns an object that can list and get Cowboys in one namespace.
func (s *cowboyScopedLister) Cowboys(namespace string) CowboyNamespaceLister {
	return &cowboyScopedNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// cowboyScopedNamespaceLister helps list and get Cowboys.
type cowboyScopedNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Cowboys in the indexer for a given workspace and namespace.
func (s *cowboyScopedNamespaceLister) List(selector labels.Selector) (ret []*wildwestv1alpha1.Cowboy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*wildwestv1alpha1.Cowboy))
	})
	return ret, err
}

// Get retrieves the Cowboy from the indexer for a given workspace, namespace and name.
func (s *cowboyScopedNamespaceLister) Get(name string) (*wildwestv1alpha1.Cowboy, error) {
	key := s.namespace + "/" + name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(wildwestv1alpha1.Resource("cowboys"), name)
	}
	return obj.(*wildwestv1alpha1.Cowboy), nil
}
