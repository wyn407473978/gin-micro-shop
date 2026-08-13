package config

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
import "gorm.io/driver/postgres"

type DataBase struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Port     int    `yaml:"port" mapstructure:"port"`
	UserName string `yaml:"username" mapstructure:"username"`
	Password string `yaml:"password" mapstructure:"password"`
	Database string `yaml:"database" mapstructure:"database"`
}

func (d DataBase) getDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.UserName, d.Password, d.Database,
	)
}
func (d DataBase) GetDB() *gorm.DB {
	dB, err := gorm.Open(postgres.Open(d.getDSN()), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return dB
}
