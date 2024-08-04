package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gretchelg/Go_BudgetApp/src/service"
)

const (
	timeout = 10 * time.Second
)

type DBClient struct {
	dbConn                 *mongo.Client
	transactionsCollection *mongo.Collection
	usersCollection        *mongo.Collection
}

// NewDBClient returns a DB client that satisfies the StorageProvider interface defined at the service layer
func NewDBClient(uri string) (service.StorageProvider, error) {
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
	return &DBClient{
		dbConn:                 dbConnection,
		transactionsCollection: txnCollection,
		usersCollection:        usersCollection,
	}, nil
}

// Close cleans up after the connection. It should be invoked in a defer clause after DBClient is instantiated
func (d *DBClient) Close() error {
	// create context used to enforce timeouts
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return d.dbConn.Disconnect(ctx)
}
