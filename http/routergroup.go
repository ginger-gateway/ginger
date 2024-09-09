package http

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger/http/internal/router"
)

func (s *srv) NewRouterGroup(path string) gateway.RouterGroup {
	g := router.NewGroup(
		&s.engine.RouterGroup,
		path,
		&s.config.Router,
		s.controller,
	)
	return g
}
