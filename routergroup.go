package ginger

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger/internal/router"
)

func (s *server) NewRouterGroup(path string) gateway.RouterGroup {
	g := router.NewGroup(
		&s.engine.RouterGroup,
		path,
		&s.config.Router,
		s.controller,
	)
	return g
}
