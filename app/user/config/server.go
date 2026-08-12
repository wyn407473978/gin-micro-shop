package config

type Server struct {
	Name string `yaml:"name" mapstructure:"name"`
	Port int    `yaml:"port" mapstructure:"port"`
}
