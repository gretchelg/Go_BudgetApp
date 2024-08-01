package transactions

import (
	"github.com/gretchelg/Go_BudgetApp/src/models"
)

type TransactionsRepository interface {
	GetAllTransactions() ([]models.Transaction, error)
}

type Service struct {
	txnRepo TransactionsRepository
}

func NewService(txnRepo TransactionsRepository) *Service {
	return &Service{
		txnRepo: txnRepo,
	}
}

func (s *Service) GetAllTransactions() ([]models.Transaction, error) {
	return s.txnRepo.GetAllTransactions()
}
