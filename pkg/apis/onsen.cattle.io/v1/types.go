package v1

import (
	"github.com/rancher/wrangler/pkg/genericcondition"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterActionSpec   `json:"spec"`
	Status            ClusterActionStatus `json:"status"`
}

type ClusterActionSpec struct {
	Option bool `json:"option"`
}

type ClusterActionStatus struct {
	Conditions []genericcondition.GenericCondition `json:"conditions"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterTemplateSpec `json:"spec"`
}

type ClusterTemplateSpec struct {
	Option bool `json:"option"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status"`
}

type ClusterSpec struct {
	Option bool `json:"option"`
}

type ClusterStatus struct {
	Conditions []genericcondition.GenericCondition `json:"conditions"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterBootstrapConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterBootstrapConfigSpec `json:"spec"`
}

type ClusterBootstrapConfigSpec struct {
	Option bool `json:"option"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterHealthCheckTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterHealthCheckTemplateSpec `json:"spec"`
}

type ClusterHealthCheckTemplateSpec struct {
	Option bool `json:"option"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type WarmCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WarmClusterSpec   `json:"spec"`
	Status            WarmClusterStatus `json:"status"`
}

type WarmClusterSpec struct {
	Option bool `json:"option"`
}

type WarmClusterStatus struct {
	Conditions []genericcondition.GenericCondition `json:"conditions"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type WarmPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WarmPoolSpec   `json:"spec"`
	Status            WarmPoolStatus `json:"status"`
}

type WarmPoolSpec struct {
	Option bool `json:"option"`
}

type WarmPoolStatus struct {
	Conditions []genericcondition.GenericCondition `json:"conditions"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type WarmClusterClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WarmClusterClaimSpec   `json:"spec"`
	Status            WarmClusterClaimStatus `json:"status"`
}

type WarmClusterClaimSpec struct {
	Option bool `json:"option"`
}

type WarmClusterClaimStatus struct {
	Conditions []genericcondition.GenericCondition `json:"conditions"`
}
