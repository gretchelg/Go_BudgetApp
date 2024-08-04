package service

import (
	"context"

	"github.com/gretchelg/Go_BudgetApp/src/models"
)

func (s *Service) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.db.GetAllUsers(ctx)
}
