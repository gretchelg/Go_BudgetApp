package service

import (
	"github.com/gretchelg/Go_BudgetApp/src/data_sources/mongodb"
	"github.com/gretchelg/Go_BudgetApp/src/workflows"
)

type Service struct {
	Transactions *workflows.TransactionsWorkflow
	Users        *workflows.UsersWorkflow
}

func NewService(config Config) (*Service, error) {
	// setup dependencies
	db, err := mongodb.NewClient(config.MongoURI)
	if err != nil {
		return nil, err
	}

	// setup workflows
	txnsWorkflow := workflows.NewTransactionsWorkflow(db)
	usersWorkflow := workflows.NewUsersWorkflow(db)

	// respond with ready service
	return &Service{
		Transactions: txnsWorkflow,
		Users:        usersWorkflow,
	}, nil
}
