package ginger

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log/logger"
	"github.com/ginger-gateway/ginger/internal/router"
)

type serverConfig struct {
	Logger struct {
		SkipPaths []string
	}
	ListenAddr string
	Router     router.Config
}

func (c *serverConfig) initialize() {
}

type server struct {
	logger logger.Logger
	config *serverConfig

	*http.Server
	engine     *gin.Engine
	controller gateway.Controller
}

func NewServer(logger logger.Logger, registry registry.Registry) gateway.Server {
	s := &server{
		logger: logger,
		config: new(serverConfig),
	}

	if err := registry.Unmarshal(&s.config); err != nil {
		panic(err)
	}
	s.config.initialize()

	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.Use(s.newLoggerHandler(logger))
	engine.Use(s.options)
	s.engine = engine

	return s
}

func (s *server) SetController(controller gateway.Controller) {
	s.controller = controller
}

func (s *server) GetController() gateway.Controller {
	return s.controller
}

func (s *server) Run() errors.Error {
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

func (s *server) Shutdown(timeout time.Duration) errors.Error {
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
