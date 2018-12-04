package config

import (
	"github.com/go-mesh/registrator/cmd"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Config = &Configuration{}

type Configuration struct {
	Discovery   Discovery   `yaml:"discovery"`
	Registrator Registrator `yaml:"registrator"`
	Auth        Auth        `yaml:"auth"`
}

type Discovery struct {
	Address string `yaml:"address"`
	T       string `yaml:"type"`
	Auth    Auth   `yaml:"auth"` //TODO register between tenants
}
type Registrator struct {
	Address string `yaml:"address"`
	T       string `yaml:"type"`
	Auth    Auth   `yaml:"auth"` //TODO register between tenants
}

type Auth struct {
	AK      string `yaml:"accessKey"`
	SK      string `yaml:"secretKey"`
	Project string `yaml:"project"`
}

func ReadYAML() error {
	b, err := ioutil.ReadFile(cmd.CLIParam.ConfPath)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(b, Config); err != nil {
		return err
	}
	return nil
}
