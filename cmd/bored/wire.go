// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jkuri/bore/pkg/logger"
	"github.com/jkuri/bore/server"
)

var providerSet = wire.NewSet(
	logger.ProviderSet,
	server.ProviderSet,
)

func CreateApp(cfg string) (*server.BoreServer, error) {
	panic(wire.Build(providerSet))
}
