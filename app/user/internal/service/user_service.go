package service

import (
	"context"
	"gin-micro-shop/app/user/internal/model"
	"gin-micro-shop/app/user/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) FindUserByID(ctx context.Context, id int) (user model.User, err error) {
	return us.userRepository.FindUserByID(ctx, id)
}
