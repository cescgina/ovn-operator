/*
Copyright 2020 Red Hat

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
	"github.com/operator-framework/operator-lib/status"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Important: Run "make" to regenerate code after modifying this file

// OVNCentralServer defines the observed state of a member of cluster
type OVNCentralServerStatus struct {
	Name string                         `json:"name"`
	NB   OVNCentralServerDatabaseStatus `json:"nb,omitempty"`
	SB   OVNCentralServerDatabaseStatus `json:"sb,omitempty"`
}

type OVNCentralServerDatabaseStatus struct {
	ClusterID string `json:"clusterID,omitempty"`
	Name      string `json:"name,omitempty"`
	ServerID  string `json:"serverID,omitempty"`
	Address   string `json:"address,omitempty"`
}

// OVNCentralSpec defines the desired state of OVNCentral
type OVNCentralSpec struct {
	// Required properties

	Replicas    int               `json:"replicas"`
	Image       string            `json:"image"`
	StorageSize resource.Quantity `json:"storageSize,omitempty"`

	// Required properties with default values

	ConnectionConfig string `json:"connectionConfig,omitempty"`
	ConnectionCA     string `json:"connectionCA,omitempty"`
	ConnectionCert   string `json:"connectionCert,omitempty"`
	NBSchemaVersion  string `json:"nbSchemaVersion,omitempty"`
	SBSchemaVersion  string `json:"sbSchemaVersion,omitempty"`

	// Optional properties

	StorageClass *string `json:"storageClass,omitempty"`
}

// OVNCentralStatus defines the observed state of OVNCentral
type OVNCentralStatus struct {
	Conditions      status.Conditions        `json:"conditions,omitempty"`
	Replicas        *int                     `json:"replicas,omitempty"`
	NBClusterID     *string                  `json:"nbClusterID,omitempty"`
	SBClusterID     *string                  `json:"sbClusterID,omitempty"`
	NBSchemaVersion *string                  `json:"nbSchemaVersion,omitEmpty"`
	SBSchemaVersion *string                  `json:"sbSchemaVersion,omitEmpty"`
	Servers         []OVNCentralServerStatus `json:"servers" patchStrategy:"merge" patchMergeKey:"name"`
}

const (
	OVNCentralFailed    status.ConditionType = "Failed"
	OVNCentralAvailable status.ConditionType = "Available"
)

const (
	OVNCentralInconsistentCluster status.ConditionReason = "InconsistentCluster"
	OVNCentralBootstrapFailed     status.ConditionReason = "BootstrapFailed"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// OVNCentral is the Schema for the ovncentrals API
type OVNCentral struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OVNCentralSpec   `json:"spec,omitempty"`
	Status OVNCentralStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OVNCentralList contains a list of OVNCentral
type OVNCentralList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OVNCentral `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OVNCentral{}, &OVNCentralList{})
}