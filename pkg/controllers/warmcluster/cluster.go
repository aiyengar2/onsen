package warmcluster

import v1 "github.com/aiyengar2/onsen/pkg/apis/onsen.cattle.io/v1"

func (h *handler) OnClusterChange(key string, cluster *v1.Cluster) (*v1.Cluster, error) {
	//change logic, return original warmCluster if no changes

	clusterCopy := cluster.DeepCopy()
	//make changes to warmClusterCopy
	return h.clusters.Update(clusterCopy)
}

func (h *handler) OnClusterRemove(key string, cluster *v1.Cluster) (*v1.Cluster, error) {
	//remove logic, return original warmCluster if no changes

	clusterCopy := cluster.DeepCopy()
	//make changes to warmClusterCopy
	return h.clusters.Update(clusterCopy)
}
