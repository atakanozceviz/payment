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
	m := transactionDBModel{
		ID:                   primitive.NewObjectID(),
		PaymentTransactionID: t.PaymentTransactionID(),
		Amount:               t.Amount(),
		Address: addressDBModel{
			City:         t.Address().City(),
			Street:       t.Address().Street(),
			StreetNumber: t.Address().StreetNumber(),
			PostCode:     t.Address().PostCode(),
		},
		PaymentMethodType: t.PaymentMethodType().String(),
		Action:            t.Action().String(),
		Status:            t.Status().String(),
		Metadata:          t.Metadata(),
	}
	id, err := r.c.InsertOne(ctx, m)
	if err != nil {
		return nil, fmt.Errorf("creating transaction: %w", err)
	}
	t.SetId(core.ID(id.InsertedID.(primitive.ObjectID).Hex()))
	return t, nil
}
func (r transactionRepo) Get(ctx context.Context, id core.ID) (*core.Transaction, error) {
	t := new(core.Transaction)
	objectID, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return nil, fmt.Errorf("getting transaction by id: %w", err)
	}
	if err := r.c.FindOne(ctx, bson.M{"_id": objectID}).Decode(t); err != nil {
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
