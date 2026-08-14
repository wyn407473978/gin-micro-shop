package config

type UserGrpc struct {
	Name string `yaml:"name" mapstructure:"name"`
	IP   string `yaml:"ip" mapstructure:"ip"`
	Port int    `yaml:"port" mapstructure:"port"`
}
