package config

type Server struct {
	IP   string `yaml:"ip" mapstructure:"ip"`
	Name string `yaml:"name" mapstructure:"name" json:"name"`
	Port int    `yaml:"port" mapstructure:"port" json:"port"`
}
