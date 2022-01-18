package main

import (
	"fmt"
	"net"
	"payment/internal/config"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"
)

type app struct {
	g    *grpc.Server
	addr string
	log  logr.Logger
}

func newApp(s *grpc.Server, c config.Server, log logr.Logger) *app {
	return &app{g: s, addr: c.GRPC.Addr, log: log}
}

func (a app) Serve() error {
	listener, err := net.Listen("tcp", a.addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listener.Addr(), err)
	}
	a.log.Info("serving gRPC server", "address", a.addr)
	return a.g.Serve(listener)
}
