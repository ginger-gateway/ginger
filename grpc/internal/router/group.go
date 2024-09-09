package router

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger/grpc/internal/controller"
	"google.golang.org/grpc"
)

type RouterGroup interface {
	gateway.RouterGroup
	Initialize(desc *grpc.ServiceDesc)
	GetDesc() *grpc.ServiceDesc
}

type group struct {
	router
	path string
	desc *grpc.ServiceDesc
}

func NewGroup(path string, controller controller.Controller) RouterGroup {
	rg := &group{
		path: path,
		router: router{
			controller: controller,
		},
	}
	return rg
}

func (g *group) Initialize(desc *grpc.ServiceDesc) {
	g.desc = desc
	g.path = "/" + desc.ServiceName
	g.controller.RegisterService(desc)
}

func (g *group) GetDesc() *grpc.ServiceDesc {
	return g.desc
}

func (g *group) Group(path string) gateway.RouterGroup {
	return NewGroup(
		path,
		g.controller,
	)
}

func (g *group) RegisterMiddlewares(middlewares ...gateway.Handler) {
	panic("not implemented")
}

func (g *group) On(method gateway.Method, handlers ...gateway.Handler) {
	g.OnPath(method, "", handlers...)
}

func (g *group) OnPath(_ gateway.Method,
	path string, handlers ...gateway.Handler) {
	g.router.controller.RegisterHandlers(
		g.path+"/"+path,
		handlers...,
	)
}
