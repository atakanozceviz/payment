package server

import (
	paymentv1 "payment/api/payment/v1"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGRPCServer(payment paymentv1.PaymentServiceServer) *grpc.Server {
	s := grpc.NewServer(grpc.UnaryInterceptor(UnaryServerInterceptor()), grpc.ConnectionTimeout(time.Second))
	reflection.Register(s)
	paymentv1.RegisterPaymentServiceServer(s, payment)
	return s
}
