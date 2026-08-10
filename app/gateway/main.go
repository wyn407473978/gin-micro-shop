package main

import (
	"gin-micro-shop/app/gateway/router"
	"github.com/gin-gonic/gin"
)
import "fmt"

func main() {
	fmt.Println("hello world")
	engine := gin.Default()
	router.RouterInit(engine)
	err := engine.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
