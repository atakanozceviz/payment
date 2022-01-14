package main

import (
	"fmt"
	"net"
	"payment/internal/server"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

func main() {
	zapLog, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("who watches the watchmen (%v)?", err))
	}
	log := zapr.NewLogger(zapLog)
	if err := run(log); err != nil {
		panic(err)
	}
}

func run(log logr.Logger) error {
	listenOn := "127.0.0.1:4242"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	s := server.NewGRPCServer(log)
	log.Info("serving gRPC server", "address", listenOn)
	if err := s.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
