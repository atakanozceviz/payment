// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"payment/internal/config"
	"payment/internal/data"
	"payment/internal/logger"
	"payment/internal/server"
	"payment/internal/service"
)

// Injectors from wire.go:

func initApp(configLogger config.Logger, configServer config.Server, configData config.Data) (*server.Server, func(), error) {
	logrLogger, err := logger.New(configLogger)
	if err != nil {
		return nil, nil, err
	}
	database, cleanup, err := data.NewDatabase(configData, logrLogger)
	if err != nil {
		return nil, nil, err
	}
	transactionRepo := data.NewTransactionRepo(configData, database, logrLogger)
	paymentServiceServer := service.NewPaymentServiceServer(transactionRepo, logrLogger)
	grpcServer := server.NewGRPCServer(configServer, paymentServiceServer)
	httpServer := server.NewHTTPServer(configServer)
	serverServer := server.New(grpcServer, httpServer, configServer, logrLogger)
	return serverServer, func() {
		cleanup()
	}, nil
}
