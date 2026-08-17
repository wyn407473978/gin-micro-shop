package grpcx

import (
	"google.golang.org/grpc"
)

type ServerConfig struct {
}

func NewServer(
	config ServerConfig,
	opts ...grpc.ServerOption,
) *grpc.Server {

	defaultOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(Recovery(), LogServerInterceptor()),
	}

	defaultOpts = append(
		defaultOpts,
		opts...,
	)

	return grpc.NewServer(
		defaultOpts...,
	)
}
