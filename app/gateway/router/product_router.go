package router

import (
	productv1 "gin-micro-shop/api/proto/product/v1"
	"gin-micro-shop/app/gateway/handler"
	"github.com/gin-gonic/gin"
)

func InitProductRouter(gin *gin.Engine, productServiceClient productv1.ProductGrpcServiceClient) {
	productHandler := handler.NewProductHandler(productServiceClient)
	productGroup := gin.Group("/product")
	productGroup.GET("/", productHandler.GetProduct)
	productGroup.POST("/", productHandler.CreateProduct)
	productGroup.POST("/reduce_stock", productHandler.ReduceStock)
}
