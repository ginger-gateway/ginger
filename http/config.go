package http

import "github.com/ginger-gateway/ginger/http/internal/router"

type config struct {
	Logger struct {
		SkipPaths []string
	}
	ListenAddr      string
	Router          router.Config
	RemoteIPHeaders []string
}

func (c *config) initialize() {
	if c.RemoteIPHeaders == nil {
		c.RemoteIPHeaders = []string{"X-Real-IP", "X-Forwarded-For"}
	}
}
