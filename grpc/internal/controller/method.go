package controller

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"google.golang.org/grpc"
)

type Method interface {
	GetFullName() string
	RegisterHandlers(...gateway.Handler)
	Handle(request gateway.Request) (any, errors.Error)
}

type method struct {
	desc     grpc.MethodDesc
	info     *grpc.UnaryServerInfo
	handlers []gateway.Handler
}

type MethodHandler func(in any) (any, error)

func NewMethod(desc grpc.MethodDesc, info *grpc.UnaryServerInfo) Method {
	return &method{
		desc: desc,
		info: info,
	}
}

func (m *method) GetFullName() string {
	return m.info.FullMethod
}

func (m *method) RegisterHandlers(hs ...gateway.Handler) {
	m.handlers = hs
}

func (m *method) Handle(request gateway.Request) (r any, err errors.Error) {
	for _, h := range m.handlers {
		r, err = h.Handle(request)
		if err != nil {
			return nil, err
		}
	}
	return r, err
}
