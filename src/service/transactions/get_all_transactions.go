package transactions

import "github.com/gretchelg/Go_BudgetApp/src/models"

func (s *Service) GetAllTransactions() ([]models.Transaction, error) {
	return s.txnRepo.GetAllTransactions()
}
