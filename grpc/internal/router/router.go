package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger/grpc/internal/controller"
)

type router struct {
	controller controller.Controller
}

func (r *router) getHandlerFunc(handler gateway.Handler,
	isFirstHandler, isLastHandler bool) gin.HandlerFunc {
	panic("not implemented")
}

func (r *router) getHandlerFuncArr(handlers ...gateway.Handler) []gin.HandlerFunc {
	var ginHandlerFuncArr []gin.HandlerFunc
	for i, h := range handlers {
		ginHandlerFuncArr = append(ginHandlerFuncArr,
			r.getHandlerFunc(
				h,
				i == 0,
				i == len(handlers)-1,
			),
		)
	}
	return ginHandlerFuncArr
}
