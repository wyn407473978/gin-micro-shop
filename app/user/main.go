package main

import (
	"fmt"
	pb "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/user/config"
	"gin-micro-shop/app/user/internal/grpc_service"
	"gin-micro-shop/app/user/internal/repository"
	"gin-micro-shop/app/user/internal/service"
	"gin-micro-shop/pkg/etcd"
	"gin-micro-shop/pkg/grpcx"
	clientv3 "go.etcd.io/etcd/client/v3"
	"net"
	"time"
)

func main() {
	config.InitConfig()
	myConfig := config.GetConfig()

	etcdConfig := myConfig.Etcd

	userGrpc := myConfig.UserGrpc
	grpcServer := grpcx.NewServer(grpcx.ServerConfig{})
	grpcPort := fmt.Sprintf(":%d", userGrpc.Port)
	fmt.Println("grpc服务端口", grpcPort)
	listen, _ := net.Listen("tcp", grpcPort)

	userRepository := repository.NewUserRepository(myConfig.Database.GetDB())
	userService := service.NewUserService(userRepository)
	grpcUserService := grpc_service.NewGrpcUserService(userService)

	etcdClient, _ := clientv3.New(clientv3.Config{Endpoints: []string{etcdConfig.GetAddress()}, DialTimeout: 10 * time.Second})
	etcdRegistry := etcd.NewEtcdRegistry(etcdClient, 10)
	err := etcdRegistry.Register(&etcd.ServiceInstance{Address: userGrpc.IP, ID: userGrpc.InstanceId, Name: userGrpc.Name, Port: userGrpc.Port})
	if err != nil {
		fmt.Println("etcd注册失败", err)
	}
	fmt.Println("etcd注册成功")
	defer etcdRegistry.Deregister()

	pb.RegisterUserServiceServer(grpcServer, grpcUserService)
	//采用一个goroutine 来启动 grpc 服务 防止阻塞导致下面的打印信息无法输出
	if err := grpcServer.Serve(listen); err != nil {
		fmt.Println("grpc服务启动失败", err)
	}

}
