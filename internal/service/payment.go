package service

import (
	"context"
	paymentv1 "payment/api/payment/v1"

	"github.com/go-logr/logr"
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

func (p paymentServiceServer) ClientToken(ctx context.Context, request *paymentv1.ClientTokenRequest) (*paymentv1.ClientTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p paymentServiceServer) CalculateTransactionFees(ctx context.Context, request *paymentv1.CalculateTransactionFeesRequest) (*paymentv1.CalculateTransactionFeesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p paymentServiceServer) ProcessPayment(ctx context.Context, request *paymentv1.ProcessPaymentRequest) (*paymentv1.ProcessPaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p paymentServiceServer) RefundPayment(ctx context.Context, request *paymentv1.RefundPaymentRequest) (*paymentv1.RefundPaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p paymentServiceServer) CapturePayment(ctx context.Context, request *paymentv1.CapturePaymentRequest) (*paymentv1.CapturePaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p paymentServiceServer) GetTransactions(ctx context.Context, request *paymentv1.GetTransactionsRequest) (*paymentv1.GetTransactionsResponse, error) {
	//TODO implement me
	panic("implement me")
}
