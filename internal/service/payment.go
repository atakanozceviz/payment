package service

import (
	"context"
	paymentv1 "payment/api/payment/v1"
	"payment/internal/core"

	"github.com/go-logr/logr"
	"google.golang.org/genproto/googleapis/type/money"
)

// PaymentServiceServer implements the PaymentService API.
type PaymentServiceServer struct {
	paymentv1.UnimplementedPaymentServiceServer
	r   core.TransactionRepo
	log logr.Logger
}

func NewPaymentServiceServer(r core.TransactionRepo, log logr.Logger) *PaymentServiceServer {
	return &PaymentServiceServer{
		r:   r,
		log: log,
	}
}

func (p PaymentServiceServer) ClientToken(ctx context.Context, request *paymentv1.ClientTokenRequest) (*paymentv1.ClientTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PaymentServiceServer) CalculateTransactionFees(ctx context.Context, request *paymentv1.CalculateTransactionFeesRequest) (*paymentv1.CalculateTransactionFeesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PaymentServiceServer) ProcessPayment(ctx context.Context, request *paymentv1.ProcessPaymentRequest) (*paymentv1.ProcessPaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PaymentServiceServer) RefundPayment(ctx context.Context, request *paymentv1.RefundPaymentRequest) (*paymentv1.RefundPaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PaymentServiceServer) CapturePayment(ctx context.Context, request *paymentv1.CapturePaymentRequest) (*paymentv1.CapturePaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PaymentServiceServer) GetTransactions(ctx context.Context, request *paymentv1.GetTransactionsRequest) (*paymentv1.GetTransactionsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func moneyToCents(m *money.Money) uint64 {
	if m == nil {
		return 0
	}

	return uint64(m.Units*100 + int64(m.Nanos/1_000_000_0))
}
