package main

import (
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis/pkg/httpclient"
	"github.com/go-mesh/openlogging"
	"github.com/go-mesh/registrator/cmd"
	"github.com/go-mesh/registrator/config"
	"github.com/go-mesh/registrator/reg"
	"github.com/go-mesh/registrator/resource"
	"github.com/huaweicse/auth"
)

func main() {
	if err := chassis.Init(); err != nil {
		openlogging.Error(err.Error())
	}
	chassis.RegisterSchema("rest", &resource.Admin{})
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
		openlogging.Warn("can not sign request: " + err.Error())
	}
	go func() {
		if err := reg.Start(); err != nil {
			openlogging.Fatal("can not start registrator: " + err.Error())
		}
	}()
	chassis.Run()
}
