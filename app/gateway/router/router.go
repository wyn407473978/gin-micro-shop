package router

import (
	orderv1 "gin-micro-shop/api/proto/order/v1"
	productv1 "gin-micro-shop/api/proto/product/v1"
	userv1 "gin-micro-shop/api/proto/user/v1"
	"github.com/gin-gonic/gin"
)

func RouterInit(gin *gin.Engine, userServiceClient userv1.UserServiceClient, productServiceClient productv1.ProductGrpcServiceClient, orderGrpcServiceClient orderv1.OrderGrpcServiceClient) {
	UserRouter(gin, userServiceClient)
	InitProductRouter(gin, productServiceClient)
	OrderRouter(gin, orderGrpcServiceClient)
}
