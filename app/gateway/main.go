package main

import (
	pb "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/gateway/config"
	"gin-micro-shop/app/gateway/router"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)
import "fmt"

func main() {
	config.InitConfig()
	myConfig := config.GetConfig()
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	engine := gin.Default()
	router.RouterInit(engine, client)
	serverPort := myConfig.Server.Port
	sprintf := fmt.Sprintf(":%d", serverPort)
	fmt.Printf("server listening at %s\n", sprintf)
	err = engine.Run(sprintf)
	if err != nil {
		fmt.Println(err)
	}
}
