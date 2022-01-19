package server

import (
	paymentv1 "payment/api/payment/v1"
	"payment/internal/config"
	"payment/internal/server/validator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGRPCServer(c config.Server, payment paymentv1.PaymentServiceServer) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(validator.UnaryServerInterceptor()),
		grpc.ConnectionTimeout(c.GRPC.Timeout),
	)
	reflection.Register(s)
	paymentv1.RegisterPaymentServiceServer(s, payment)
	return s
}
