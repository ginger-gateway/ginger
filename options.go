package ginger

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *server) options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "POST,GET,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}
