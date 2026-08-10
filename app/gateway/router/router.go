package router

import "github.com/gin-gonic/gin"

func RouterInit(gin *gin.Engine) {
	UserRouter(gin)
}
