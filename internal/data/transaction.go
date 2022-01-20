package data

import (
	"context"
	"fmt"
	"payment/internal/config"
	"payment/internal/core"

	"github.com/go-logr/logr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type transactionRepo struct {
	log logr.Logger
	c   *mongo.Collection
}

func NewTransactionRepo(c config.Data, d *Database, log logr.Logger) *transactionRepo {
	collection := d.db.Collection(c.MongoDB.Collection)
	return &transactionRepo{
		log: log,
		c:   collection,
	}
}

func (r transactionRepo) Create(ctx context.Context, t *core.Transaction) (*core.Transaction, error) {
	id, err := r.c.InsertOne(ctx, t)
	if err != nil {
		return nil, fmt.Errorf("creating transaction: %w", err)
	}
	t.ID = id.InsertedID.(primitive.ObjectID)
	return t, nil
}
func (r transactionRepo) Get(ctx context.Context, id primitive.ObjectID) (*core.Transaction, error) {
	t := new(core.Transaction)
	if err := r.c.FindOne(ctx, bson.M{"_id": id}).Decode(t); err != nil {
		return nil, fmt.Errorf("getting transaction by id: %w", err)
	}
	return t, nil
}
func (r transactionRepo) GetByPaymentTransactionID(ctx context.Context, ptID string) (*core.Transaction, error) {
	match := bson.D{{Key: "payment_transaction_id", Value: ptID}}
	opts := options.FindOne().SetSort(bson.D{{Key: "_id", Value: -1}})

	t := new(core.Transaction)
	if err := r.c.FindOne(ctx, match, opts).Decode(t); err != nil {
		return nil, fmt.Errorf("getting transaction by payment transaction id: %w", err)
	}
	return t, nil
}
