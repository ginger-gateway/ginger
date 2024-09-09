package ginger

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log/logger"
	"github.com/ginger-gateway/ginger/grpc"
)

func NewGRPC(logger logger.Logger,
	registry registry.Registry) grpc.Server {
	return grpc.NewServer(logger, registry)
}
