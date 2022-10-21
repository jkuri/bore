// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jkuri/bore/pkg/logger"
	"github.com/jkuri/bore/server"
)

// Injectors from wire.go:

func CreateApp(cfg string) (*server.BoreServer, error) {
	viper, err := server.NewConfig(cfg)
	if err != nil {
		return nil, err
	}
	options, err := server.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	loggerOptions, err := logger.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	zapLogger, err := logger.NewLogger(loggerOptions)
	if err != nil {
		return nil, err
	}
	boreServer := server.NewBoreServer(options, zapLogger)
	return boreServer, nil
}

// wire.go:

var providerSet = wire.NewSet(logger.ProviderSet, server.ProviderSet)
