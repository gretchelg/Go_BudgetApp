package service

import (
	"context"

	"github.com/gretchelg/Go_BudgetApp/src/models"
)

// StorageProvider defines methods required of the storage layer (database).
type StorageProvider interface {
	// Transactions-related
	GetAllTransactions(ctx context.Context, filter *models.TransactionsFilter) ([]models.Transaction, error)
	GetTransactionByID(ctx context.Context, tranID string) (*models.Transaction, error)
	InsertTransaction(ctx context.Context, txn models.Transaction) (string, error)
	UpdateTransaction(ctx context.Context, tranID string, txn models.Transaction) error

	// Users-related
	GetAllUsers(ctx context.Context) ([]models.User, error)
}

// BankDetailsProvider defines methods required of external banking integrations provider
type BankDetailsProvider interface {
	GetLatestTransactions(ctx context.Context) ([]models.Transaction, error)
}

// Service defines the core service of our app, and provides access to the underlying Workflow functionalities
// It has no HTTP functionality. See "Server" pkg, which wraps this Service to expose its functionalities over HTTP.
type Service struct {
	db    StorageProvider
	banks BankDetailsProvider
}

// NewService return an initialized Service
func NewService(db StorageProvider, banks BankDetailsProvider) *Service {
	return &Service{
		db:    db,
		banks: banks,
	}
}
