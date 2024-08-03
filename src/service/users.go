package service

import "github.com/gretchelg/Go_BudgetApp/src/models"

func (s *Service) GetAllUsers() ([]models.User, error) {
	return s.db.GetAllUsers()
}
