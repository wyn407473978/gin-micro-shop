package main

import (
	"fmt"
	productv1 "gin-micro-shop/api/proto/product/v1"
	"gin-micro-shop/app/product/config"
	"gin-micro-shop/app/product/internal/grpc_service"
	"gin-micro-shop/app/product/internal/repository"
	"gin-micro-shop/app/product/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("product")
	config.InitConfig()
	myConfig := config.GetConfig()
	database := myConfig.Database

	db := database.GetDB()

	productGrpc := myConfig.ProductGrpc
	grpcServer := grpc.NewServer()
	sprintf := fmt.Sprintf(":%d", productGrpc.Port)
	listen, err := net.Listen("tcp", sprintf)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productv1.RegisterProductGrpcServiceServer(grpcServer, grpc_service.NewProductGrpcService(productService))

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("product Grpc 服务器启动成功")

}
