package service

import (
	"context"
	paymentv1 "payment/api/payment/v1"
	"payment/internal/data"

	"github.com/go-logr/logr"
	"google.golang.org/genproto/googleapis/type/money"
)

// paymentServiceServer implements the PaymentService API.
type paymentServiceServer struct {
	paymentv1.UnimplementedPaymentServiceServer
	r   data.TransactionRepo
	log logr.Logger
}

func NewPaymentServiceServer(r data.TransactionRepo, log logr.Logger) paymentv1.PaymentServiceServer {
	return &paymentServiceServer{
		r:   r,
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
	paymentMethodType := data.PaymentMethodType(request.PaymentMethod)
	if err := paymentMethodType.IsValid(); err != nil {
		return nil, err
	}

	t := &data.Transaction{
		UserId:  request.UserId,
		UserKey: request.UserKey,
		Amount:  moneyToCents(request.Amount),
		Address: data.Address{
			City:         request.Address.City,
			Street:       request.Address.Street,
			StreetNumber: request.Address.StreetNumber,
			PostCode:     request.Address.PostCode,
		},
		PaymentMethodType: paymentMethodType,
		TransactionAction: data.ProcessPayment,
		TransactionStatus: data.Processing,
	}
	err := p.r.Create(ctx, t)
	if err != nil {
		return nil, err
	}
	return &paymentv1.ProcessPaymentResponse{
		Id:                   t.ID.Hex(),
		PaymentTransactionId: t.PaymentTransactionID,
		ClientSecret:         "hihi",
		TransactionStatus:    string(t.TransactionStatus),
	}, nil
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

func moneyToCents(m *money.Money) uint64 {
	if m == nil {
		return 0
	}

	return uint64(m.Units*100 + int64(m.Nanos/1_000_000_0))
}
