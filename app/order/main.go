package main

import (
	"fmt"
	orderv1 "gin-micro-shop/api/proto/order/v1"
	productv1 "gin-micro-shop/api/proto/product/v1"
	userv1 "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/order/config"
	"gin-micro-shop/app/order/internal/repository"
	impl2 "gin-micro-shop/app/order/internal/service/impl"
	"gin-micro-shop/pkg/etcd"
	"gin-micro-shop/pkg/grpcx"
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
	etcdClient, _ := clientv3.New(clientv3.Config{Endpoints: []string{etcdConfig.GetAddress()}, DialTimeout: 10 * time.Second})
	//discoveryService := etcd.NewDiscovery(etcdClient)
	builder := etcd.NewEtcdResolverBuilder(etcdClient, "/services")
	userConnect, err2 := grpcx.NewClient("etcd:///user-service-grpc", "round_robin", builder)
	if err2 != nil {
		log.Fatalf("did not connect: %v", err2)
	}
	userConnect.Connect()
	userGrpcServiceClient := userv1.NewUserServiceClient(userConnect)
	defer userConnect.Close()

	userService := impl2.NewUserServiceImpl(userGrpcServiceClient)

	//err2, productConnect := serviceDiscoveryMethod(discoveryService, "product-service-grpc")
	productConnect, err2 := grpcx.NewClient("etcd:///product-service-grpc", "round_robin", builder)
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

	server := grpcx.NewServer(
		grpcx.ServerConfig{},
	)
	orderv1.RegisterOrderGrpcServiceServer(server, orderGrpcService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", orderGrpc.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	etcdRegistry := etcd.NewEtcdRegistry(etcdClient, 10)
	err = etcdRegistry.Register(&etcd.ServiceInstance{Address: orderGrpc.IP, ID: orderGrpc.Name, Name: orderGrpc.Name, Port: orderGrpc.Port})
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

func serviceDiscoveryMethod(serviceDiscovery *etcd.Discovery, name string) (error, *grpc.ClientConn) {
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
