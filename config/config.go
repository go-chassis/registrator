package config

import (
	"io/ioutil"
	"os"

	"github.com/go-mesh/registrator/cmd"
	"gopkg.in/yaml.v2"
)

const (
	SourceAddressEnv     = "SOURCE_ADDRESS"
	TargetAddressEnv     = "TARGET_ADDRESS"
	FetchIntervalEnv     = "FETCH_INTERVAL"
	HeartbeatIntervalEnv = "HEARTBEAT_INTERVAL"
)

var Config = &Configuration{}

func Init() error {
	b, err := ioutil.ReadFile(cmd.CLIParam.ConfPath)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(b, Config); err != nil {
		return err
	}
	if v := os.Getenv(SourceAddressEnv); v != "" {
		Config.Source.Address = v
	}
	if v := os.Getenv(FetchIntervalEnv); v != "" {
		Config.Source.FetchInterval = v
	}
	if v := os.Getenv(TargetAddressEnv); v != "" {
		Config.Target.Address = v
	}
	if v := os.Getenv(HeartbeatIntervalEnv); v != "" {
		Config.Target.HeartbeatInterval = v
	}
	setDefaultValue()
	return nil
}

func setDefaultValue() {
	if Config.Source.FetchInterval == "" {
		Config.Source.FetchInterval = "120s"
	}
	if Config.Target.HeartbeatInterval == "" {
		Config.Target.HeartbeatInterval = "30s"
	}
}
