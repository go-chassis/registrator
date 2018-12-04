package reg

import (
	"github.com/go-chassis/go-chassis/core/registry"
	"github.com/patrickmn/go-cache"
)

var serviceCache = cache.New(0, 0)
var InstanceCache = registry.NewIndexCache()

func SaveServices(s []*registry.MicroService) {
	serviceCache.Set("services", s, 0)

}
func GetServices() []*registry.MicroService {
	s, ok := serviceCache.Get("services")
	if !ok {
		return nil
	}
	return s.([]*registry.MicroService)

}
