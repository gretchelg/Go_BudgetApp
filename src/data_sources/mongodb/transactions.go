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
func (d *DBClient) GetAllTransactions(ctx context.Context, filter *models.TransactionsFilter) ([]models.Transaction, error) {

	// create context used to enforce timeouts
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// set filter to narrow results
	dbFilter := bson.D{}

	if filter != nil {
		// for each  non-zero-value field, add it to the list of filter fields.
		filterFields := []bson.E{}

		if filter.User != "" {
			filterFields = append(filterFields, bson.E{Key: "user", Value: filter.User})

			// TODO: debug issue that it doesn't return consistently. Consider changing db field type from "ObjectID" to plain "string"
		}

		dbFilter = filterFields
	}

	// get all transactions
	cursor, err := d.transactionsCollection.Find(ctxWithTimeout, dbFilter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctxWithTimeout)

	// parse the db call response.
	// for each row, convert the data from dbModel to appModel and append to the list of results.
	var results []models.Transaction
	for cursor.Next(ctxWithTimeout) {

		// parse the db row date into the dbModel
		var aDbTxn dbTransaction
		err = cursor.Decode(&aDbTxn)
		if err != nil {
			return nil, err
		}

		// convert dbModel to our appModel
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
func (d *DBClient) GetTransactionByID(ctx context.Context, tranID string) (*models.Transaction, error) {
	// create context used to enforce timeouts
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// create filter that matches the given ID
	filter := bson.D{
		{
			Key:   "tran_id",
			Value: tranID,
		},
	}

	// do find
	var aTransaction dbTransaction
	err := d.transactionsCollection.FindOne(ctxWithTimeout, filter).Decode(&aTransaction)
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
func (d *DBClient) InsertTransaction(ctx context.Context, txn models.Transaction) (string, error) {
	// create context used to enforce timeouts
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// convert from application model to db model
	dbModelTransaction := convertTransactionToDBModel(txn)

	// do insert
	res, err := d.transactionsCollection.InsertOne(ctxWithTimeout, dbModelTransaction)
	if err != nil {
		return "", err
	}

	// respond
	id := fmt.Sprintf("%s", res.InsertedID) // get the string representation of the InsertedID object
	return id, nil
}

// UpdateTransaction updates the transaction specified by the given ID using the provided transaction details.
// Only non-empty fields are updated, empty fields remain unchanged
func (d *DBClient) UpdateTransaction(ctx context.Context, tranID string, txnDelta models.Transaction) error {
	// create context used to enforce timeouts
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// create filter that matches the given ID
	filter := bson.D{
		{
			Key:   "tran_id",
			Value: tranID,
		},
	}

	// convert from application model to db model
	dbModelTransaction := convertTransactionToDBModel(txnDelta)

	// for each  non-zero-value field, add it to the list of fields to update.
	// this is to keep other (zero-value fields) unchanged.
	fieldsToUpdate := []bson.E{}

	if dbModelTransaction.CategoryName != "" {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "category_name", Value: dbModelTransaction.CategoryName})
	}

	if txnDelta.TranAmount != 0 { // keep this check as the original txnDelta, so we can evaluate as a float
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "tran_amount", Value: dbModelTransaction.TranAmount})
	}

	if dbModelTransaction.TranCurrency != "" {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "tran_currency", Value: dbModelTransaction.TranCurrency})
	}

	if !dbModelTransaction.TranDate.IsZero() {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "tran_date", Value: dbModelTransaction.TranDate})
	}

	if dbModelTransaction.TranDescription != "" {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "tran_description", Value: dbModelTransaction.TranDescription})
	}

	if dbModelTransaction.TranSign != "" {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "tran_sign", Value: dbModelTransaction.TranSign})
	}

	if dbModelTransaction.User != "" {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "user", Value: dbModelTransaction.User})
	}

	// Create the updates object
	updates := bson.D{
		{
			Key:   "$set",
			Value: fieldsToUpdate,
		},
	}

	// do update/patch
	_, err := d.transactionsCollection.UpdateOne(ctxWithTimeout, filter, updates)
	if err != nil {
		return err
	}

	// respond
	return nil
}

// convertTransactionToAppModel converts from the internal db model to the application-wide data model.
func convertTransactionToAppModel(dbModel dbTransaction) models.Transaction {
	// convert the string TranAmount field into a proper Float64 value
	floatTranAmount, err := strconv.ParseFloat(dbModel.TranAmount, 64)
	if err != nil {
		floatTranAmount = 0
	}

	// ensure TranSign is valid
	tranSign := models.TranSign(dbModel.TranSign)

	if err := tranSign.Validate(); err != nil {

		// if tranSign is invalid, default to Debit
		tranSign = models.TranSignDebit
	}

	// response
	return models.Transaction{
		TranID:          dbModel.TranID,
		CategoryName:    dbModel.CategoryName,
		TranAmount:      floatTranAmount,
		TranCurrency:    dbModel.TranCurrency,
		TranDate:        dbModel.TranDate,
		TranDescription: dbModel.TranDescription,
		TranSign:        tranSign,
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
		TranSign:        string(appModel.TranSign),
		User:            appModel.User,
	}
}
