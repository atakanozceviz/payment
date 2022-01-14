package service

import (
	"context"
	"github.com/go-logr/logr"
	paymentv1 "payment/api/payment/v1"
)

// paymentServiceServer implements the PaymentService API.
type paymentServiceServer struct {
	paymentv1.UnimplementedPaymentServiceServer
	log logr.Logger
}

func NewPaymentServiceServer(log logr.Logger) paymentv1.PaymentServiceServer {
	return &paymentServiceServer{
		log: log,
	}
}

func (s paymentServiceServer) ClientToken(context.Context, *paymentv1.ClientTokenRequest) (*paymentv1.ClientTokenResponse, error) {
	response := &paymentv1.ClientTokenResponse{ClientToken: "it works!"}
	return response, nil
}
func (s paymentServiceServer) CalculateTransactionFees(context.Context, *paymentv1.CalculateTransactionFeesRequest) (*paymentv1.CalculateTransactionFeesResponse, error) {
	return nil, nil
}
func (s paymentServiceServer) ProcessPayment(context.Context, *paymentv1.ProcessPaymentRequest) (*paymentv1.ProcessPaymentResponse, error) {
	return nil, nil
}
func (s paymentServiceServer) RefundPayment(context.Context, *paymentv1.RefundPaymentRequest) (*paymentv1.RefundPaymentResponse, error) {
	return nil, nil
}
func (s paymentServiceServer) CapturePayment(context.Context, *paymentv1.CapturePaymentRequest) (*paymentv1.CapturePaymentResponse, error) {
	return nil, nil
}
func (s paymentServiceServer) GetTransactions(context.Context, *paymentv1.GetTransactionsRequest) (*paymentv1.GetTransactionsResponse, error) {
	return nil, nil
}
