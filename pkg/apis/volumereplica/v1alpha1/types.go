package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VolumeReplicaSpec struct {
	SourcePVC string `json:"sourcePvc,omitempty"`
	TargetPVC string `json:"targetPvc,omitempty"`
	ReplicaClass string `json:"replicaClass,omitempty"`
}

type VolumeReplica struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:metadata,omitempty`
	Spec VolumeReplicaSpec `json:"spec,omitempty"`
}

type VolumeReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items []VolumeReplica `json:"items"`
}