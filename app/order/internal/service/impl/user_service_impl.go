package impl

import (
	"context"
	userv1 "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/order/internal/service/dto"
)

type UserServiceImpl struct {
	userGrpcServiceClient userv1.UserServiceClient
}

func NewUserServiceImpl(userGrpcServiceClient userv1.UserServiceClient) *UserServiceImpl {
	return &UserServiceImpl{
		userGrpcServiceClient: userGrpcServiceClient,
	}
}

func (u *UserServiceImpl) GetUserById(ctx context.Context, id int64) (*dto.UserDto, error) {
	user, err := u.userGrpcServiceClient.GetUser(ctx, &userv1.GetUserRequest{
		Id: id,
	})
	return &dto.UserDto{
		ID:       user.Id,
		UserName: user.Name,
		Age:      user.Age,
	}, err
}
