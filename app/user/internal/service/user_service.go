package service

import (
	"context"
	pb "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/gateway/response"
	"gin-micro-shop/app/user/internal/model"
	"gin-micro-shop/app/user/internal/repository"
	"strconv"
	"time"
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

func (us *UserService) GetPageUsers(ctx context.Context, req response.UserPageReq) (grpcResp *pb.ListUsersResponse, err error) {
	users, err := us.userRepository.PageUserByName(ctx, req)
	if err != nil {
		return nil, err
	}
	var grpcUsers []*pb.User
	for _, user := range users {
		grpcUsers = append(grpcUsers, &pb.User{
			Id:   int64(user.ID),
			Name: user.Username,
			Age:  strconv.FormatInt(int64(user.Age), 10),
		})
	}
	return &pb.ListUsersResponse{
		Users:         grpcUsers,
		NextPageToken: "next_page_token",
	}, nil
}

func (us *UserService) CreateUser(ctx context.Context, username string, age int32, sex int32) (bool, error) {
	user := model.User{
		Username:   username,
		Password:   "123456",
		Age:        int(age),
		Sex:        int(sex),
		CreateTime: time.Now(),
	}
	return us.userRepository.CreateUser(ctx, user)

}
