/*
Copyright 2021 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

// +k8s:deepcopy-gen=package
// +groupName=onsen.cattle.io
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterTemplateList is a list of ClusterTemplate resources
type ClusterTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ClusterTemplate `json:"items"`
}

func NewClusterTemplate(namespace, name string, obj ClusterTemplate) *ClusterTemplate {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("ClusterTemplate").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterList is a list of Cluster resources
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Cluster `json:"items"`
}

func NewCluster(namespace, name string, obj Cluster) *Cluster {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("Cluster").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterActionList is a list of ClusterAction resources
type ClusterActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ClusterAction `json:"items"`
}

func NewClusterAction(namespace, name string, obj ClusterAction) *ClusterAction {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("ClusterAction").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterBootstrapConfigList is a list of ClusterBootstrapConfig resources
type ClusterBootstrapConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ClusterBootstrapConfig `json:"items"`
}

func NewClusterBootstrapConfig(namespace, name string, obj ClusterBootstrapConfig) *ClusterBootstrapConfig {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("ClusterBootstrapConfig").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterHealthCheckTemplateList is a list of ClusterHealthCheckTemplate resources
type ClusterHealthCheckTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ClusterHealthCheckTemplate `json:"items"`
}

func NewClusterHealthCheckTemplate(namespace, name string, obj ClusterHealthCheckTemplate) *ClusterHealthCheckTemplate {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("ClusterHealthCheckTemplate").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WarmClusterList is a list of WarmCluster resources
type WarmClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []WarmCluster `json:"items"`
}

func NewWarmCluster(namespace, name string, obj WarmCluster) *WarmCluster {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("WarmCluster").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WarmPoolList is a list of WarmPool resources
type WarmPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []WarmPool `json:"items"`
}

func NewWarmPool(namespace, name string, obj WarmPool) *WarmPool {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("WarmPool").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WarmClusterClaimList is a list of WarmClusterClaim resources
type WarmClusterClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []WarmClusterClaim `json:"items"`
}

func NewWarmClusterClaim(namespace, name string, obj WarmClusterClaim) *WarmClusterClaim {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("WarmClusterClaim").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}
