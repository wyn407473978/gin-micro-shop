package grpcx

import (
	"gin-micro-shop/pkg/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

const roundRobinServiceConfig = `
{
    "loadBalancingConfig": [
        {
            "round_robin": {}
        }
    ]
}
`

type ClientConfig struct {
	target  string
	timeout time.Duration
}

func NewClient(target, round string, builder *resolver.EtcdResolverBuilder, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
	defaultOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithResolvers(builder),
		grpc.WithChainUnaryInterceptor(LogClientInterceptor(), Timeout(10*time.Second)),
	}

	if round == "round_robin" {
		defaultOpts = append(defaultOpts, grpc.WithDefaultServiceConfig(roundRobinServiceConfig))
	}

	defaultOpts = append(defaultOpts, opts...)
	return grpc.NewClient(target, defaultOpts...)
}
