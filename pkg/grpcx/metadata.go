package grpcx

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	MetadataTraceID = "x-trace-id"
	MetadataUserID  = "x-user-id"
)

func GetMetadata(
	ctx context.Context,
	key string,
) string {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	values := md.Get(key)

	if len(values) == 0 {
		return ""
	}

	return values[0]
}
func PutMetadata(
	ctx context.Context,
	key string,
	value string,
) context.Context {
	return metadata.AppendToOutgoingContext(ctx, key, value)
}
