package main

import (
	"flag"
	"fmt"
	"net"
	"payment/internal/config"
	"payment/internal/data"
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

func run(c *config.Config, log logr.Logger) error {
	listener, err := net.Listen("tcp", c.Server.GRPC.Addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", c.Server.GRPC.Addr, err)
	}

	d, cleanup, err := data.NewData(c.Data, log)
	if err != nil {
		return fmt.Errorf("crating data: %w", err)
	}
	defer cleanup()
	repo := data.NewTransactionRepo(c.Data.MongoDB, d, log)
	s := server.NewGRPCServer(repo, log)
	log.Info("serving gRPC server", "address", c.Server.GRPC.Addr)
	if err := s.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
