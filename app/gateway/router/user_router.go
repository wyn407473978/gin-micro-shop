package router

import (
	userv1 "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/gateway/handler"
	"github.com/gin-gonic/gin"
)

func UserRouter(gin *gin.Engine, userServiceClient userv1.UserServiceClient) {
	userHandler := &handler.UserHandler{
		UserServiceClient: userServiceClient,
	}
	userGroup := gin.Group("/user")

	userGroup.GET("/get", userHandler.GetUser)

}
