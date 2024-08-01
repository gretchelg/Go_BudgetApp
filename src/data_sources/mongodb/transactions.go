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
		aTransaction := convertTransaction(aDbTxn)

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
		// Do something when no record was found
		//fmt.Println("record does not exist")
		return nil, fmt.Errorf("DB.GetTransactionByID: %w", models.ErrorNotFound)
	}

	if err != nil {
		return nil, err
	}

	// respond
	result := convertTransaction(aTransaction)
	return &result, nil

}

// convertTransaction converts from the internal db model to the application-wide data model.
func convertTransaction(dbModel dbTransaction) models.Transaction {
	// convert the string TxnAmt field into a proper Float64 value
	floatTxnAmt, err := strconv.ParseFloat(dbModel.TranAmount, 64)
	if err != nil {
		floatTxnAmt = 0
	}

	// response
	return models.Transaction{
		TranID:          dbModel.TranID,
		CategoryName:    dbModel.CategoryName,
		TranAmount:      floatTxnAmt,
		TranCurrency:    dbModel.TranCurrency,
		TranDate:        dbModel.TranDate,
		TranDescription: dbModel.TranDescription,
		TranSign:        dbModel.TranSign,
		User:            dbModel.User,
	}
}
