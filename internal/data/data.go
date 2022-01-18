package data

import (
	"context"
	"fmt"
	"payment/internal/config"
	"time"

	"github.com/go-logr/logr"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Data struct {
	db *mongo.Database
}

func NewData(c config.Data, log logr.Logger) (*Data, func(), error) {
	// create connection
	client, err := mongo.NewClient(options.Client().ApplyURI(c.MongoDB.ConnectionString))
	if err != nil {
		return nil, nil, fmt.Errorf("creating mongodb client: %w", err)
	}
	cleanup := func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Error(err, "error while disconnecting from mongodb")
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("connecting to mongodb: %w", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, nil, fmt.Errorf("sending ping command: %w", err)
	}

	return &Data{db: client.Database(c.MongoDB.Database)}, cleanup, nil
}
