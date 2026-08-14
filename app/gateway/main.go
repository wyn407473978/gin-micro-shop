package main

import (
	orderv1 "gin-micro-shop/api/proto/order/v1"
	productv1 "gin-micro-shop/api/proto/product/v1"
	pb "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/pkg/etcd/discovery"
	"gin-micro-shop/pkg/etcd/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"net"
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
	etcdClient, _ := clientv3.New(clientv3.Config{Endpoints: []string{etcdConfig.GetAddress()}, DialTimeout: 10 * time.Second})
	serviceDiscovery := discovery.NewDiscovery(etcdClient)

	err, userConnect := serviceDiscoveryMethod(serviceDiscovery, "user-service-grpc")
	userClient := pb.NewUserServiceClient(userConnect)
	defer userConnect.Close()

	err, productConnect := serviceDiscoveryMethod(serviceDiscovery, "product-service-grpc")
	productClient := productv1.NewProductGrpcServiceClient(productConnect)
	defer productConnect.Close()

	err, orderConnect := serviceDiscoveryMethod(serviceDiscovery, "order-service-grpc")
	orderClient := orderv1.NewOrderGrpcServiceClient(orderConnect)
	defer orderConnect.Close()

	engine := gin.Default()

	router.RouterInit(engine, userClient, productClient, orderClient)
	gatewayServer := myConfig.Server
	serverPort := gatewayServer.Port
	sprintf := fmt.Sprintf(":%d", serverPort)
	fmt.Printf("server listening at %s\n", sprintf)

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

func serviceDiscoveryMethod(serviceDiscovery *discovery.Discovery, name string) (error, *grpc.ClientConn) {
	if err := serviceDiscovery.WatchService(name); err != nil {
		log.Fatalf("watch user service failed: %v", err)
	}
	instances, err := serviceDiscovery.GetInstances(name)
	if err != nil || len(instances) == 0 {
		log.Fatal("no available service-grpc instance")
	}

	// 先用第一个实例；后续可改成轮询/随机负载均衡
	userInstance := instances[0]
	userAddr := net.JoinHostPort(userInstance.Address, fmt.Sprint(userInstance.Port))

	userConnect, err := grpc.NewClient(userAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect user service failed: %v", err)
	}
	return err, userConnect
}
