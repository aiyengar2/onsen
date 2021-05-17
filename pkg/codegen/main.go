package main

import (
	"os"

	v1 "github.com/aiyengar2/onsen/pkg/apis/onsen.cattle.io/v1"
	controllergen "github.com/rancher/wrangler/pkg/controller-gen"
	"github.com/rancher/wrangler/pkg/controller-gen/args"
)

func main() {
	os.Unsetenv("GOPATH")
	controllergen.Run(args.Options{
		OutputPackage: "github.com/aiyengar2/onsen/pkg/generated",
		Boilerplate:   "scripts/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"onsen.cattle.io": {
				Types: []interface{}{
					v1.ClusterTemplate{},
					v1.Cluster{},
					v1.ClusterAction{},
					v1.ClusterBootstrapConfig{},
					v1.ClusterHealthCheckTemplate{},
					v1.WarmCluster{},
					v1.WarmPool{},
					v1.WarmClusterClaim{},
				},
				GenerateTypes: true,
			},
		},
	})
}
