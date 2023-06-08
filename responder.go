package ginger

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger/internal/response"
)

func (s *server) NewResponder() gateway.Responder {
	return response.NewResponder()
}
