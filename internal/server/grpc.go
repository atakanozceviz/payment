package server

import (
	paymentv1 "payment/api/payment/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(payment paymentv1.PaymentServiceServer) *grpc.Server {
	s := grpc.NewServer(grpc.UnaryInterceptor(UnaryServerInterceptor()))
	reflection.Register(s)
	paymentv1.RegisterPaymentServiceServer(s, payment)
	return s
}
