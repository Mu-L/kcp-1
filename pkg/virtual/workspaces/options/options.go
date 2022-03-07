/*
Copyright 2021 The KCP Authors.

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

package options

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"

	"github.com/kcp-dev/kcp/pkg/apis/tenancy/v1alpha1/helper"
	kcpclient "github.com/kcp-dev/kcp/pkg/client/clientset/versioned"
	kcpinformer "github.com/kcp-dev/kcp/pkg/client/informers/externalversions"
	"github.com/kcp-dev/kcp/pkg/virtual/framework"
	"github.com/kcp-dev/kcp/pkg/virtual/framework/rootapiserver"
	"github.com/kcp-dev/kcp/pkg/virtual/workspaces/builder"
)

const DefaultRootPathPrefix string = "/services/workspaces"

type Workspaces struct {
	RootPathPrefix string
}

func NewWorkspaces() *Workspaces {
	return &Workspaces{
		RootPathPrefix: DefaultRootPathPrefix,
	}
}

func (o *Workspaces) AddFlags(flags *pflag.FlagSet, prefix string) {
	if o == nil {
		return
	}

	flags.StringVar(&o.RootPathPrefix, prefix+"workspaces-base-path", o.RootPathPrefix, ""+
		"The prefix of the workspaces API server root path.\n"+
		"The final workspaces API root path will be of the form:\n    <root-path-prefix>/<org-name>/personal|all")
}

func (o *Workspaces) Validate(flagPrefix string) []error {
	if o == nil {
		return nil
	}
	errs := []error{}

	if !strings.HasPrefix(o.RootPathPrefix, "/") {
		errs = append(errs, fmt.Errorf("--%s-workspaces-base-path %v should start with /", flagPrefix, o.RootPathPrefix))
	}

	return errs
}

func (o *Workspaces) NewVirtualWorkspaces(
	kubeClusterClient kubernetes.ClusterInterface,
	kcpClusterClient kcpclient.ClusterInterface,
	wildcardKubeInformers informers.SharedInformerFactory,
	wildcardKcpInformers kcpinformer.SharedInformerFactory,
) (extraInformers []rootapiserver.InformerStart, workspaces []framework.VirtualWorkspace, err error) {
	rootKubeClient := kubeClusterClient.Cluster(helper.RootCluster)
	rootKcpClient := kcpClusterClient.Cluster(helper.RootCluster)

	virtualWorkspaces := []framework.VirtualWorkspace{
		builder.BuildVirtualWorkspace(o.RootPathPrefix, wildcardKcpInformers.Tenancy().V1alpha1().ClusterWorkspaces(), wildcardKubeInformers.Rbac().V1(), rootKcpClient, rootKubeClient, kcpClusterClient, kubeClusterClient),
	}
	return nil, virtualWorkspaces, nil
}
