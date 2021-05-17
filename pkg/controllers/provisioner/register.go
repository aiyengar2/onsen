package provisioner

import (
	"context"

	onsenController "github.com/aiyengar2/onsen/pkg/generated/controllers/onsen.cattle.io/v1"
)

type handler struct {
	clusterTemplates onsenController.ClusterTemplateController
	clusters         onsenController.ClusterController
}

func Register(
	ctx context.Context,
	clusterTemplates onsenController.ClusterTemplateController,
	clusters onsenController.ClusterController) {

	controller := &handler{
		clusterTemplates: clusterTemplates,
		clusters:         clusters,
	}

	clusters.OnChange(ctx, "cluster-handler", controller.OnClusterChange)
	clusters.OnRemove(ctx, "cluster-handler", controller.OnClusterRemove)
}
