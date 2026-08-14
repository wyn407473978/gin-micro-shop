package _interface

import (
	"context"
	"gin-micro-shop/app/order/internal/service/dto"
)

type UserService interface {
	GetUserById(ctx context.Context, id int64) (*dto.UserDto, error)
}
