/*
Copyright 2023.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StandingsSpec defines the desired state of Standings
type StandingsSpec struct {
	// DataSourceUrl is the URL of the data source
	//+kubebuilder:validation:Required
	DataSourceUrl string `json:"dataSourceUrl,omitempty"`

	// ConfigMapName is the name of the ConfigMap to store the standings
	//+kubebuilder:validation:Required
	ConfigMapName string `json:"configMapName,omitempty"`

	// ConfigMapKey is the key of the ConfigMap to store the standings
	//+kubebuilder:validation:Required
	ConfigMapKey string `json:"configMapKey,omitempty"`
}

// StandingsStatus defines the observed state of Standings
type StandingsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Standings is the Schema for the standings API
type Standings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StandingsSpec   `json:"spec,omitempty"`
	Status StandingsStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// StandingsList contains a list of Standings
type StandingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Standings `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Standings{}, &StandingsList{})
}
