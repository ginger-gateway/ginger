package request

import (
	"github.com/gin-gonic/gin"
	"github.com/ginger-core/gateway"
)

func (r *request) WithConn(conn any) gateway.Request {
	r.Context = conn.(*gin.Context)
	return r
}

func (r *request) GetConn() any {
	return r.Context
}
