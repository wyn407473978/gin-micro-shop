package main

import (
	pb "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/user/internal/grpc_service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	listen, _ := net.Listen("tcp", ":50001")
	pb.RegisterUserServiceServer(grpcServer, &grpc_service.GrpcUserServiceImpl{})

	_ = grpcServer.Serve(listen)
	engine := gin.Default()
	engine.Run(":18081")
}
