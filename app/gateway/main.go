package main

import (
	orderv1 "gin-micro-shop/api/proto/order/v1"
	productv1 "gin-micro-shop/api/proto/product/v1"
	pb "gin-micro-shop/api/proto/user/v1"
	//productv1 "gin-micro-shop/api/proto/product/v1"

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

	newClient, err := grpc.NewClient("localhost:50002", grpc.WithInsecure())
	defer newClient.Close()
	productClient := productv1.NewProductGrpcServiceClient(newClient)

	orderClient, err := grpc.NewClient("localhost:50003", grpc.WithInsecure())
	defer orderClient.Close()
	orderGrpcService := orderv1.NewOrderGrpcServiceClient(orderClient)

	router.RouterInit(engine, client, productClient, orderGrpcService)
	serverPort := myConfig.Server.Port
	sprintf := fmt.Sprintf(":%d", serverPort)
	fmt.Printf("server listening at %s\n", sprintf)
	err = engine.Run(sprintf)
	if err != nil {
		fmt.Println(err)
	}
}
