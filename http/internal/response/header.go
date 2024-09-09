package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ginger-core/gateway"
)

func (r *responder) SetHeader(request gateway.Request, key string, value any) {
	c := request.GetConn().(*gin.Context)
	c.Header(key, fmt.Sprint(value))
}
