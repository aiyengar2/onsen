package claim

import (
	"context"

	onsenController "github.com/aiyengar2/onsen/pkg/generated/controllers/onsen.cattle.io/v1"
)

type handler struct {
	warmPools         onsenController.WarmPoolController
	warmClusterClaims onsenController.WarmClusterClaimController
}

func Register(
	ctx context.Context,
	warmPools onsenController.WarmPoolController,
	warmClusterClaims onsenController.WarmClusterClaimController) {

	controller := &handler{
		warmPools:         warmPools,
		warmClusterClaims: warmClusterClaims,
	}

	warmPools.OnChange(ctx, "warm-pool-handler", controller.OnWarmPoolChange)
	warmPools.OnRemove(ctx, "warm-pool-handler", controller.OnWarmPoolRemove)

	warmClusterClaims.OnChange(ctx, "warm-cluster-claim-handler", controller.OnWarmClusterClaimChange)
	warmClusterClaims.OnRemove(ctx, "warm-cluster-claim-handler", controller.OnWarmClusterClaimRemove)
}
