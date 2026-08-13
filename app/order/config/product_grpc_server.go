package config

type ProductGrpc struct {
	Host string `yaml:"host" mapstructure:"host"`
	Port int    `yaml:"port" mapstructure:"port"`
}
