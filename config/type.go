package config

type Configuration struct {
	Source Source `yaml:"source"`
	Target Target `yaml:"target"`
	Auth   Auth   `yaml:"auth"`
}

type Source struct {
	Address       string `yaml:"address"`
	Auth          Auth   `yaml:"auth"`    //TODO register between tenants
	Exclude       string `yaml:"exclude"` // service names, separated by commas
	FetchInterval string `yaml:"fetchInterval"`
}
type Target struct {
	Address           string `yaml:"address"`
	Auth              Auth   `yaml:"auth"` //TODO register between tenants
	HeartbeatInterval string `yaml:"heartbeatInterval"`
}

type Auth struct {
	AK      string `yaml:"accessKey"`
	SK      string `yaml:"secretKey"`
	Project string `yaml:"project"`
}
