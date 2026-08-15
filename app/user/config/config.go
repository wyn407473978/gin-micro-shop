package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var config Config

type Config struct {
	Server   Server   `yaml:"server" mapstructure:"server"`
	Database DataBase `yaml:"database" mapstructure:"database"`
	UserGrpc UserGrpc `yaml:"grpc" mapstructure:"grpc"`
	Etcd     Etcd     `yaml:"etcd" mapstructure:"etcd"`
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
	if name := os.Getenv("USER_GRPC_NAME"); name != "" {
		viper.Set("grpc.name", name)
	}
	if instanceID := os.Getenv("USER_GRPC_INSTANCE_ID"); instanceID != "" {
		viper.Set("grpc.instance_id", instanceID)
	}
	if port := os.Getenv("USER_GRPC_PORT"); port != "" {
		viper.Set("grpc.port", port)
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
