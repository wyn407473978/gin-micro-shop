package handler

import (
	"fmt"
	"gin-micro-shop/app/user/internal/model"
	"gin-micro-shop/pkg/response"
)

type UserHandler struct{}

func (h *UserHandler) GetUserById(id int64) (int64, error) {
	fmt.Println("id:", id)
	return id, nil
}
func (h *UserHandler) GetUsers(req response.UserPageReq) ([]model.User, error) {
	fmt.Println(req)
	users := []model.User{
		{ID: 1, Name: "Alice", Age: 18, Sex: "Female"},
		{ID: 2, Name: "Bob", Age: 20, Sex: "Male"},
	}
	return users, nil
}
