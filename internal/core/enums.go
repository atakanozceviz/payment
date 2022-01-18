package core

import "errors"

type PaymentMethodType string

var ErrInvalidPaymentMethodType = errors.New("invalid payment method type")

const (
	GooglePay       PaymentMethodType = "GooglePay"
	ApplePay                          = "ApplePay"
	Mastercard                        = "Mastercard"
	Visa                              = "Visa"
	AmericanExpress                   = "AmericanExpress"
	PayPal                            = "PayPal"
	BankTransfer                      = "BankTransfer"
	Cash                              = "Cash"
)

func (pmt PaymentMethodType) IsValid() error {
	switch pmt {
	case GooglePay, ApplePay, Mastercard, Visa, AmericanExpress, PayPal, BankTransfer, Cash:
		return nil
	}
	return ErrInvalidPaymentMethodType
}

type TransactionAction string

var ErrInvalidTransactionAction = errors.New("invalid transaction action")

const (
	CapturePayment TransactionAction = "CapturePayment"
	ProcessPayment                   = "ProcessPayment"
	RefundPayment                    = "RefundPayment"
	UpdateStatus                     = "UpdateStatus"
)

func (ta TransactionAction) IsValid() error {
	switch ta {
	case CapturePayment, ProcessPayment, RefundPayment, UpdateStatus:
		return nil
	}
	return ErrInvalidTransactionAction
}

type TransactionStatus string

var ErrInvalidTransactionStatus = errors.New("invalid transaction status")

const (
	RequiresPaymentMethod TransactionStatus = "RequiresPaymentMethod"
	RequiresConfirmation                    = "RequiresConfirmation"
	RequiresCapture                         = "RequiresCapture"
	RequiresAction                          = "RequiresAction"
	Processing                              = "Processing"
	Succeeded                               = "Succeeded"
	Canceled                                = "Canceled"
)

func (ta TransactionStatus) IsValid() error {
	switch ta {
	case RequiresPaymentMethod, RequiresConfirmation, RequiresCapture,
		RequiresAction, Processing, Succeeded, Canceled:
		return nil
	}
	return ErrInvalidTransactionStatus
}
