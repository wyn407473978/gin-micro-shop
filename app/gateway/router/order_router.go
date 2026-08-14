package router

import (
	orderv1 "gin-micro-shop/api/proto/order/v1"
	"gin-micro-shop/app/gateway/handler"
	"github.com/gin-gonic/gin"
)

func OrderRouter(engine *gin.Engine, orderGrpcService orderv1.OrderGrpcServiceClient) {
	orderHandler := handler.NewOrderHandler(orderGrpcService)
	orderGroup := engine.Group("/order")
	orderGroup.GET("/", orderHandler.GetOrderById)
	orderGroup.POST("/", orderHandler.CreateOrder)
}
