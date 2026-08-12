package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var config Config

type Config struct {
	Server   Server   `yaml:"server" mapstructure:"server"`
	Database DataBase `yaml:"database" mapstructure:"database"`
	UserGrpc UserGrpc `yaml:"grpc" mapstructure:"grpc"`
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("app/user/config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = viper.Unmarshal(&config)
	fmt.Printf("%+v\n", config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("配置初始化完成！")
}

func GetConfig() Config {
	return config
}
