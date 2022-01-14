package server

import (
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	paymentv1 "payment/api/payment/v1"
	"payment/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(log logr.Logger) *grpc.Server {
	s := grpc.NewServer(grpc.UnaryInterceptor(UnaryServerInterceptor()))
	reflection.Register(s)
	paymentv1.RegisterPaymentServiceServer(s, service.NewPaymentServiceServer(log))
	return s
}
