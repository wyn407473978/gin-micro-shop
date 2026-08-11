package grpc_service

import (
	"context"
	"fmt"
	pb "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/user/internal/handler"
	"gin-micro-shop/pkg/response"
)

type GrpcUserServiceImpl struct {
	UserHandler handler.UserHandler
	pb.UnimplementedUserServiceServer
}

func (s *GrpcUserServiceImpl) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	id, _ := s.UserHandler.GetUserById(req.Id)
	return &pb.User{
		Id:   id,
		Name: "John",
		Age:  "30",
	}, nil
}

func (s *GrpcUserServiceImpl) ListUsers(ctx context.Context, listUserReq *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	req := response.UserPageReq{}
	fmt.Println("req:", listUserReq)
	req.PageNum = int(listUserReq.PageNum)
	req.PageSize = int(listUserReq.PageSize)
	users, _ := s.UserHandler.GetUsers(req)
	fmt.Println(users)
	users2 := []*pb.User{{Id: 1, Name: "John", Age: "30"}}
	return &pb.ListUsersResponse{
		Users:         users2,
		NextPageToken: "next_page_token",
	}, nil
}
