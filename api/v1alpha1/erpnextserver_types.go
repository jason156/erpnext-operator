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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Important: Run "make" to regenerate code after modifying this file
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ERPNextServerSpec defines the desired state of ERPNextServer
type ERPNextServerSpec struct {
	// Version is the ERPNext server version to install and manage.
	Version string `json:"version,omitempty"`
}

// ERPNextServerStatus defines the observed state of ERPNextServer
type ERPNextServerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ERPNextServer is the Schema for the erpnextservers API
type ERPNextServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ERPNextServerSpec   `json:"spec,omitempty"`
	Status ERPNextServerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ERPNextServerList contains a list of ERPNextServer
type ERPNextServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ERPNextServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ERPNextServer{}, &ERPNextServerList{})
}
