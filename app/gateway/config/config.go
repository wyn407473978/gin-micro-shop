package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var config Config

type Config struct {
	Server Server `yaml:"server" mapstructure:"server"`
	Etcd   Etcd   `yaml:"etcd" mapstructure:"etcd"`
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("app/gateway/")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("%+v\n", config)
}

func GetConfig() Config {
	return config
}
