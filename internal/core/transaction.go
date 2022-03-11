package core

import (
	"context"
	"payment/internal/core/payment_method_type"
	"payment/internal/core/transaction_action"
	"payment/internal/core/transaction_status"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID                   primitive.ObjectID                    `bson:"_id,omitempty"`
	PaymentTransactionID string                                `bson:"payment_transaction_id,omitempty"`
	Amount               uint64                                `bson:"amount,omitempty"`
	Address              Address                               `bson:"address,omitempty"`
	PaymentMethodType    payment_method_type.PaymentMethodType `bson:"payment_method_type,omitempty"`
	Action               transaction_action.TransactionAction  `bson:"action,omitempty"`
	Status               transaction_status.TransactionStatus  `bson:"status,omitempty"`
	Metadata             map[string]string                     `bson:"metadata,omitempty"`
}
type Address struct {
	City         string `bson:"city,omitempty"`
	Street       string `bson:"street,omitempty"`
	StreetNumber string `bson:"street_number,omitempty"`
	PostCode     string `bson:"post_code,omitempty"`
}

type TransactionRepo interface {
	Create(context.Context, *Transaction) (*Transaction, error)
	Get(context.Context, primitive.ObjectID) (*Transaction, error)
	GetByPaymentTransactionID(context.Context, string) (*Transaction, error)
}
