package grpcx

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

func Timeout(
	timeout time.Duration,
) grpc.UnaryClientInterceptor {

	return func(
		ctx context.Context,
		method string,
		req any,
		reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {

		// 如果上层已经设置 Deadline，
		// 不要粗暴覆盖。
		if _, ok := ctx.Deadline(); ok {
			return invoker(
				ctx,
				method,
				req,
				reply,
				cc,
				opts...,
			)
		}

		ctx, cancel := context.WithTimeout(
			ctx,
			timeout,
		)
		defer cancel()

		return invoker(
			ctx,
			method,
			req,
			reply,
			cc,
			opts...,
		)
	}
}
