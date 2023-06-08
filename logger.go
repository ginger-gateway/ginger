package ginger

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log/logger"
	"github.com/ginger-gateway/ginger/internal"
)

func (s *server) newLoggerHandler(l logger.Logger) gin.HandlerFunc {
	l = l.WithTrace("handler")

	skipPaths := make(map[string]bool, 0)
	for _, sp := range s.config.Logger.SkipPaths {
		skipPaths[sp] = true
	}
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		exists := skipPaths[c.FullPath()]
		if exists {
			return
		}

		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		_l := l.WithTrace("ginger")

		var request gateway.Request
		if r, _ := c.Get(internal.RequestKey); r != nil {
			request = r.(gateway.Request)
		}
		if request != nil {
			_l = _l.WithUid(request.GetId())
		}

		status := c.Writer.Status()

		field := logger.Field{
			"Path":    path,
			"Method":  c.Request.Method,
			"IP":      c.ClientIP(),
			"Elapsed": time.Since(start).Milliseconds(),
			"Status":  status,
		}
		if len(c.Errors) != 0 {
			field["Errors"] = c.Errors.ByType(gin.ErrorTypePrivate)
		}
		if status < 400 {
			_l.With(field).Debugf("")
		} else if status < 500 {
			if status == http.StatusFailedDependency {
				_l.With(field).Criticalf("StatusFailedDependency")
			} else {
				_l.With(field).Infof("")
			}
		} else {
			_l.With(field).Criticalf("")
		}
	}
}
