package response

import (
	"github.com/ginger-core/gateway"
)

type responder struct {
}

func NewResponder() gateway.Responder {
	r := &responder{}
	return r
}
