//go:generate go run pkg/codegen/cleanup/main.go
//go:generate /bin/rm -rf pkg/generated
//go:generate go run pkg/codegen/main.go

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/aiyengar2/onsen/pkg/controllers/action"
	"github.com/aiyengar2/onsen/pkg/controllers/claim"
	"github.com/aiyengar2/onsen/pkg/controllers/pool"
	"github.com/aiyengar2/onsen/pkg/controllers/provisioner"
	"github.com/aiyengar2/onsen/pkg/controllers/warmcluster"
	"github.com/aiyengar2/onsen/pkg/generated/controllers/onsen.cattle.io"
	"github.com/rancher/wrangler/pkg/kubeconfig"
	"github.com/rancher/wrangler/pkg/signals"
	"github.com/rancher/wrangler/pkg/start"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	Version    = "v0.0.0-dev"
	GitCommit  = "HEAD"
	KubeConfig string
)

func main() {
	app := cli.NewApp()
	app.Name = "onsen"
	app.Version = fmt.Sprintf("%s (%s)", Version, GitCommit)
	app.Usage = "A k8s controller for creating and maintaining warm pools of provisioned clusters"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "kubeconfig",
			EnvVar:      "KUBECONFIG",
			Destination: &KubeConfig,
		},
	}
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) {
	flag.Parse()

	logrus.Info("Starting controller")
	ctx := signals.SetupSignalHandler(context.Background())

	kubeConfig, err := kubeconfig.GetNonInteractiveClientConfig(KubeConfig).ClientConfig()
	if err != nil {
		logrus.Fatalf("failed to find kubeconfig: %v", err)
	}

	types := onsen.NewFactoryFromConfigOrDie(kubeConfig)

	provisioner.Register(ctx,
		types.Onsen().V1().ClusterTemplate(),
		types.Onsen().V1().Cluster(),
	)

	action.Register(ctx,
		types.Onsen().V1().Cluster(),
		types.Onsen().V1().ClusterAction(),
	)

	warmcluster.Register(ctx,
		types.Onsen().V1().Cluster(),
		types.Onsen().V1().ClusterAction(),
		types.Onsen().V1().ClusterBootstrapConfig(),
		types.Onsen().V1().ClusterHealthCheckTemplate(),
		types.Onsen().V1().WarmCluster(),
	)

	pool.Register(ctx,
		types.Onsen().V1().WarmCluster(),
		types.Onsen().V1().WarmPool(),
	)

	claim.Register(ctx,
		types.Onsen().V1().WarmPool(),
		types.Onsen().V1().WarmClusterClaim(),
	)

	if err := start.All(ctx, 5, types); err != nil {
		logrus.Fatalf("Error starting: %s", err.Error())
	}

	<-ctx.Done()
}
