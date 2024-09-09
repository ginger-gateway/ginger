package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log/logger"
)

type Server interface {
	gateway.Server
}

type srv struct {
	*http.Server

	logger logger.Logger
	config config

	engine     *gin.Engine
	controller gateway.Controller
}

func NewServer(logger logger.Logger, registry registry.Registry) Server {
	s := &srv{
		logger: logger,
	}

	if err := registry.Unmarshal(&s.config); err != nil {
		panic(err)
	}
	s.config.initialize()

	engine := gin.New()

	if len(s.config.RemoteIPHeaders) > 0 {
		engine.RemoteIPHeaders = s.config.RemoteIPHeaders
	}

	engine.Use(gin.Recovery())

	engine.Use(s.newLoggerHandler(logger))
	engine.Use(s.options)
	s.engine = engine

	return s
}

func (s *srv) SetController(controller gateway.Controller) {
	s.controller = controller
}

func (s *srv) GetController() gateway.Controller {
	return s.controller
}

func (s *srv) Run() errors.Error {
	s.Server = &http.Server{
		Addr:    s.config.ListenAddr,
		Handler: s.engine,
	}
	s.logger.Infof("Starting server at %s", s.config.ListenAddr)
	err := s.Server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			return nil
		}
		return errors.New(err)
	}
	return nil
}

func (s *srv) Shutdown(timeout time.Duration) errors.Error {
	if s.Server == nil {
		return errors.Internal().WithMessage("Server is not started yet.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := s.Server.Shutdown(ctx); err != nil {
		return errors.New(err)
	}
	return nil
}
