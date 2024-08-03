package mongodb

import (
	"context"
	"github.com/gretchelg/Go_BudgetApp/src/service"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	timeout = 10 * time.Second
)

type Client struct {
	dbConn                 *mongo.Client
	transactionsCollection *mongo.Collection
	usersCollection        *mongo.Collection
}

// NewClient returns a DB client that satisfies the Storage interface defined at the service layer
// func NewClient(uri string) (*Client, error) {
func NewClient(uri string) (service.Storage, error) {
	// create context used to enforce timeouts
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// connect to db
	dbConnection, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// define collections
	txnCollection := dbConnection.Database("test").Collection("transactions")
	usersCollection := dbConnection.Database("test").Collection("users")

	// respond
	return &Client{
		dbConn:                 dbConnection,
		transactionsCollection: txnCollection,
		usersCollection:        usersCollection,
	}, nil
}

// Close cleans up after the connection. It should be invoked in a defer clause after Client is instantiated
func (c *Client) Close() error {
	// create context used to enforce timeouts
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return c.dbConn.Disconnect(ctx)
}
