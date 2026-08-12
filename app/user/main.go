package main

import (
	"fmt"
	pb "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/user/config"
	"gin-micro-shop/app/user/internal/grpc_service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	config.InitConfig()
	myConfig := config.GetConfig()
	userGrpc := myConfig.UserGrpc
	grpcServer := grpc.NewServer()
	grpcPort := fmt.Sprintf(":%d", userGrpc.Port)
	fmt.Println("grpc服务端口", grpcPort)
	listen, _ := net.Listen("tcp", grpcPort)
	pb.RegisterUserServiceServer(grpcServer, &grpc_service.GrpcUserServiceImpl{})
	//采用一个goroutine 来启动 grpc 服务 防止阻塞导致下面的打印信息无法输出
	//go func() {
	if err := grpcServer.Serve(listen); err != nil {
		fmt.Println("grpc服务启动失败", err)
	}
	//}()

	//engine := gin.Default()
	//userServicePort := fmt.Sprintf(":%d", server.Port)
	//fmt.Println("用户服务模块端口", userServicePort)
	//err := engine.Run(userServicePort)
	//if err != nil {
	//	fmt.Println("用户服务模块启动失败", err)
	//	panic(err)
	//}
}
