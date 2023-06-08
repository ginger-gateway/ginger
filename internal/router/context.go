package router

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ginger-core/gateway"
)

func (r *router) procContext(c *gin.Context,
	ctx context.Context) context.Context {
	ip := c.GetHeader("X-Forwarded-For")
	ip = strings.TrimSpace(strings.Split(ip, ",")[0])
	if ip == "" {
		ip = strings.TrimSpace(c.GetHeader("X-Real-Ip"))
	}
	if ip == "" {
		ip = strings.TrimSpace(strings.Split(c.Request.RemoteAddr, ":")[0])
	}
	ctx = context.WithValue(ctx, gateway.IPKey, ip)
	ctx = context.WithValue(ctx, gateway.AgentKey, c.GetHeader("User-Agent"))

	ctx, cancel := context.WithTimeout(ctx, r.config.ConnTimeout)

	c.Set("cancel", cancel)

	return ctx
}
