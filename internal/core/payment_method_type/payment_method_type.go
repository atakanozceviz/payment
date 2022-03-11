package payment_method_type

import "errors"

type PaymentMethodType struct {
	slug string
}

func (p PaymentMethodType) String() string {
	return p.slug
}

var ErrInvalidPaymentMethodType = errors.New("invalid payment method type")

var (
	Unknown         = PaymentMethodType{""}
	GooglePay       = PaymentMethodType{"GooglePay"}
	ApplePay        = PaymentMethodType{"ApplePay"}
	Mastercard      = PaymentMethodType{"Mastercard"}
	Visa            = PaymentMethodType{"Visa"}
	AmericanExpress = PaymentMethodType{"AmericanExpress"}
	PayPal          = PaymentMethodType{"PayPal"}
	BankTransfer    = PaymentMethodType{"BankTransfer"}
	Cash            = PaymentMethodType{"Cash"}
)

func FromString(s string) (PaymentMethodType, error) {
	switch s {
	case GooglePay.slug:
		return GooglePay, nil
	case ApplePay.slug:
		return ApplePay, nil
	case Mastercard.slug:
		return Mastercard, nil
	case Visa.slug:
		return Visa, nil
	case AmericanExpress.slug:
		return AmericanExpress, nil
	case PayPal.slug:
		return PayPal, nil
	case BankTransfer.slug:
		return BankTransfer, nil
	case Cash.slug:
		return Cash, nil
	}

	return Unknown, ErrInvalidPaymentMethodType
}
