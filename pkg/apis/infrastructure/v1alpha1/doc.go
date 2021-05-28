/*
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

// Package v1alpha1 contains API Schema definitions for the v1alpha1 API group
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:defaulter-gen=TypeMeta
// +groupName=kit.k8s.amazonaws.com
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"knative.dev/pkg/apis"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// APIVersion is the current API version used to register these objects
	APIVersion = "v1alpha1"

	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "kit.k8s.amazonaws.com", Version: APIVersion}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme is required by pkg/client/...
	AddToScheme = SchemeBuilder.AddToScheme
)

const (
	// Active is a condition implemented by all resources. It indicates that the
	// controller is able to take actions: it's correctly configured, can make
	// necessary API calls, and isn't disabled.
	Active apis.ConditionType = "Active"
)

func init() {
	SchemeBuilder.Register(&ControlPlane{}, &ControlPlaneList{})
	SchemeBuilder.Register(&VPC{}, &VPCList{})
	SchemeBuilder.Register(&Subnet{}, &SubnetList{})
	SchemeBuilder.Register(&InternetGateway{}, &InternetGatewayList{})
}
