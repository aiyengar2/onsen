package claim

import v1 "github.com/aiyengar2/onsen/pkg/apis/onsen.cattle.io/v1"

func (h *handler) OnWarmPoolChange(key string, warmPool *v1.WarmPool) (*v1.WarmPool, error) {
	//change logic, return original warmPool if no changes

	warmPoolCopy := warmPool.DeepCopy()
	//make changes to warmPoolCopy
	return h.warmPools.Update(warmPoolCopy)
}

func (h *handler) OnWarmPoolRemove(key string, warmPool *v1.WarmPool) (*v1.WarmPool, error) {
	//remove logic, return original warmPool if no changes

	warmPoolCopy := warmPool.DeepCopy()
	//make changes to warmPoolCopy
	return h.warmPools.Update(warmPoolCopy)
}
