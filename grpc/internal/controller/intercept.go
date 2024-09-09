package controller

import (
	"context"

	"google.golang.org/grpc"
)

func (c *c) interceptor(desc grpc.MethodDesc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp any, err error) {
		m := NewMethod(desc, info)
		return m, nil
	}
}
