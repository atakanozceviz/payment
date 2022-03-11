package core

import (
	"context"
	"payment/internal/core/payment_method_type"
	"payment/internal/core/transaction_action"
	"payment/internal/core/transaction_status"
)

type ID string

type Transaction struct {
	ID                   ID
	PaymentTransactionID string
	Amount               uint64
	Address              Address
	PaymentMethodType    payment_method_type.PaymentMethodType
	Action               transaction_action.TransactionAction
	Status               transaction_status.TransactionStatus
	Metadata             map[string]string
}
type Address struct {
	City         string
	Street       string
	StreetNumber string
	PostCode     string
}

type TransactionRepo interface {
	Create(context.Context, *Transaction) (*Transaction, error)
	Get(context.Context, ID) (*Transaction, error)
	GetByPaymentTransactionID(context.Context, string) (*Transaction, error)
}
