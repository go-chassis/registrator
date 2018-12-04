package main

import (
	"github.com/go-chassis/go-chassis/pkg/httpclient"
	"github.com/go-chassis/paas-lager"
	"github.com/go-mesh/openlogging"
	"github.com/go-mesh/registrator/cmd"
	"github.com/go-mesh/registrator/config"
	"github.com/go-mesh/registrator/reg"
	"github.com/huaweicse/auth"
)

func main() {
	log.Init(log.Config{
		Writers:       []string{"file", "stdout"},
		LoggerLevel:   "DEBUG",
		LoggerFile:    "./reg.log",
		LogFormatText: true,
	})

	logger := log.NewLogger("reg")
	openlogging.SetLogger(logger)
	if err := cmd.ReadParams(); err != nil {
		openlogging.Fatal("can not init CLI: " + err.Error())
	}
	if err := config.ReadYAML(); err != nil {
		openlogging.Fatal("can not read config: " + err.Error())
	}
	var err error
	httpclient.SignRequest, err = auth.GetShaAKSKSignFunc(config.Config.Auth.AK,
		config.Config.Auth.SK,
		config.Config.Auth.Project)
	if err != nil {
		openlogging.Fatal("can not sign request: " + err.Error())
	}
	if err := reg.Start(); err != nil {
		openlogging.Fatal("can not start registrator: " + err.Error())
	}
}
