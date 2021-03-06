package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type transactionDBModel struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	PaymentTransactionID string             `bson:"payment_transaction_id,omitempty"`
	Amount               uint64             `bson:"amount,omitempty"`
	Address              addressDBModel     `bson:"address,omitempty"`
	PaymentMethodType    string             `bson:"payment_method_type,omitempty"`
	Action               string             `bson:"action,omitempty"`
	Status               string             `bson:"status,omitempty"`
	Metadata             map[string]string  `bson:"metadata,omitempty"`
}
type addressDBModel struct {
	City         string `bson:"city,omitempty"`
	Street       string `bson:"street,omitempty"`
	StreetNumber string `bson:"street_number,omitempty"`
	PostCode     string `bson:"post_code,omitempty"`
}
