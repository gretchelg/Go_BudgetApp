package service

import (
	"errors"
	"fmt"
	"github.com/gretchelg/Go_BudgetApp/src/models"
	"github.com/gretchelg/Go_BudgetApp/src/utils/random"
	"time"
)

func (s *Service) GetAllTransactions() ([]models.Transaction, error) {
	return s.db.GetAllTransactions()
}

func (s *Service) GetTransactionByID(id string) (*models.Transaction, error) {
	return s.db.GetTransactionByID(id)
}

func (s *Service) InsertTransaction(txn models.Transaction) error {
	// validate
	if txn.TranCurrency == "" {
		return errors.New("TranCurrency missing")
	}

	if err := txn.TranSign.Validate(); err != nil {
		return err
	}

	if txn.User == "" {
		return errors.New("field User is missing")
	}

	// supply generated values
	txn.TranID = generateTranID()
	txn.TranDate = time.Now() // TODO only generate if empty / zero time

	// TODO return the generated TranID

	// do insert
	_, err := s.db.InsertTransaction(txn)
	return err
}

func (s *Service) UpdateTransaction(tranID string, txn models.Transaction) error {
	// validate
	if txn.TranSign == "" {
		// only validate TranSign if we're asked to update it
		if err := txn.TranSign.Validate(); err != nil {
			return err
		}
	}

	// ensure tranID is consistent
	txn.TranID = tranID

	// do insert
	return s.db.UpdateTransaction(tranID, txn)
}

func generateTranID() string {
	part1 := random.GenerateRandomUUID(4)
	part2 := random.GenerateRandomUUID(4)
	part3 := random.GenerateRandomUUID(4)
	part4 := random.GenerateRandomUUID(4)

	return fmt.Sprintf("%s-%s-%s-%s", part1, part2, part3, part4)
}
