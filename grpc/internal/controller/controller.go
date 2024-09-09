package controller

import (
	"context"
	"fmt"

	"github.com/ginger-core/gateway"
	"google.golang.org/grpc"
)

type Controller interface {
	gateway.Controller
	RegisterService(desc *grpc.ServiceDesc)
	RegisterHandlers(path string, handlers ...gateway.Handler)
	Handle(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error)
}

type c struct {
	gateway.Controller
	methods map[string]Method
}

func New(responder gateway.Responder) Controller {
	return &c{
		Controller: gateway.NewController(responder),
		methods:    make(map[string]Method),
	}
}

func (c *c) RegisterService(desc *grpc.ServiceDesc) {
	for _, m := range desc.Methods {
		info, err := m.Handler(
			nil,
			context.Background(),
			func(in any) error {
				return nil
			},
			c.interceptor(m),
		)
		if err != nil {
			panic(err)
		}
		method := info.(Method)
		c.methods[method.GetFullName()] = method
	}
}

func (c *c) RegisterHandlers(path string, handlers ...gateway.Handler) {
	m := c.methods[path]
	if m == nil {
		panic(fmt.Sprintf("method of path `%s` not found", path))
	}
	m.RegisterHandlers(handlers...)
}
