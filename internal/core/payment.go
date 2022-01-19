package core

import "github.com/go-logr/logr"

type Payment struct {
	log logr.Logger
	r   TransactionRepo
}
