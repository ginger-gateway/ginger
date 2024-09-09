package http

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger/http/internal/response"
)

func (s *srv) NewResponder() gateway.Responder {
	return response.NewResponder()
}
