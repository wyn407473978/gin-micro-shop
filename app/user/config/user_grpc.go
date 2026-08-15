package config

type UserGrpc struct {
	Name       string `yaml:"name" mapstructure:"name"`
	IP         string `yaml:"ip" mapstructure:"ip"`
	Port       int    `yaml:"port" mapstructure:"port"`
	InstanceId string `yaml:"instance_id" mapstructure:"instance_id"`
}
