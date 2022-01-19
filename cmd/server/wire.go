//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"payment/internal/config"
	"payment/internal/data"
	"payment/internal/logger"
	"payment/internal/server"
	"payment/internal/service"

	"github.com/google/wire"
)

func initApp(config.Logger, config.Server, config.Data) (*server.Server, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, logger.ProviderSet, server.New))
}
