package handler

import (
	"context"
	"fmt"
	userv1 "gin-micro-shop/api/proto/user/v1"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserHandler struct {
	UserServiceClient userv1.UserServiceClient
}

func (h *UserHandler) GetUser(gin *gin.Context) {
	id := gin.Query("id")
	userId, _ := strconv.Atoi(id)
	fmt.Println(id)
	user, err := h.UserServiceClient.GetUser(context.Background(), &userv1.GetUserRequest{Id: int64(userId)})
	if err != nil {
		gin.JSON(500, err)
		return
	}
	gin.JSON(200, user)
	return
}
