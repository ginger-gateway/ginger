package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger/internal"
	"github.com/ginger-gateway/ginger/internal/request"
)

type router struct {
	config     *Config
	controller gateway.Controller
}

func (r *router) getGinHandlerFunc(handler gateway.Handler,
	isFirstHandler, isLastHandler bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req gateway.Request

		defer func() {
			if isLastHandler {
				cancel, ok := c.Get(internal.CancelKey)
				if ok {
					defer cancel.(context.CancelFunc)()
				}
			}
		}()

		if isFirstHandler {
			var language gateway.Language
			if bundle := r.controller.GetLanguageBundle(); bundle != nil {
				acceptLanguage := c.GetHeader("Accept-Language")
				language = gateway.NewLanguage(bundle, acceptLanguage, "EN")
			}
			req = request.New(c, language).
				WithContext(r.procContext(c, c.Request.Context()))

			c.Set(internal.RequestKey, req)
		} else {
			if _req, exists := c.Get(internal.RequestKey); exists {
				req = _req.(gateway.Request)
			} else {
				c.AbortWithStatus(http.StatusNotImplemented)
				isLastHandler = true
				return
			}
		}
		if r.controller.Process(handler, req, isLastHandler) {
			c.Next()
		} else {
			c.Abort()
		}
	}
}

func (r *router) getGinHandlerFuncArr(handlers ...gateway.Handler) []gin.HandlerFunc {
	var ginHandlerFuncArr []gin.HandlerFunc
	for i, h := range handlers {
		ginHandlerFuncArr = append(ginHandlerFuncArr,
			r.getGinHandlerFunc(
				h,
				i == 0,
				i == len(handlers)-1,
			),
		)
	}
	return ginHandlerFuncArr
}
