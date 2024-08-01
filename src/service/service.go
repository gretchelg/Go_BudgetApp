package service

import (
	"github.com/gretchelg/Go_BudgetApp/src/data_sources/mongodb"
	"github.com/gretchelg/Go_BudgetApp/src/models"
	"github.com/gretchelg/Go_BudgetApp/src/workflows"
)

// TransactionsRepository defines method required of a Transactions Repository / DB
type TransactionsRepository interface {
	GetAllTransactions() ([]models.Transaction, error)
}

type Service struct {
	Transactions *workflows.TransactionsWorkflow
}

func NewService(config Config) (*Service, error) {
	// setup dependencies
	db, err := mongodb.NewClient(config.MongoURI)

	if err != nil {
		return nil, err
	}

	txnsWorkflow := workflows.NewTransactionsWorkflow(db)

	// respond with ready service
	return &Service{
		Transactions: txnsWorkflow,
	}, nil
}
