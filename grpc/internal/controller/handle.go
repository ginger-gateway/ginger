package controller

import (
	"context"

	"github.com/ginger-core/errors"
	grpcerr "github.com/ginger-core/errors/grpc"
	"github.com/ginger-gateway/ginger/grpc/internal/request"
	"google.golang.org/grpc"
)

func (c *c) Handle(ctx context.Context, req any, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp any, err error) {
	method := c.methods[info.FullMethod]
	if method == nil {
		return nil, grpcerr.Generate(
			errors.New().WithTrace("methodNotFound").
				WithDetail(errors.NewDetail().
					With("method", info.FullMethod)))
	}
	request := request.New(nil, c.GetLanguageBundle(), req).
		WithContext(ctx)
	r, err := method.Handle(request)
	if err != nil {
		return nil, grpcerr.Parse(err)
	}
	return r, nil
}
