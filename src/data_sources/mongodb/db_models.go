package mongodb

import "time"

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
