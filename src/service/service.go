package service

import (
	"github.com/gretchelg/Go_BudgetApp/src/data_sources/mongodb"
	"github.com/gretchelg/Go_BudgetApp/src/workflows"
)

// Service defines the core service of our app, and provides access to the underlying Workflow functionalities
// It has no HTTP functionality.
// See "Server" pkg, which wraps this Service in order to expose its functionalities over HTTP.
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
	transactionsWorkflow := workflows.NewTransactionsWorkflow(db)
	usersWorkflow := workflows.NewUsersWorkflow(db)

	// respond with ready service
	return &Service{
		Transactions: transactionsWorkflow,
		Users:        usersWorkflow,
	}, nil
}
