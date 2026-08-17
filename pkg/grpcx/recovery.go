package grpcx

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func Recovery() grpc.UnaryServerInterceptor {

	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (
		resp any,
		err error,
	) {

		defer func() {
			if r := recover(); r != nil {

				defer log.Printf(
					"grpc panic: method=%s panic=%v",
					info.FullMethod,
					r,
				)

				err = status.Error(
					codes.Internal,
					"internal server error",
				)
			}
		}()

		return handler(ctx, req)
	}
}
