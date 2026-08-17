package grpc_service

import (
	"context"
	"fmt"
	pb "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/gateway/response"
	"gin-micro-shop/app/user/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	user, err := s.userService.FindUserByID(ctx, int(req.Id))
	if err != nil {
		err := status.Error(codes.NotFound, "user not found")
		return nil, err
	}
	return &pb.User{
		Id:   int64(user.ID),
		Name: user.Username,
		Age:  strconv.FormatInt(int64(user.Age), 10),
	}, nil
}

func (s *GrpcUserServiceImpl) ListUsers(ctx context.Context, listUserReq *pb.ListUsersRequest) (grpcResp *pb.ListUsersResponse, err error) {
	req := response.UserPageReq{}
	fmt.Println("req:", listUserReq)
	req.PageNum = int(listUserReq.PageNum)
	req.PageSize = int(listUserReq.PageSize)
	req.UserName = *listUserReq.UserName
	grpcResp, _ = s.userService.GetPageUsers(ctx, req)
	return grpcResp, nil
}

func (s *GrpcUserServiceImpl) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	success, err := s.userService.CreateUser(ctx, req.Username, req.Age, req.Sex)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &pb.CreateUserResponse{Success: success}, nil
}
