package apis

import (
	"github.com/redhat/service-config-operator/pkg/apis/cloudnative/v1alpha11"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1alpha11.SchemeBuilder.AddToScheme)
}
