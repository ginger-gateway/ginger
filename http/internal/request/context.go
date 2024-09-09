package request

import (
	"context"

	"github.com/ginger-core/gateway"
)

func (r *request) WithContext(ctx context.Context) gateway.Request {
	r.context = context.WithValue(ctx, gateway.RequestIdKey, r.id)
	return r
}

func (r *request) GetContext() context.Context {
	return r.context
}
