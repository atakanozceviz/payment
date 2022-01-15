package main

import (
	"flag"
	"fmt"
	"net"
	"payment/internal/config"
	"payment/internal/logger"
	"payment/internal/server"

	"github.com/go-logr/logr"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "conf", "configs/config.toml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	c, err := config.Configure(confPath)
	if err != nil {
		panic(fmt.Sprintf("error configuring service: %v", err))
	}
	log, err := logger.New(c.Logger.Env)
	if err != nil {
		panic(fmt.Sprintf("error creating logger: %v", err))
	}
	if err := run(c, log); err != nil {
		panic(err)
	}
}

func run(config *config.Config, log logr.Logger) error {
	listener, err := net.Listen("tcp", config.Server.GRPC.Addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", config.Server.GRPC.Addr, err)
	}

	s := server.NewGRPCServer(log)
	log.Info("serving gRPC server", "address", config.Server.GRPC.Addr)
	if err := s.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
