package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ginger-core/gateway"
)

type group struct {
	rg *gin.RouterGroup
	router
}

func NewGroup(source *gin.RouterGroup, path string,
	config *Config, controller gateway.Controller) gateway.RouterGroup {
	rg := &group{
		rg: source.Group(path),
		router: router{
			config:     config,
			controller: controller,
		},
	}
	return rg
}

func (g *group) Group(path string) gateway.RouterGroup {
	return NewGroup(
		g.rg,
		path,
		g.config,
		g.controller,
	)
}

func (g *group) RegisterMiddlewares(middlewares ...gateway.Handler) {
	g.rg.Use(g.getGinHandlerFuncArr(middlewares...)...)
}

func (g *group) On(method gateway.Method, handlers ...gateway.Handler) {
	g.OnPath(method, "", handlers...)
}

func (g *group) OnPath(method gateway.Method,
	path string, handlers ...gateway.Handler) {
	switch method {
	case gateway.Create:
		g.rg.POST(path, g.getGinHandlerFuncArr(handlers...)...)
	case gateway.Read:
		g.rg.GET(path, g.getGinHandlerFuncArr(handlers...)...)
	case gateway.Update:
		g.rg.PUT(path, g.getGinHandlerFuncArr(handlers...)...)
	case gateway.Delete:
		g.rg.DELETE(path, g.getGinHandlerFuncArr(handlers...)...)
	default:
		panic("method handler not found")
	}
}
