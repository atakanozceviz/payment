package service

import (
	paymentv1 "payment/api/payment/v1"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewPaymentServiceServer,
	wire.Bind(new(paymentv1.PaymentServiceServer), new(*PaymentServiceServer)),
)
