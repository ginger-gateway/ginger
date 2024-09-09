package ginger

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log/logger"
	"github.com/ginger-gateway/ginger/http"
)

func NewHTTP(logger logger.Logger,
	registry registry.Registry) gateway.Server {
	return http.NewServer(logger, registry)
}
