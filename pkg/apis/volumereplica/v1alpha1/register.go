package v1alpha1

import (
	"github.com/numanelahi/volumereplication/pkg/apis/volumereplica"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemaGroupVersion = schema.GroupVersion{
	Group:   volumereplica.GroupName,
	Version: "v1alpha1",
}

func Kind(kind string) schema.GroupKind {
	return SchemaGroupVersion.WithKind(kind).GroupKind()
}

func Resource(resource string) schema.GroupResource {
	return SchemaGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemaGroupVersion, &VolumeReplica{}, &VolumeReplicaList{})
	metav1.AddToGroupVersion(scheme, SchemaGroupVersion)
	return nil
}
