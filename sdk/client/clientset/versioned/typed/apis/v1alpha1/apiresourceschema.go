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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"

	v1alpha1 "github.com/kcp-dev/kcp/sdk/apis/apis/v1alpha1"
	apisv1alpha1 "github.com/kcp-dev/kcp/sdk/client/applyconfiguration/apis/v1alpha1"
	scheme "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/scheme"
)

// APIResourceSchemasGetter has a method to return a APIResourceSchemaInterface.
// A group's client should implement this interface.
type APIResourceSchemasGetter interface {
	APIResourceSchemas() APIResourceSchemaInterface
}

// APIResourceSchemaInterface has methods to work with APIResourceSchema resources.
type APIResourceSchemaInterface interface {
	Create(ctx context.Context, aPIResourceSchema *v1alpha1.APIResourceSchema, opts v1.CreateOptions) (*v1alpha1.APIResourceSchema, error)
	Update(ctx context.Context, aPIResourceSchema *v1alpha1.APIResourceSchema, opts v1.UpdateOptions) (*v1alpha1.APIResourceSchema, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.APIResourceSchema, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.APIResourceSchemaList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIResourceSchema, err error)
	Apply(ctx context.Context, aPIResourceSchema *apisv1alpha1.APIResourceSchemaApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.APIResourceSchema, err error)
	APIResourceSchemaExpansion
}

// aPIResourceSchemas implements APIResourceSchemaInterface
type aPIResourceSchemas struct {
	*gentype.ClientWithListAndApply[*v1alpha1.APIResourceSchema, *v1alpha1.APIResourceSchemaList, *apisv1alpha1.APIResourceSchemaApplyConfiguration]
}

// newAPIResourceSchemas returns a APIResourceSchemas
func newAPIResourceSchemas(c *ApisV1alpha1Client) *aPIResourceSchemas {
	return &aPIResourceSchemas{
		gentype.NewClientWithListAndApply[*v1alpha1.APIResourceSchema, *v1alpha1.APIResourceSchemaList, *apisv1alpha1.APIResourceSchemaApplyConfiguration](
			"apiresourceschemas",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.APIResourceSchema { return &v1alpha1.APIResourceSchema{} },
			func() *v1alpha1.APIResourceSchemaList { return &v1alpha1.APIResourceSchemaList{} }),
	}
}