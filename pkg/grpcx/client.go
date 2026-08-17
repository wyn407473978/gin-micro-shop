package grpcx

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type ClientConfig struct {
	target  string
	timeout time.Duration
}

func NewClient(target string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
	defaultOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	}

	defaultOpts = append(defaultOpts, opts...)
	return grpc.NewClient(target, opts...)
}
