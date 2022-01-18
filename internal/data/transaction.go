package data

import (
	"context"
	"fmt"
	"payment/internal/config"

	"github.com/go-logr/logr"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Transaction struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	PaymentTransactionID string             `bson:"payment_transaction_id,omitempty"`
	UserId               uint32             `bson:"user_id,omitempty"`
	UserKey              string             `bson:"user_key,omitempty"`
	Amount               uint64             `bson:"amount,omitempty"`
	Address              Address            `bson:"address,omitempty"`
	PaymentMethodType    PaymentMethodType  `bson:"payment_method_type,omitempty"`
	TransactionAction    TransactionAction  `bson:"transaction_action,omitempty"`
	TransactionStatus    TransactionStatus  `bson:"transaction_status,omitempty"`
}
type Address struct {
	City         string `bson:"city,omitempty"`
	Street       string `bson:"street,omitempty"`
	StreetNumber string `bson:"street_number,omitempty"`
	PostCode     string `bson:"post_code,omitempty"`
}
type TransactionRepo interface {
	Create(context.Context, *Transaction) error
}

type transactionRepo struct {
	log logr.Logger
	c   *mongo.Collection
}

func NewTransactionRepo(c config.Data, d *Data, log logr.Logger) *transactionRepo {
	collection := d.db.Collection(c.MongoDB.Collection)
	return &transactionRepo{
		log: log,
		c:   collection,
	}
}

func (r transactionRepo) Create(ctx context.Context, t *Transaction) error {
	id, err := r.c.InsertOne(ctx, t)
	if err != nil {
		return fmt.Errorf("creating transaction: %w", err)
	}
	t.ID = id.InsertedID.(primitive.ObjectID)
	return nil
}
