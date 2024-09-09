package grpc

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger/grpc/internal/response"
)

func (s *srv) NewResponder() gateway.Responder {
	return response.NewResponder()
}
