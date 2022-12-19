/*
Copyright 2022.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TimeSeriesDBSpec defines the desired state of TimeSeriesDB
type TimeSeriesDBSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	DBType   string `json:"dbType,omitempty"`
	Replicas int    `json:"replicas,omitempty"`
}

// TimeSeriesDBStatus defines the observed state of TimeSeriesDB
type TimeSeriesDBStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TimeSeriesDB is the Schema for the timeseriesdbs API
type TimeSeriesDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TimeSeriesDBSpec   `json:"spec,omitempty"`
	Status TimeSeriesDBStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TimeSeriesDBList contains a list of TimeSeriesDB
type TimeSeriesDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TimeSeriesDB `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TimeSeriesDB{}, &TimeSeriesDBList{})
}
