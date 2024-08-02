package workflows

import "github.com/gretchelg/Go_BudgetApp/src/models"

type UsersStorage interface {
	GetAllUsers() ([]models.User, error)
}

type UsersWorkflow struct {
	userStorage UsersStorage
}

func NewUsersWorkflow(transactionsStorage UsersStorage) *UsersWorkflow {
	return &UsersWorkflow{
		userStorage: transactionsStorage,
	}
}

func (s *UsersWorkflow) GetAllUsers() ([]models.User, error) {
	return s.userStorage.GetAllUsers()
}
