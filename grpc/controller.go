package grpc

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger/grpc/internal/controller"
)

func (s *srv) SetController(c gateway.Controller) {
	s.setController(c.(controller.Controller))
}

func (s *srv) GetController() gateway.Controller {
	return s.controller
}

func (s *srv) setController(controller controller.Controller) {
	s.controller = controller
}

func (s *srv) getController() controller.Controller {
	return s.controller
}
