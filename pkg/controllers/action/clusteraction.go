package action

import v1 "github.com/aiyengar2/onsen/pkg/apis/onsen.cattle.io/v1"

func (h *handler) OnClusterActionChange(key string, clusterAction *v1.ClusterAction) (*v1.ClusterAction, error) {
	//change logic, return original warmCluster if no changes

	clusterActionCopy := clusterAction.DeepCopy()
	//make changes to warmClusterCopy
	return h.clusterActions.Update(clusterActionCopy)
}

func (h *handler) OnClusterActionRemove(key string, clusterAction *v1.ClusterAction) (*v1.ClusterAction, error) {
	//remove logic, return original warmCluster if no changes

	clusterActionCopy := clusterAction.DeepCopy()
	//make changes to warmClusterCopy
	return h.clusterActions.Update(clusterActionCopy)
}
