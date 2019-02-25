package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

//MyResource describes a MyResource resoruce
type MyCustomResource struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:"metadata, omitempty"`

	// Spec is the custom resource  Spec
	Spec MyCustomResourceSpc `json:"spec"`
}

//MyResourceSpec is the spec fot a MyResource resource
type MyCustomResourceSpc struct {
	Message   string `json:"message"`
	SomeValue *int32 `json:"someValue"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MyCustomResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metedata"`

	Items []MyCustomResource `json:"items"`
}
