package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gretchelg/Go_BudgetApp/src/models"
	"github.com/gretchelg/Go_BudgetApp/src/utils/random"
)

// GetAllTransactions returns all transactions. An optional Filter object can be provided to limit the results
func (s *Service) GetAllTransactions(ctx context.Context, filter *models.TransactionsFilter) ([]models.Transaction, error) {
	return s.db.GetAllTransactions(ctx, filter)
}

func (s *Service) GetTransactionByID(ctx context.Context, tranID string) (*models.Transaction, error) {
	return s.db.GetTransactionByID(ctx, tranID)
}

// InsertTransaction inserts the given Transaction into storage, and returns the generated tranID
func (s *Service) InsertTransaction(ctx context.Context, txn models.Transaction) (newTranID string, e error) {
	// validate
	if txn.TranCurrency == "" {
		return "", errors.New("TranCurrency missing")
	}

	if err := txn.TranSign.Validate(); err != nil {
		return "", err
	}

	if txn.User == "" {
		return "", errors.New("field User is missing")
	}

	// supply generated values
	txn.TranID = generateTranID()

	if txn.TranDate.IsZero() {
		// only generate if empty / zero
		txn.TranDate = time.Now()
	}

	// do insert
	_, err := s.db.InsertTransaction(ctx, txn)
	return txn.TranID, err
}

func (s *Service) UpdateTransaction(ctx context.Context, tranID string, txn models.Transaction) error {
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
	return s.db.UpdateTransaction(ctx, tranID, txn)
}

func generateTranID() string {
	part1 := random.GenerateRandomUUID(4)
	part2 := random.GenerateRandomUUID(4)
	part3 := random.GenerateRandomUUID(4)
	part4 := random.GenerateRandomUUID(4)

	return fmt.Sprintf("%s-%s-%s-%s", part1, part2, part3, part4)
}
