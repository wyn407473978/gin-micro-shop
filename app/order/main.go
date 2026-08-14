package main

import (
	"fmt"
	orderv1 "gin-micro-shop/api/proto/order/v1"
	productv1 "gin-micro-shop/api/proto/product/v1"
	userv1 "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/order/config"
	"gin-micro-shop/app/order/internal/repository"
	impl2 "gin-micro-shop/app/order/internal/service/impl"
	"gin-micro-shop/pkg/etcd/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {

	config.InitConfig()
	myConfig := config.GetConfig()
	etcdConfig := myConfig.Etcd

	database := myConfig.Database
	orderGrpc := myConfig.OrderGrpc
	db := database.GetDB()

	//TODO 把用户服务和Product 注入进来给Service 使用

	userConnect, err2 := grpc.NewClient("localhost:50001", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("did not connect: %v", err2)
	}
	userConnect.Connect()

	userGrpcServiceClient := userv1.NewUserServiceClient(userConnect)
	defer userConnect.Close()

	userService := impl2.NewUserServiceImpl(userGrpcServiceClient)

	productConnect, err2 := grpc.NewClient("localhost:50002", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("did not connect: %v", err2)
	}
	productConnect.Connect()

	//其实注册到这里就可以进行远程调用了，之所以在封装一层是要将不需要的接口隔离开 只暴露需要的接口
	productGrpcServiceClient := productv1.NewProductGrpcServiceClient(productConnect)
	defer productConnect.Close()
	productService := impl2.NewProductServiceImpl(productGrpcServiceClient)

	//注册Order Grpc 服务
	orderRepository := repository.NewOrderRepository(db)

	orderGrpcService := impl2.NewOrderGrpcService(orderRepository, productService, userService)

	server := grpc.NewServer()
	orderv1.RegisterOrderGrpcServiceServer(server, orderGrpcService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", orderGrpc.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	etcdClient, _ := clientv3.New(clientv3.Config{Endpoints: []string{etcdConfig.GetAddress()}, DialTimeout: 10 * time.Second})
	etcdRegistry := registry.NewEtcdRegistry(etcdClient, 10)
	err = etcdRegistry.Register(&registry.ServiceInstance{Address: orderGrpc.IP, ID: orderGrpc.Name, Name: orderGrpc.Name, Port: orderGrpc.Port})
	if err != nil {
		fmt.Println("etcd注册失败", err)
	}
	fmt.Println("etcd注册成功")
	defer etcdRegistry.Deregister()

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	defer lis.Close()

}
