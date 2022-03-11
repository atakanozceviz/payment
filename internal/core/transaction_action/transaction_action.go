package transaction_action

import "errors"

type TransactionAction struct {
	slug string
}

var ErrInvalidTransactionAction = errors.New("invalid transaction action")

var (
	Unknown        = TransactionAction{""}
	CapturePayment = TransactionAction{"CapturePayment"}
	ProcessPayment = TransactionAction{"ProcessPayment"}
	RefundPayment  = TransactionAction{"RefundPayment"}
	UpdateStatus   = TransactionAction{"UpdateStatus"}
)

func FromString(s string) (TransactionAction, error) {
	switch s {
	case CapturePayment.slug:
		return CapturePayment, nil
	case ProcessPayment.slug:
		return ProcessPayment, nil
	case RefundPayment.slug:
		return RefundPayment, nil
	case UpdateStatus.slug:
		return UpdateStatus, nil
	}

	return Unknown, ErrInvalidTransactionAction
}
