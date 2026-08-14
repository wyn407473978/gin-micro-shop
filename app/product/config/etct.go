package config

import "fmt"

type Etcd struct {
	IP      string `yaml:"ip" mapstructure:"ip"`
	Port    string `yaml:"port" mapstructure:"port"`
	Timeout int    `yaml:"timeout" mapstructure:"timeout"`
}

func (e *Etcd) GetAddress() string {
	fmt.Println("etcd地址", e.IP+":"+e.Port)
	return e.IP + ":" + e.Port
}
