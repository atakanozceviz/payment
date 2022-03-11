package transaction_status

import "errors"

type TransactionStatus struct {
	slug string
}

var ErrInvalidTransactionStatus = errors.New("invalid transaction status")

var (
	Unknown               = TransactionStatus{""}
	RequiresAuthorization = TransactionStatus{"RequiresAuthorization"}
	Authorized            = TransactionStatus{"Authorized"}
	Captured              = TransactionStatus{"Captured"}
	Refunded              = TransactionStatus{"Refunded"}
	Voided                = TransactionStatus{"Voided"}
	Expired               = TransactionStatus{"Expired"}
	Failed                = TransactionStatus{"Failed"}
)

func FromString(s string) (TransactionStatus, error) {
	switch s {
	case RequiresAuthorization.slug:
		return RequiresAuthorization, nil
	case Authorized.slug:
		return Authorized, nil
	case Captured.slug:
		return Captured, nil
	case Refunded.slug:
		return Refunded, nil
	case Voided.slug:
		return Voided, nil
	case Expired.slug:
		return Expired, nil
	case Failed.slug:
		return Failed, nil
	}

	return Unknown, ErrInvalidTransactionStatus
}
