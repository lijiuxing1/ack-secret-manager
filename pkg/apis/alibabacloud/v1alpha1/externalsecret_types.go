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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExternalSecretSpec defines the desired state of ExternalSecret
// +k8s:openapi-gen=true
type ExternalSecretSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Data []DataSource `json:"data,omitempty"`
	Type string       `json:"type,omitempty"`
}

type DataSource struct {
	Key          string `json:"key"`
	Name         string `json:"name"`
	VersionStage string `json:"versionStage,omitempty"`
	VersionId    string `json:"versionId,omitempty"`

	//Optional array to specify what json key value pairs to extract from a secret and mount as individual secrets
	JMESPath []JMESPathObject `json:"jmesPath"`
}

// ExternalSecretStatus defines the observed state of ExternalSecret
type ExternalSecretStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExternalSecret is the Schema for the externalsecrets API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=externalsecrets,scope=Namespaced
type ExternalSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExternalSecretSpec   `json:"spec,omitempty"`
	Status ExternalSecretStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExternalSecretList contains a list of ExternalSecret
type ExternalSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalSecret `json:"items"`
}

//An individual json key value pair to mount
type JMESPathObject struct {
	//JMES path to use for retrieval
	Path string `json:"path"`

	//File name in which to store the secret in.
	ObjectAlias string `json:"objectAlias"`
}

func init() {
	SchemeBuilder.Register(&ExternalSecret{}, &ExternalSecretList{})
}
