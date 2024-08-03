package mongodb

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gretchelg/Go_BudgetApp/src/models"
)

// dbTransaction defines a transaction as specified in the DB
// the `bson` tag maps to the field names in the db.
type dbTransaction struct {
	TranID          string    `bson:"tran_id"`
	CategoryName    string    `bson:"category_name"`
	TranAmount      string    `bson:"tran_amount"` // NOTE this needs conversion
	TranCurrency    string    `bson:"tran_currency"`
	TranDate        time.Time `bson:"tran_date"`
	TranDescription string    `bson:"tran_description"`
	TranSign        string    `bson:"tran_sign"`
	User            string    `bson:"user"`
}

// GetAllTransactions returns all Transactions
func (c *Client) GetAllTransactions() ([]models.Transaction, error) {
	// create context used to enforce timeouts
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// get all transactions
	cursor, err := c.transactionsCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	// parse the db call response
	var results []models.Transaction
	for cursor.Next(ctx) {

		//var result bson.D
		//var result bson.M
		var aDbTxn dbTransaction
		err = cursor.Decode(&aDbTxn)
		if err != nil {
			return nil, err
		}

		// convert the row from an internal db model to the application model
		aTransaction := convertTransactionToAppModel(aDbTxn)

		// append to the list of results
		results = append(results, aTransaction)
	}

	// final check for any errors reported
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	// respond
	return results, nil
}

// GetTransactionByID returns one transaction specified by the given ID
func (c *Client) GetTransactionByID(id string) (*models.Transaction, error) {
	// create context used to enforce timeouts
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// create filter that matches the given ID
	filter := bson.D{
		{
			Key:   "tran_id",
			Value: id,
		},
	}

	// do find
	var aTransaction dbTransaction
	err := c.transactionsCollection.FindOne(ctx, filter).Decode(&aTransaction)
	if err == mongo.ErrNoDocuments {
		// if no matching docs found, return sentinel error "models.ErrorNotFound" that callers can inspect in order to
		// handle in a custom way, such as returning 404-NotFound rather than a generic 500-InternalServerError
		return nil, fmt.Errorf("DB.GetTransactionByID: %w", models.ErrorNotFound)
	}

	if err != nil {
		return nil, err
	}

	// respond
	result := convertTransactionToAppModel(aTransaction)
	return &result, nil
}

// InsertTransaction inserts the given transaction into the database. It returns the database ID of the inserted row
func (c *Client) InsertTransaction(txn models.Transaction) (string, error) {
	// create context used to enforce timeouts
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// convert from application model to db model
	dbModelTransaction := convertTransactionToDBModel(txn)

	// do insert
	res, err := c.transactionsCollection.InsertOne(ctx, dbModelTransaction)
	if err != nil {
		return "", err
	}

	// respond
	id := fmt.Sprintf("%s", res.InsertedID) // get the string representation of the InsertedID object
	return id, nil
}

// convertTransactionToAppModel converts from the internal db model to the application-wide data model.
func convertTransactionToAppModel(dbModel dbTransaction) models.Transaction {
	// convert the string TranAmount field into a proper Float64 value
	floatTranAmount, err := strconv.ParseFloat(dbModel.TranAmount, 64)
	if err != nil {
		floatTranAmount = 0
	}

	// response
	return models.Transaction{
		TranID:          dbModel.TranID,
		CategoryName:    dbModel.CategoryName,
		TranAmount:      floatTranAmount,
		TranCurrency:    dbModel.TranCurrency,
		TranDate:        dbModel.TranDate,
		TranDescription: dbModel.TranDescription,
		TranSign:        dbModel.TranSign,
		User:            dbModel.User,
	}
}

// convertTransactionToDBModel converts from the application model to the internal db model
func convertTransactionToDBModel(appModel models.Transaction) dbTransaction {
	// convert the float64 TranAmount field into string as specified in the db model
	strTranAmount := fmt.Sprintf("%f", appModel.TranAmount)

	// response
	return dbTransaction{
		TranID:          appModel.TranID,
		CategoryName:    appModel.CategoryName,
		TranAmount:      strTranAmount,
		TranCurrency:    appModel.TranCurrency,
		TranDate:        appModel.TranDate,
		TranDescription: appModel.TranDescription,
		TranSign:        appModel.TranSign,
		User:            appModel.User,
	}
}
