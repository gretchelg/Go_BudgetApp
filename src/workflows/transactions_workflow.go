package workflows

import (
	"errors"
	"github.com/gretchelg/Go_BudgetApp/src/models"
	"time"
)

// TransactionsStorage defines methods require of a storage for Transactions, as in a DB.
type TransactionsStorage interface {
	GetAllTransactions() ([]models.Transaction, error)
	GetTransactionByID(id string) (*models.Transaction, error)
	InsertTransaction(txn models.Transaction) (string, error)
}

// TransactionsWorkflow provides functionality for transactions-related workflows
type TransactionsWorkflow struct {
	storage TransactionsStorage
}

// NewTransactionsWorkflow is the constructor function for TransactionsWorkflow type
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

func (t *TransactionsWorkflow) InsertTransaction(txn models.Transaction) error {
	// validate
	if txn.TranCurrency == "" {
		return errors.New("TranCurrency missing")
	}

	if txn.TranSign == "" {
		return errors.New("TranSign missing")
	}

	if txn.User == "" {
		return errors.New("field User is missing")
	}

	// supply generated values
	txn.TranID = generateTranID()
	txn.TranDate = time.Now()

	// do insert
	_, err := t.storage.InsertTransaction(txn)
	return err
}

func generateTranID() string {
	return "as2lnfdasvlasmea233"
}
