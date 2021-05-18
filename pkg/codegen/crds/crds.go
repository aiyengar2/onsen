package crds

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	wrangler "github.com/rancher/wrangler/pkg/crd"
	"github.com/rancher/wrangler/pkg/data/convert"
	_ "github.com/rancher/wrangler/pkg/generated/controllers/apiextensions.k8s.io/v1" // Imported to use init function
	"github.com/rancher/wrangler/pkg/yaml"
	apiext "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type CustomResourceDefinition struct {
	Type interface{}

	APIGroup   string
	Version    string
	Namespaced bool

	Customizer customizer
}

func (c *CustomResourceDefinition) WriteCRD(toDir string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	toAbsDir := filepath.Join(dir, toDir)
	_, err = os.Stat(toAbsDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(toAbsDir, os.ModePerm)
	}
	if err != nil {
		return err
	}

	crdObject, err := c.GetCRD().ToCustomResourceDefinition()
	if err != nil {
		return err
	}
	var crd apiext.CustomResourceDefinition
	if err := convert.ToObj(crdObject, &crd); err != nil {
		return fmt.Errorf("could not convert object into apiextensions/v1 crd: %s", err)
	}

	if c.Customizer != nil {
		c.Customizer.CustomizeOutput(&crd)
	}

	yamlBytes, err := yaml.Export(&crd)
	if err != nil {
		return err
	}

	filename := filepath.Join(toDir, strings.ToLower(crd.Spec.Names.Kind)+".yaml")
	return ioutil.WriteFile(filename, yamlBytes, 0644)
}

func (c *CustomResourceDefinition) GetCRD() wrangler.CRD {
	crd := wrangler.CRD{
		GVK: schema.GroupVersionKind{
			Group:   c.APIGroup,
			Version: c.Version,
		},
		Status:       true,
		SchemaObject: c.Type,
		NonNamespace: !c.Namespaced,
	}
	if c.Customizer != nil {
		crd = c.Customizer.CustomizeObject(crd)
	}
	return crd
}
