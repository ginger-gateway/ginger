package grpc

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger/grpc/internal/router"
	"google.golang.org/grpc"
)

func (s *srv) Register(desc *grpc.ServiceDesc) gateway.RouterGroup {
	g := router.NewGroup(desc.ServiceName, s.getController())
	g.Initialize(desc)
	s.groups = append(s.groups, g)
	return g
}
