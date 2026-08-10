package router

import (
	"gin-micro-shop/app/gateway/handler"
	"github.com/gin-gonic/gin"
)

func UserRouter(gin *gin.Engine) {
	userHandler := &handler.UserHandler{}
	userGroup := gin.Group("/user")

	userGroup.GET("/get", userHandler.GetUser)

}
