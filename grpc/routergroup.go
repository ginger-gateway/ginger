package grpc

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger/grpc/internal/router"
)

func (s *srv) NewRouterGroup(path string) gateway.RouterGroup {
	g := router.NewGroup(
		path,
		s.controller,
	)
	s.groups = append(s.groups, g)
	return g
}
