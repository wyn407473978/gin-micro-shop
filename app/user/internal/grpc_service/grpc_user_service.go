package grpc_service

import (
	"context"
	pb "gin-micro-shop/api/proto/user/v1"
	"gin-micro-shop/app/user/internal/handler"
)

type GrpcUserServiceImpl struct {
	pb.UnimplementedUserServiceServer
}

func (s *GrpcUserServiceImpl) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	userHandler := &handler.UserHandler{}
	id, _ := userHandler.GetUserById(req.Id)
	return &pb.User{
		Id:   id,
		Name: "John",
		Age:  "30",
	}, nil
}
