package workflows

import (
	"github.com/gretchelg/Go_BudgetApp/src/models"
)

type TransactionsStorage interface {
	GetAllTransactions() ([]models.Transaction, error)
}

type TransactionsWorkflow struct {
	transactionsStorage TransactionsStorage
}

func NewTransactionsWorkflow(transactionsStorage TransactionsStorage) *TransactionsWorkflow {
	return &TransactionsWorkflow{
		transactionsStorage: transactionsStorage,
	}
}

func (s *TransactionsWorkflow) GetAllTransactions() ([]models.Transaction, error) {
	return s.transactionsStorage.GetAllTransactions()
}