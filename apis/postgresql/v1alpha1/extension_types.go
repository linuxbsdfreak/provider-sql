/*
Copyright 2020 The Crossplane Authors.

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

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ExtensionParameters are the configurable fields of a Extension.
type ExtensionParameters struct {
	// Extension name to be installed
	Extension *string `json:"extension"`

	// Extension Version to be installed
	// +immutable
	// +optional
	Version *string `json:"version,omitempty"`

	// Database for extension install.
	// +optional
	Database *string `json:"database"`

	// DatabaseRef references the database object this extensions it for.
	// +immutable
	// +optional
	DatabaseRef *xpv1.Reference `json:"databaseRef"`

	// DatabaseSelector selects a reference to a Database this extension is for.
	// +immutable
	// +optional
	DatabaseSelector *xpv1.Selector `json:"databaseSelector"`
}

// A ExtensionSpec defines the desired state of a Extension.
type ExtensionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ExtensionParameters `json:"forProvider"`
}

// A ExtensionStatus represents the observed state of a Extension.
type ExtensionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// A Extension represents the declarative state of a PostgreSQL Extension.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,sql}
type Extension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExtensionSpec   `json:"spec"`
	Status ExtensionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExtensionList contains a list of Extension
type ExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Extension `json:"items"`
}
