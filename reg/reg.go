package reg

import (
	"github.com/go-chassis/go-chassis/core/registry"
	"github.com/go-chassis/go-chassis/core/registry/servicecenter"
	"github.com/go-chassis/go-chassis/third_party/forked/k8s.io/apimachinery/pkg/util/sets"
	"github.com/go-mesh/openlogging"
	"github.com/go-mesh/registrator/cmd"
	"github.com/go-mesh/registrator/config"
	"time"
)

var registrator registry.Registrator
var discovery registry.ServiceDiscovery

func getRegistrator() (registry.Registrator, error) {

	return registry.NewRegistrator(config.Config.Registrator.T, registry.Options{
		Addrs: []string{config.Config.Discovery.Address},
	})
}

func getDiscoveryService() (registry.ServiceDiscovery, error) {

	return registry.NewDiscovery(config.Config.Discovery.T, registry.Options{
		Addrs: []string{config.Config.Discovery.Address},
	})
}

func FetchService() error {
	fi, err := time.ParseDuration(cmd.CLIParam.FetchInterval)
	if err != nil {
		return err
	}
	ft := time.NewTicker(fi)
	go func() {
		for range ft.C {
			ms, err := discovery.GetAllMicroServices()
			if err != nil {
				openlogging.Error("can not get services: " + err.Error())
			}
			SaveServices(ms)
			appService := sets.NewString()
			for _, s := range ms {
				s.ServiceID, err = registrator.RegisterService(s)
				if err != nil {
					openlogging.Error("can not register service:" + err.Error())
				}
				k := s.ServiceName + "::" + s.AppID
				if appService.Has(k) {
					continue
				}
				appService.Insert(k)
				//TODO batch find
				//TODO set cache

			}
		}
	}()
	return nil
}
func SyncService() error {
	hb, err := time.ParseDuration(cmd.CLIParam.RegisterInterval)
	if err != nil {
		return err
	}
	hbt := time.NewTicker(hb)
	for range hbt.C {
		services := GetServices()
		for _, s := range services {
			instances, ok := InstanceCache.Get(s.ServiceName, map[string]string{
				"app":     s.AppID,
				"version": s.Version,
			})
			if ok {
				for _, i := range instances {
					_, err := registrator.Heartbeat(s.ServiceID, i.InstanceID)
					if err != nil {
						openlogging.Error("hb failed" + err.Error())
						_, err := registrator.RegisterServiceInstance(s.ServiceID, i)
						if err != nil {
							openlogging.Error("register instance failed" + err.Error())
						}
					}
				}
			}
		}
	}
	return nil
}
func Start() error {
	registry.InstallRegistrator("servicecenter", servicecenter.NewRegistrator)
	registry.InstallServiceDiscovery("servicecenter", servicecenter.NewServiceDiscovery)

	var err error
	registrator, err = getRegistrator()
	if err != nil {
		return err
	}
	discovery, err = getDiscoveryService()
	if err != nil {
		return err
	}

	if err := FetchService(); err != nil {
		return err
	}
	if err := SyncService(); err != nil {
		return err
	}
	return nil
}
