package main

import (
	"fmt"
	productv1 "gin-micro-shop/api/proto/product/v1"
	"gin-micro-shop/app/product/config"
	"gin-micro-shop/app/product/internal/grpc_service"
	"gin-micro-shop/app/product/internal/repository"
	"gin-micro-shop/app/product/internal/service"
	"gin-micro-shop/pkg/etcd"
	"gin-micro-shop/pkg/grpcx"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("product")
	config.InitConfig()
	myConfig := config.GetConfig()
	etcdConfig := myConfig.Etcd

	database := myConfig.Database

	db := database.GetDB()

	productGrpc := myConfig.ProductGrpc
	grpcServer := grpcx.NewServer(grpcx.ServerConfig{})
	sprintf := fmt.Sprintf(":%d", productGrpc.Port)
	listen, err := net.Listen("tcp", sprintf)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productv1.RegisterProductGrpcServiceServer(grpcServer, grpc_service.NewProductGrpcService(productService))

	etcdClient, _ := clientv3.New(clientv3.Config{Endpoints: []string{etcdConfig.GetAddress()}, DialTimeout: 10 * time.Second})
	etcdRegistry := etcd.NewEtcdRegistry(etcdClient, 10)
	err = etcdRegistry.Register(&etcd.ServiceInstance{Address: productGrpc.IP, ID: productGrpc.Name, Name: productGrpc.Name, Port: productGrpc.Port})
	if err != nil {
		fmt.Println("etcd注册失败", err)
	}
	fmt.Println("etcd注册成功")
	defer etcdRegistry.Deregister()

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("product Grpc 服务器启动成功")

}
