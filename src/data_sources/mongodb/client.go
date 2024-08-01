package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gretchelg/Go_BudgetApp/src/workflows"
)

const (
	timeout = 10 * time.Second
)

type Client struct {
	db                     *mongo.Client
	transactionsCollection *mongo.Collection
}

// NewClient returns a DB client that satisfies the TransactionsStorage defined at the service layer
func NewClient(uri string) (workflows.TransactionsStorage, error) {
	// create context used to enforce timeouts
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// connect to db
	mongoDbClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// define collections
	txnCollection := mongoDbClient.Database("test").Collection("transactions")

	// respond
	return &Client{
		db:                     mongoDbClient,
		transactionsCollection: txnCollection,
	}, nil
}

// Close cleans up after the connection. It should be invoked in a defer clause after Client is instantiated
func (c *Client) Close() error {
	// create context used to enforce timeouts
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return c.db.Disconnect(ctx)
}
