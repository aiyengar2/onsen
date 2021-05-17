package claim

import v1 "github.com/aiyengar2/onsen/pkg/apis/onsen.cattle.io/v1"

func (h *handler) OnWarmClusterClaimChange(key string, warmClusterClaim *v1.WarmClusterClaim) (*v1.WarmClusterClaim, error) {
	//change logic, return original warmClusterClaim if no changes

	warmClusterClaimCopy := warmClusterClaim.DeepCopy()
	//make changes to warmClusterClaimCopy
	return h.warmClusterClaims.Update(warmClusterClaimCopy)
}

func (h *handler) OnWarmClusterClaimRemove(key string, warmClusterClaim *v1.WarmClusterClaim) (*v1.WarmClusterClaim, error) {
	//remove logic, return original warmClusterClaim if no changes

	warmClusterClaimCopy := warmClusterClaim.DeepCopy()
	//make changes to warmClusterClaimCopy
	return h.warmClusterClaims.Update(warmClusterClaimCopy)
}
