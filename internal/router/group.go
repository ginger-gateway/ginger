package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ginger-core/gateway"
)

type group struct {
	rg *gin.RouterGroup
	router
}

func NewGroup(source *gin.RouterGroup, path string, config *Config, controller gateway.Controller) gateway.RouterGroup {
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

func (g *group) Create(path string, handlers ...gateway.Handler) {
	g.rg.POST(path, g.getGinHandlerFuncArr(handlers...)...)
}

func (g *group) Read(path string, handlers ...gateway.Handler) {
	g.rg.GET(path, g.getGinHandlerFuncArr(handlers...)...)
}

func (g *group) Update(path string, handlers ...gateway.Handler) {
	g.rg.PUT(path, g.getGinHandlerFuncArr(handlers...)...)
}

func (g *group) Delete(path string, handlers ...gateway.Handler) {
	g.rg.DELETE(path, g.getGinHandlerFuncArr(handlers...)...)
}
