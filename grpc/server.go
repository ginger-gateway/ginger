package grpc

import (
	"net"
	"time"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log/logger"
	"github.com/ginger-gateway/ginger/grpc/internal/controller"
	"github.com/ginger-gateway/ginger/grpc/internal/router"
	"google.golang.org/grpc"
)

type Server interface {
	gateway.Server
	Initialize(responder gateway.Responder)
	Register(desc *grpc.ServiceDesc) gateway.RouterGroup
}

type srv struct {
	*grpc.Server

	logger logger.Logger
	config *config

	controller controller.Controller
	groups     []router.RouterGroup
}

func NewServer(logger logger.Logger,
	registry registry.Registry) Server {
	s := &srv{
		logger: logger,
		config: new(config),
		groups: make([]router.RouterGroup, 0),
	}

	if err := registry.Unmarshal(&s.config); err != nil {
		panic(err)
	}
	s.config.initialize()
	return s
}

func (s *srv) Initialize(responder gateway.Responder) {
	controller := controller.New(responder)
	s.setController(controller)
	var sererOptions []grpc.ServerOption = []grpc.ServerOption{
		grpc.UnaryInterceptor(controller.Handle),
	}
	s.Server = grpc.NewServer(sererOptions...)
}

func (s *srv) Run() errors.Error {
	l, err := net.Listen("tcp", s.config.ListenAddr)
	if err != nil {
		return errors.New(err).
			WithTrace("net.Listen")
	}
	for _, g := range s.groups {
		s.Server.RegisterService(g.GetDesc(), nil)
	}
	s.logger.Infof("grpc server listening to %s", s.config.ListenAddr)
	if err = s.Server.Serve(l); err != nil {
		return errors.New(err).
			WithTrace("gRpcServer.Serve")
	}
	return nil
}

func (s *srv) Shutdown(timeout time.Duration) (err errors.Error) {
	closeCh := make(chan bool)
	stopped := false
	go func() {
		s.Server.GracefulStop()
		stopped = true
		closeCh <- true
	}()
	go func() {
		time.Sleep(timeout)
		if !stopped {
			// err = errors.New().
			// 	WithTrace("gRpcServer.GracefulStop.timeout")
			s.Server.Stop()
			close(closeCh)
		}
	}()
	<-closeCh
	return
}
