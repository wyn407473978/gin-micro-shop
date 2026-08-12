package main

import (
	"fmt"
	pb "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/user/config"
	"gin-micro-shop/app/user/internal/grpc_service"
	"gin-micro-shop/app/user/internal/repository"
	"gin-micro-shop/app/user/internal/service"
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
	userRepository := repository.NewUserRepository(myConfig.Database.GetDB())
	userService := service.NewUserService(userRepository)
	grpcUserService := grpc_service.NewGrpcUserService(userService)
	pb.RegisterUserServiceServer(grpcServer, grpcUserService)
	//采用一个goroutine 来启动 grpc 服务 防止阻塞导致下面的打印信息无法输出
	if err := grpcServer.Serve(listen); err != nil {
		fmt.Println("grpc服务启动失败", err)
	}

}
