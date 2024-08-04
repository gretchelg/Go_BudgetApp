package service

import "github.com/gretchelg/Go_BudgetApp/src/models"

// Storage defines methods require of the storage layer (database).
type Storage interface {
	// Transactions-related
	GetAllTransactions() ([]models.Transaction, error)
	GetTransactionByID(tranID string) (*models.Transaction, error)
	InsertTransaction(txn models.Transaction) (string, error)
	UpdateTransaction(tranID string, txn models.Transaction) error

	// Users-related
	GetAllUsers() ([]models.User, error)
}

// Service defines the core service of our app, and provides access to the underlying Workflow functionalities
// It has no HTTP functionality. See "Server" pkg, which wraps this Service to expose its functionalities over HTTP.
type Service struct {
	db Storage
}

// NewService return an initialized Service
func NewService(db Storage) *Service {
	return &Service{
		db: db,
	}
}
