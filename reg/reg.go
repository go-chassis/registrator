package reg

import (
	"crypto/tls"
	"os"
	"strings"
	"time"

	"fmt"
	"github.com/go-chassis/go-chassis/core/registry"
	"github.com/go-chassis/go-chassis/pkg/runtime"
	"github.com/go-chassis/go-chassis/pkg/scclient"
	"github.com/go-chassis/go-chassis/pkg/scclient/proto"
	"github.com/go-mesh/openlogging"
	"github.com/go-mesh/registrator/cmd"
	"github.com/go-mesh/registrator/config"
)

var targetRegistry *client.RegistryClient
var sourceRegistry *client.RegistryClient
var Self = &proto.MicroService{
	ServiceName: "REGISTRATOR",
	Version:     "0.1",
	AppId:       "default",
	Environment: os.Getenv("CSE_ENV"),
}

func getSourceRegistry() (*client.RegistryClient, error) {
	opts, err := GetOptions(config.Config.Source.Address)
	if err != nil {
		return nil, err
	}
	r := &client.RegistryClient{}
	if err := r.Initialize(opts); err != nil {
		openlogging.GetLogger().Errorf("RegistryClient initialization failed, err %s", err)
		return nil, err
	}
	return r, nil
}

func getTargetRegistry() (*client.RegistryClient, error) {
	opts, err := GetOptions(config.Config.Target.Address)
	if err != nil {
		return nil, err
	}
	r := &client.RegistryClient{}
	if err := r.Initialize(opts); err != nil {
		openlogging.GetLogger().Errorf("RegistryClient initialization failed, err %s", err)
		return nil, err
	}
	return r, nil
}
func GetOptions(address string) (client.Options, error) {
	hosts, schema, err := registry.URIs2Hosts(strings.Split(address, ","))
	if err != nil {
		return client.Options{}, err
	}
	ssl := false
	var tlsConfig *tls.Config
	if schema == "https" {
		ssl = true
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	sco := client.Options{}
	sco.TLSConfig = tlsConfig
	sco.Addrs = hosts
	sco.EnableSSL = ssl
	return sco, nil
}
func doFetchService() {
	ms, err := sourceRegistry.GetAllMicroServices()
	if err != nil {
		openlogging.Error("can not get services: " + err.Error())
	}
	criteria := make([]*proto.FindService, 0)
	for _, s := range ms {
		if strings.Contains(config.Config.Source.Exclude, s.ServiceName) {
			openlogging.Info("skip: " + s.ServiceName)
			continue
		}
		_, err = targetRegistry.RegisterService(s)
		if err != nil {
			openlogging.Error("can not register service:" + err.Error())
			continue
		}
		f := &proto.FindService{
			Service: &proto.MicroServiceKey{
				ServiceName: s.ServiceName,
				Version:     s.Version,
				AppId:       s.AppId,
			},
		}
		criteria = append(criteria, f)

	}
	instancesMap, err := sourceRegistry.BatchFindInstances(runtime.InstanceID, criteria)
	if err != nil {
		openlogging.Error("can not get instances: " + err.Error())
	}
	for _, instances := range instancesMap {
		for _, ins := range instances {
			_, err := targetRegistry.RegisterMicroServiceInstance(ins)
			if err != nil {
				openlogging.Error(
					fmt.Sprintf("can not register instance [%s]: %s", ins.InstanceId, err.Error()))
			}
		}

	}
	SaveInstances(instancesMap)
}

//FetchService pull services and instances from source registry
//register them to target registry
func FetchService() error {
	fi, err := time.ParseDuration(cmd.CLIParam.FetchInterval)
	if err != nil {
		return err
	}
	ft := time.NewTicker(fi)
	doFetchService()
	go func() {

		for range ft.C {
			doFetchService()
		}
	}()
	return nil
}
func Heartbeat() error {
	hb, err := time.ParseDuration(cmd.CLIParam.RegisterInterval)
	if err != nil {
		return err
	}
	hbt := time.NewTicker(hb)
	for range hbt.C {
		instanceMap := GetInstances()
		for _, instances := range instanceMap {
			for _, ins := range instances {
				_, err := targetRegistry.Heartbeat(ins.ServiceId, ins.InstanceId)
				if err != nil {
					openlogging.Error("hb failed" + err.Error())
					_, err := targetRegistry.RegisterMicroServiceInstance(ins)
					if err != nil {
						openlogging.Error(
							fmt.Sprintf("can not register instance [%s]: %s", ins.InstanceId, err.Error()))

					}
				}
			}
		}
	}
	return nil
}

func Start() error {
	var err error
	targetRegistry, err = getTargetRegistry()
	if err != nil {
		return err
	}
	sourceRegistry, err = getSourceRegistry()
	if err != nil {
		return err
	}

	if err := FetchService(); err != nil {
		return err
	}
	if err := Heartbeat(); err != nil {
		return err
	}
	return nil
}
