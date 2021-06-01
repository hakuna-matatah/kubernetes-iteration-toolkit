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

package v1alpha1

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis"
)

// LaunchTemplate is the Schema for the LaunchTemplate API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type LaunchTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LaunchTemplateSpec   `json:"spec,omitempty"`
	Status LaunchTemplateStatus `json:"status,omitempty"`
}

// LaunchTemplateSpec
type LaunchTemplateSpec struct {
	ClusterName string `json:"clusterName,omitempty"`
}

// LaunchTemplateList contains a list of LaunchTemplate
// +kubebuilder:object:root=true
type LaunchTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LaunchTemplate `json:"items"`
}

// LaunchTemplateStatus defines the observed state of the LaunchTemplate of a cluster
type LaunchTemplateStatus struct {
	// Conditions is the set of conditions required for this LaunchTemplate to create
	// its objects, and indicates whether or not those conditions are met.
	// +optional
	Conditions apis.Conditions `json:"conditions,omitempty"`
}

func (c *LaunchTemplate) StatusConditions() apis.ConditionManager {
	return apis.NewLivingConditionSet(
		Active,
	).Manage(c)
}

func (c *LaunchTemplate) GetConditions() apis.Conditions {
	return c.Status.Conditions
}

func (c *LaunchTemplate) SetConditions(conditions apis.Conditions) {
	c.Status.Conditions = conditions
}

// <foo-master-instances>-<LaunchTemplate>
func LaunchTemplateName(component string) string {
	return fmt.Sprintf("%s-launchTemplate", component)
}
