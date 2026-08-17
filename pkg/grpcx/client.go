package grpcx

import (
	"gin-micro-shop/pkg/etcd"
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

func NewClient(target, round string, builder *etcd.EtcdResolverBuilder, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
	defaultOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithResolvers(builder),
	}

	if round == "round_robin" {
		defaultOpts = append(defaultOpts, grpc.WithDefaultServiceConfig(roundRobinServiceConfig))
	}

	defaultOpts = append(defaultOpts, opts...)
	return grpc.NewClient(target, defaultOpts...)
}
