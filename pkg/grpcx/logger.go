package grpcx

import (
	"context"
	"fmt"
	"gin-micro-shop/app/gateway/middleware"
	"google.golang.org/grpc"
	"log"
	"time"
)

func LogClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req any,
		reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		start := time.Now()
		value := ctx.Value(middleware.RequestID).(string)
		fmt.Println("Request ID:", value)
		ctx = PutMetadata(ctx, MetadataTraceID, value)
		err := invoker(ctx, method, req, reply, cc, opts...)
		latency := time.Since(start)
		log.Printf("client grpc request: %s, %s, %v", method, value, latency)
		return err
	}
}

func LogServerInterceptor() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		start := time.Now()
		resp, err = handler(ctx, req)
		latency := time.Since(start)
		log.Printf("server grpc request: %s, %s, %v", info.FullMethod, GetMetadata(ctx, MetadataTraceID), latency)
		return resp, err
	}

}
