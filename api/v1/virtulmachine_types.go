/*
Copyright 2020 alex.

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

// VirtulMachineSpec defines the desired state of VirtulMachine
type VirtulMachineSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of VirtulMachine. Edit VirtulMachine_types.go to remove/update
	Foo    string `json:"foo,omitempty"`
	Cpu    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

// VirtulMachineStatus defines the observed state of VirtulMachine
type VirtulMachineStatus struct {
	Status string `json:"status"`
}

// +kubebuilder:subresource:status
// +kubebuilder:object:root=true

// VirtulMachine is the Schema for the virtulmachines API
type VirtulMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtulMachineSpec   `json:"spec,omitempty"`
	Status VirtulMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtulMachineList contains a list of VirtulMachine
type VirtulMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtulMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtulMachine{}, &VirtulMachineList{})
}
