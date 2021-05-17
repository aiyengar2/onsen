package action

import (
	"context"

	onsenController "github.com/aiyengar2/onsen/pkg/generated/controllers/onsen.cattle.io/v1"
)

type handler struct {
	clusters       onsenController.ClusterController
	clusterActions onsenController.ClusterActionController
}

func Register(
	ctx context.Context,
	clusters onsenController.ClusterController,
	clusterActions onsenController.ClusterActionController) {

	controller := &handler{
		clusters:       clusters,
		clusterActions: clusterActions,
	}

	clusterActions.OnChange(ctx, "cluster-action-handler", controller.OnClusterActionChange)
	clusterActions.OnRemove(ctx, "cluster-action-handler", controller.OnClusterActionRemove)
}
