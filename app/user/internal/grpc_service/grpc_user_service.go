package grpc_service

import (
	"context"
	"fmt"
	pb "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/user/internal/service"
	"strconv"
)

type GrpcUserServiceImpl struct {
	userService *service.UserService
	pb.UnimplementedUserServiceServer
}

func NewGrpcUserService(userService *service.UserService) *GrpcUserServiceImpl {
	return &GrpcUserServiceImpl{userService: userService}
}
func (s *GrpcUserServiceImpl) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	//id, _ := s.userService.GetUserById(req.Id)
	user, err := s.userService.FindUserByID(ctx, int(req.Id))
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return &pb.User{
		Id:   int64(user.ID),
		Name: user.UserName,
		Age:  strconv.FormatInt(int64(user.Age), 10),
	}, nil
}

func (s *GrpcUserServiceImpl) ListUsers(ctx context.Context, listUserReq *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	//req := response.UserPageReq{}
	//fmt.Println("req:", listUserReq)
	//req.PageNum = int(listUserReq.PageNum)
	//req.PageSize = int(listUserReq.PageSize)
	//users, _ := s.userHandler.GetUsers(req)
	//fmt.Println(users)
	//users2 := []*pb.User{{Id: 1, Name: "John", Age: "30"}}
	return &pb.ListUsersResponse{
		Users:         nil,
		NextPageToken: "next_page_token",
	}, nil
}
