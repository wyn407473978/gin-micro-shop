package main

import (
	pb "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/gateway/router"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)
import "fmt"

func main() {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	engine := gin.Default()
	router.RouterInit(engine, client)
	err = engine.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
