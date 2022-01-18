package main

import (
	"fmt"
	"net"
	"net/http"
	"payment/internal/config"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"
)

type app struct {
	g        *grpc.Server
	r        http.Handler
	grpcAddr string
	httpAddr string
	log      logr.Logger
}

func newApp(g *grpc.Server, r http.Handler, c config.Server, log logr.Logger) *app {
	return &app{g: g, r: r, grpcAddr: c.GRPC.Addr, httpAddr: c.HTTP.Addr, log: log}
}

func (a app) Serve() error {
	listener, err := net.Listen("tcp", a.grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listener.Addr(), err)
	}

	a.log.Info("serving gRPC server", "address", a.grpcAddr)
	go func(l net.Listener) {
		if err := a.g.Serve(l); err != nil {
			panic(err)
		}
	}(listener)

	a.log.Info("serving HTTP server", "address", a.httpAddr)
	return http.ListenAndServe(a.httpAddr, a.r)
}
