package main

import (
	"fmt"
	"gin-micro-shop/app/product/config"
)

func main() {
	fmt.Println("product")
	config.InitConfig()
	myConfig := config.GetConfig()

	database := myConfig.Database

	db := database.GetDB()

}
