package reg

import (
	"github.com/go-chassis/go-sc-client/proto"
	"github.com/patrickmn/go-cache"
)

var serviceCache = cache.New(0, 0)

func SaveInstances(s map[string][]*proto.MicroServiceInstance) {
	serviceCache.Set("instances", s, 0)

}
func GetInstances() map[string][]*proto.MicroServiceInstance {
	s, ok := serviceCache.Get("instances")
	if !ok {
		return nil
	}
	return s.(map[string][]*proto.MicroServiceInstance)

}
