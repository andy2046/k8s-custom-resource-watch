package v1

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomResource defines a CustomResource resource
type CustomResource struct {
	// TypeMeta is the metadata for the resource, e.g. kind / apiversion
	metaV1.TypeMeta `json:",inline"`
	// ObjectMeta is the metadata for the particular object, e.g. name / namespace / labels
	metaV1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the custom resource spec
	Spec CustomResourceSpec `json:"spec"`
}

// CustomResourceSpec is the spec for a CustomResource resource
type CustomResourceSpec struct {
	// Message and Count are example custom spec fields
	Message string `json:"message"`
	Count   *int32 `json:"count"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomResourceList is a list of CustomResource resources
type CustomResourceList struct {
	metaV1.TypeMeta `json:",inline"`
	metaV1.ListMeta `json:"metadata"`

	Items []CustomResource `json:"items"`
}
