package router

import (
	userv1 "gin-micro-shop/api/proto/user/v1"
	"github.com/gin-gonic/gin"
)

func RouterInit(gin *gin.Engine, userServiceClient userv1.UserServiceClient) {
	UserRouter(gin, userServiceClient)
}
