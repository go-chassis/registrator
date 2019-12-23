package reg

import (
	"github.com/go-chassis/go-chassis/pkg/scclient/proto"
	"github.com/patrickmn/go-cache"
)

const instanceKey = "instance"

var serviceCache = cache.New(0, 0)

// SaveInstances merges new instances to old ones, not remove any old instance
func SaveInstances(s map[string][]*proto.MicroServiceInstance) {
	oldInstances := GetInstances()
	if oldInstances == nil {
		serviceCache.Set(instanceKey, s, 0)
		return
	}
	result := make(map[string][]*proto.MicroServiceInstance, len(oldInstances))
	for k, v := range oldInstances {
		result[k] = v
	}
	for k, v := range s {
		result[k] = v
	}
	serviceCache.Set(instanceKey, result, 0)
}
func GetInstances() map[string][]*proto.MicroServiceInstance {
	s, ok := serviceCache.Get(instanceKey)
	if !ok {
		return nil
	}
	return s.(map[string][]*proto.MicroServiceInstance)
}
