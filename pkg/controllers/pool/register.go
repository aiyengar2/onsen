package pool

import (
	"context"

	onsenController "github.com/aiyengar2/onsen/pkg/generated/controllers/onsen.cattle.io/v1"
)

type handler struct {
	warmClusters onsenController.WarmClusterController
	warmPools    onsenController.WarmPoolController
}

func Register(
	ctx context.Context,
	warmClusters onsenController.WarmClusterController,
	warmPools onsenController.WarmPoolController) {

	controller := &handler{
		warmClusters: warmClusters,
		warmPools:    warmPools,
	}

	warmClusters.OnChange(ctx, "warm-cluster-handler", controller.OnWarmClusterChange)
	warmClusters.OnRemove(ctx, "warm-cluster-handler", controller.OnWarmClusterRemove)

	warmPools.OnChange(ctx, "warm-pool-handler", controller.OnWarmPoolChange)
	warmPools.OnRemove(ctx, "warm-pool-handler", controller.OnWarmPoolRemove)
}
