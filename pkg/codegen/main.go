package main

import (
	"os"

	v1 "github.com/aiyengar2/onsen/pkg/apis/onsen.cattle.io/v1"
	"github.com/aiyengar2/onsen/pkg/codegen/crds"
	controllergen "github.com/rancher/wrangler/pkg/controller-gen"
	"github.com/rancher/wrangler/pkg/controller-gen/args"
)

var (
	CustomResourceDefinitions = []crds.CustomResourceDefinition{
		{
			APIGroup: "onsen.cattle.io",
			Version:  "v1",
			Type:     v1.ClusterTemplate{},
		},
		{
			APIGroup: "onsen.cattle.io",
			Version:  "v1",
			Type:     v1.Cluster{},
		},
		{
			APIGroup: "onsen.cattle.io",
			Version:  "v1",
			Type:     v1.ClusterAction{},
		},
		{
			APIGroup: "onsen.cattle.io",
			Version:  "v1",
			Type:     v1.ClusterBootstrapConfig{},
		},
		{
			APIGroup: "onsen.cattle.io",
			Version:  "v1",
			Type:     v1.ClusterHealthCheckTemplate{},
		},
		{
			APIGroup: "onsen.cattle.io",
			Version:  "v1",
			Type:     v1.WarmCluster{},
		},
		{
			APIGroup: "onsen.cattle.io",
			Version:  "v1",
			Type:     v1.WarmPool{},
		},
		{
			APIGroup: "onsen.cattle.io",
			Version:  "v1",
			Type:     v1.WarmClusterClaim{},
		},
	}
)

func main() {
	os.Unsetenv("GOPATH")
	controllergen.Run(args.Options{
		OutputPackage: "github.com/aiyengar2/onsen/pkg/generated",
		Boilerplate:   "scripts/boilerplate.go.txt",
		Groups:        crds.GetGroups(CustomResourceDefinitions),
	})
	if err := crds.WriteCRDs(CustomResourceDefinitions, "crds"); err != nil {
		panic(err)
	}
}
