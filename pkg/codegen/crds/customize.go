package crds

import (
	"github.com/rancher/wrangler/pkg/crd"
	apiext "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type customizer interface {
	// CustomizeObject is used to modify the wrangler.CRD.
	// This can be used to add in custom columns, e.g. c.WithColumn("Location", ".status.storageLocation")
	// The modified wrangler.CRD should be returned.
	CustomizeObject(c crd.CRD) crd.CRD

	// CustomizeOutput is used to modify the final CustomResourceDefinition before writing it.
	// This can be used to add additional validation to the schema or to modify the spec.
	// The CustomResourceDefinition should be modified in place.
	CustomizeOutput(c *apiext.CustomResourceDefinition)
}
