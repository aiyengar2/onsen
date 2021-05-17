package pool

import v1 "github.com/aiyengar2/onsen/pkg/apis/onsen.cattle.io/v1"

func (h *handler) OnWarmClusterChange(key string, warmCluster *v1.WarmCluster) (*v1.WarmCluster, error) {
	//change logic, return original warmCluster if no changes

	warmClusterCopy := warmCluster.DeepCopy()
	//make changes to warmClusterCopy
	return h.warmClusters.Update(warmClusterCopy)
}

func (h *handler) OnWarmClusterRemove(key string, warmCluster *v1.WarmCluster) (*v1.WarmCluster, error) {
	//remove logic, return original warmCluster if no changes

	warmClusterCopy := warmCluster.DeepCopy()
	//make changes to warmClusterCopy
	return h.warmClusters.Update(warmClusterCopy)
}
