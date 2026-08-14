package config

type ProductGrpc struct {
	IP   string `yaml:"ip" mapstructure:"ip"`
	Name string `yaml:"name" mapstructure:"name"`
	Port int    `yaml:"port" mapstructure:"port"`
}
