package main

import (
	"fmt"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
	"net"
	"os"
	"payment/internal/server"
)

func main() {
	zl := zerolog.New(os.Stderr)
	zl = zl.With().Caller().Timestamp().Logger()
	log := zerologr.New(&zl)
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
	log.Info("Listening on", "address", listenOn)
	if err := s.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
