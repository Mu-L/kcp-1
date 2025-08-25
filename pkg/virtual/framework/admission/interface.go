/*
Copyright 2025 The KCP Authors.

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

package admission

import (
	"context"

	"k8s.io/apiserver/pkg/admission"
)

/*
	These interfaces are supposed to be used by the virtual workspace admission plugins.
	We're intentionally not using the Kubernetes admission interfaces as those require an
	additional function, Handle, which we don't need in admission for virtual workspaces.
*/

// Mutator is an abstract, pluggable interface for Admission Control decisions.
type Mutator interface {
	// Admit makes an admission decision based on the request attributes.
	// Context is used only for timeout/deadline/cancellation and tracing information.
	Admit(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) (err error)
}

// Validator is an abstract, pluggable interface for Admission Control decisions.
type Validator interface {
	// Validate makes an admission decision based on the request attributes.  It is NOT allowed to mutate
	// Context is used only for timeout/deadline/cancellation and tracing information.
	Validate(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) (err error)
}

// MutatorValidator is a union of Mutator and Validator interfaces.
type MutatorValidator interface {
	Mutator
	Validator
}
