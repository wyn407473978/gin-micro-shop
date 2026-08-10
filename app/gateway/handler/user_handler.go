package handler

import "github.com/gin-gonic/gin"

type UserHandler struct{}

func (h *UserHandler) GetUser(gin *gin.Context) {
	gin.JSON(200, "welcome to go")
	return
}
