package config

type Server struct {
	Name string `yaml:"name" mapstructure:"name" json:"name"`
	Port int    `yaml:"port" mapstructure:"port" json:"port"`
}
