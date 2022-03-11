package core

import (
	"context"
	"payment/internal/core/payment_method_type"
	"payment/internal/core/transaction_action"
	"payment/internal/core/transaction_status"
)

type ID string

type Transaction struct {
	id                   ID
	paymentTransactionID string
	amount               uint64
	address              *Address
	paymentMethodType    payment_method_type.PaymentMethodType
	action               transaction_action.TransactionAction
	status               transaction_status.TransactionStatus
	metadata             map[string]string
}

func (t Transaction) ID() ID {
	return t.id
}

func (t *Transaction) SetId(id ID) {
	t.id = id
}

func (t Transaction) PaymentTransactionID() string {
	return t.paymentTransactionID
}

func (t Transaction) Amount() uint64 {
	return t.amount
}

func (t Transaction) Address() *Address {
	return t.address
}

func (t Transaction) PaymentMethodType() payment_method_type.PaymentMethodType {
	return t.paymentMethodType
}

func (t Transaction) Action() transaction_action.TransactionAction {
	return t.action
}

func (t Transaction) Status() transaction_status.TransactionStatus {
	return t.status
}

func (t Transaction) Metadata() map[string]string {
	return t.metadata
}

func NewTransaction(
	id ID,
	paymentTransactionID string,
	amount uint64,
	address *Address,
	paymentMethodType payment_method_type.PaymentMethodType,
	action transaction_action.TransactionAction,
	status transaction_status.TransactionStatus,
	metadata map[string]string) *Transaction {
	return &Transaction{
		id:                   id,
		paymentTransactionID: paymentTransactionID,
		amount:               amount,
		address:              address,
		paymentMethodType:    paymentMethodType,
		action:               action,
		status:               status,
		metadata:             metadata,
	}
}

type Address struct {
	city         string
	street       string
	streetNumber string
	postCode     string
}

func (a Address) City() string {
	return a.city
}

func (a Address) Street() string {
	return a.street
}

func (a Address) StreetNumber() string {
	return a.streetNumber
}

func (a Address) PostCode() string {
	return a.postCode
}

func NewAddress(
	city string,
	street string,
	streetNumber string,
	postCode string) *Address {
	return &Address{
		city:         city,
		street:       street,
		streetNumber: streetNumber,
		postCode:     postCode,
	}
}

type TransactionRepo interface {
	Create(context.Context, *Transaction) (*Transaction, error)
	Get(context.Context, ID) (*Transaction, error)
	GetByPaymentTransactionID(context.Context, string) (*Transaction, error)
}
