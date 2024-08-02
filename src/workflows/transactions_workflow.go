package workflows

import (
	"github.com/gretchelg/Go_BudgetApp/src/models"
)

type TransactionsStorage interface {
	GetAllTransactions() ([]models.Transaction, error)
	GetTransactionByID(id string) (*models.Transaction, error)
}

type TransactionsWorkflow struct {
	storage TransactionsStorage
}

func NewTransactionsWorkflow(storage TransactionsStorage) *TransactionsWorkflow {
	return &TransactionsWorkflow{
		storage: storage,
	}
}

func (t *TransactionsWorkflow) GetAllTransactions() ([]models.Transaction, error) {
	return t.storage.GetAllTransactions()
}

func (t *TransactionsWorkflow) GetTransactionByID(id string) (*models.Transaction, error) {
	return t.storage.GetTransactionByID(id)
}
