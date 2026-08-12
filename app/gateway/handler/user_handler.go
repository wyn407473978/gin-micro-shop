package handler

import (
	"context"
	"fmt"
	userv1 "gin-micro-shop/api/proto/user/v1"
	userResponse "gin-micro-shop/app/gateway/response"
	"gin-micro-shop/pkg/response"
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
	fmt.Println("user:", user, "err:", err)
	if err != nil {
		response.ErrorWithCode(gin, 500, "requestId", "远程调用失败")
		return
	}
	response.SuccessWithData(gin, user, "requestId")
	return
}

func (h *UserHandler) UserLists(gin *gin.Context) {
	userPageReq := userResponse.UserPageReq{}

	err := gin.ShouldBindJSON(&userPageReq)
	if err != nil {
		response.ErrorWithCode(gin, 400, "requestId", "参数绑定失败")
		fmt.Println("err:", err)
		return
	}
	fmt.Println("userPageReq:", userPageReq)
	users, err := h.UserServiceClient.ListUsers(gin.Request.Context(), &userv1.ListUsersRequest{
		PageNum:  int64(userPageReq.PageNum),
		PageSize: int64(userPageReq.PageSize),
		UserName: &userPageReq.UserName,
	})
	if err != nil {
		response.ErrorWithCode(gin, 500, "requestId", "远程调用失败")
		return
	}
	response.SuccessWithData(gin, users, "requestId")
}

func (h *UserHandler) CreateUser(gin *gin.Context) {
	userCreateReq := userResponse.UserCreateReq{}

	err := gin.ShouldBindJSON(&userCreateReq)
	if err != nil {
		response.ErrorWithCode(gin, 400, "requestId", "参数绑定失败")
		fmt.Println("err:", err)
		return
	}
	fmt.Println("userCreateReq:", userCreateReq)
	createUserResponse, err := h.UserServiceClient.CreateUser(gin.Request.Context(), &userv1.CreateUserRequest{
		Username: userCreateReq.UserName,
		Age:      userCreateReq.Age,
		Sex:      userCreateReq.Sex,
	})
	if err != nil {
		response.ErrorWithCode(gin, 500, "requestId", "远程调用失败")
		return
	}
	response.SuccessWithData(gin, createUserResponse.Success, "requestId")
}
