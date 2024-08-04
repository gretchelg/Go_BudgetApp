package plaid

import (
	"context"

	"github.com/gretchelg/Go_BudgetApp/src/models"
	"github.com/gretchelg/Go_BudgetApp/src/service"
)

// Config defines the configurations required for the Plaid Client
type Config struct {
	ClientID string
	Secret   string
}

// PlaidClient provides integration to Plaid service to link with users' external bank accounts.
type PlaidClient struct {
	// TODO
}

// NewPlaidClient returns an initialized Plaid client, that adheres to the service layer's BankDetailsProvider interface
func NewPlaidClient(config Config) (service.BankDetailsProvider, error) {
	// TODO implement me

	return &PlaidClient{}, nil
}

func (p *PlaidClient) GetLatestTransactions(ctx context.Context) ([]models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}
