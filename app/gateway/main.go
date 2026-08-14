package main

import (
	orderv1 "gin-micro-shop/api/proto/order/v1"
	productv1 "gin-micro-shop/api/proto/product/v1"
	pb "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/pkg/etcd/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"

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
	etcdConfig := myConfig.Etcd
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
	gatewayServer := myConfig.Server
	serverPort := gatewayServer.Port
	sprintf := fmt.Sprintf(":%d", serverPort)
	fmt.Printf("server listening at %s\n", sprintf)

	etcdClient, _ := clientv3.New(clientv3.Config{Endpoints: []string{etcdConfig.GetAddress()}, DialTimeout: 10 * time.Second})
	etcdRegistry := registry.NewEtcdRegistry(etcdClient, 10)
	err = etcdRegistry.Register(&registry.ServiceInstance{Address: gatewayServer.IP, ID: gatewayServer.Name, Name: gatewayServer.Name, Port: gatewayServer.Port})
	if err != nil {
		fmt.Println("etcd注册失败", err)
	}
	fmt.Println("etcd注册成功")
	defer etcdRegistry.Deregister()

	err = engine.Run(sprintf)
	if err != nil {
		fmt.Println(err)
	}
}
