/*
Copyright 2021.

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

// SolaceRdpSpec defines the desired state of SolaceRdp
type SolaceRdpSpec struct {
	ClientProfile string                  `json:"clientProfile,omitempty"`
	Consumers     []SolaceRdpConsumer     `json:"consumers,omitempty"`
	QueueBindings []SolaceRdpQueueBinding `json:"queueBindings,omitempty"`
}

type SolaceRdpQueueBinding struct {
	QueueName         string `json:"queueName,omitempty"`
	PostRequestTarget string `json:"postRequestTarget,omitempty"`
}

type SolaceRdpConsumer struct {
	Host            string `json:"host,omitempty"`
	Port            int    `json:"port,omitempty"`
	HttpMethod      string `json:"httpMethod,omitempty"`
	TlsEnabled      bool   `json:"tlsEnabled"`
	ConnectionCount int    `json:"connectionCount,omitempty"`
}

// SolaceRdpStatus defines the observed state of SolaceRdp
type SolaceRdpStatus struct {
	OperationalState  string `json:"operationalState,omitempty"`
	LastFailureReason string `json:"lastFailureReason,omitempty"`
	LastFailureTime   string `json:"lastFailureTime,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Client Profile",type="string",JSONPath=".spec.clientProfile",description="The client profile used by the RDP"
//+kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.operationalState",description="The operational state of the RDP"
//+kubebuilder:printcolumn:name="Last Failure Reason",type="string",JSONPath=".status.lastFailureReason",description="The reason for the last failure"
//+kubebuilder:printcolumn:name="Last Failure Time",type="string",JSONPath=".status.lastFailureTime",description="The time of the last failure"

// SolaceRdp is the Schema for the solacerdps API
type SolaceRdp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SolaceRdpSpec   `json:"spec,omitempty"`
	Status SolaceRdpStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SolaceRdpList contains a list of SolaceRdp
type SolaceRdpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SolaceRdp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SolaceRdp{}, &SolaceRdpList{})
}
