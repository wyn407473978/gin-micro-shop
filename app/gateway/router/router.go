package router

import (
	orderv1 "gin-micro-shop/api/proto/order/v1"
	productv1 "gin-micro-shop/api/proto/product/v1"
	userv1 "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/gateway/middleware"
	"github.com/gin-gonic/gin"
)

func RouterInit(gin *gin.Engine, userServiceClient userv1.UserServiceClient, productServiceClient productv1.ProductGrpcServiceClient, orderGrpcServiceClient orderv1.OrderGrpcServiceClient) {
	gin.Use(middleware.RequestInterceptor())
	UserRouter(gin, userServiceClient)
	InitProductRouter(gin, productServiceClient)
	OrderRouter(gin, orderGrpcServiceClient)
}
