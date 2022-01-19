package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"payment/internal/config"

	"github.com/go-logr/logr"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

var ProviderSet = wire.NewSet(
	NewGRPCServer,
	NewHTTPServer,
)

type Server struct {
	g        *grpc.Server
	h        *http.Server
	grpcAddr string
	httpAddr string
	log      logr.Logger
}

func New(g *grpc.Server, h *http.Server, c config.Server, log logr.Logger) *Server {
	return &Server{g: g, h: h, grpcAddr: c.GRPC.Addr, httpAddr: c.HTTP.Addr, log: log}
}

func (a Server) ServeHTTP() error {
	if a.httpAddr == "" {
		a.log.Info("HTTP server not serving!: empty address")
		return nil
	}
	a.log.Info("serving HTTP server", "address", a.httpAddr)
	return a.h.ListenAndServe()
}
func (a Server) ServeGRPC() error {
	if a.grpcAddr == "" {
		a.log.Info("gRPC server not serving!: empty address")
		return nil
	}
	listener, err := net.Listen("tcp", a.grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", a.grpcAddr, err)
	}
	a.log.Info("serving gRPC server", "address", a.grpcAddr)
	return a.g.Serve(listener)
}

func (a Server) Shutdown() {
	a.log.Info("server shutdown")
	if a.grpcAddr != "" {
		a.g.GracefulStop()
	}
	if a.httpAddr != "" {
		_ = a.h.Shutdown(context.Background())
	}
}
