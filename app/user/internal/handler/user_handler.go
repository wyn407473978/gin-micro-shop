package handler

import (
	"fmt"
	"gin-micro-shop/app/user/internal/model"
	"gin-micro-shop/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func (h *UserHandler) GetUserById(id int64) (int64, error) {
	return id, nil
}
func (h *UserHandler) GetUsers(gin *gin.Context) {
	userPageReq := response.UserPageReq{}
	err := gin.ShouldBindJSON(userPageReq)
	if err != nil {
		gin.JSON(400, err)
		return
	}
	fmt.Println(userPageReq)
	users := []model.User{
		{ID: 1, Name: "Alice", Age: 18, Sex: "Female"},
		{ID: 2, Name: "Bob", Age: 20, Sex: "Male"},
	}
	gin.JSON(200, users)
}
