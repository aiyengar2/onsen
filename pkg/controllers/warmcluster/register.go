package warmcluster

import (
	"context"

	onsenController "github.com/aiyengar2/onsen/pkg/generated/controllers/onsen.cattle.io/v1"
)

type handler struct {
	clusters                    onsenController.ClusterController
	clusterActions              onsenController.ClusterActionController
	clusterBootstrapConfigs     onsenController.ClusterBootstrapConfigController
	clusterHealthCheckTemplates onsenController.ClusterHealthCheckTemplateController
	warmClusters                onsenController.WarmClusterController
}

func Register(
	ctx context.Context,
	clusters onsenController.ClusterController,
	clusterActions onsenController.ClusterActionController,
	clusterBootstrapConfigs onsenController.ClusterBootstrapConfigController,
	clusterHealthCheckTemplates onsenController.ClusterHealthCheckTemplateController,
	warmClusters onsenController.WarmClusterController) {

	controller := &handler{
		clusters: clusters,

		clusterActions:              clusterActions,
		clusterBootstrapConfigs:     clusterBootstrapConfigs,
		clusterHealthCheckTemplates: clusterHealthCheckTemplates,

		warmClusters: warmClusters,
	}

	clusters.OnChange(ctx, "cluster-handler", controller.OnClusterChange)
	clusters.OnRemove(ctx, "cluster-handler", controller.OnClusterRemove)

	clusterActions.OnChange(ctx, "cluster-action-handler", controller.OnClusterActionChange)
	clusterActions.OnRemove(ctx, "cluster-action-handler", controller.OnClusterActionRemove)

	warmClusters.OnChange(ctx, "warm-cluster-handler", controller.OnWarmClusterChange)
	warmClusters.OnRemove(ctx, "warm-cluster-handler", controller.OnWarmClusterRemove)
}
