package models

import (
	"time"
)

// Transaction defines a transaction
// the `bson` tag is useful for mapping to bson-encoded DB systems
type Transaction struct {
	TranID          string    `json:"tran_id" bson:"tran_id"`
	CategoryName    string    `json:"category_name" bson:"category_name"`
	TranAmount      float64   `json:"tran_amount" bson:"tran_amount"`
	TranCurrency    string    `json:"tran_currency" bson:"tran_currency"`
	TranDate        time.Time `json:"tran_date" bson:"tran_date"`
	TranDescription string    `json:"tran_description" bson:"tran_description"`
	TranSign        TranSign  `json:"tran_sign" bson:"tran_sign"`
	User            string    `json:"user" bson:"user"`
}
