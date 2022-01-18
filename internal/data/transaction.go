package data

import (
	"context"
	"fmt"
	"payment/internal/config"
	"payment/internal/core"

	"github.com/go-logr/logr"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (r transactionRepo) Create(ctx context.Context, t *core.Transaction) error {
	id, err := r.c.InsertOne(ctx, t)
	if err != nil {
		return fmt.Errorf("creating transaction: %w", err)
	}
	t.ID = id.InsertedID.(primitive.ObjectID)
	return nil
}
