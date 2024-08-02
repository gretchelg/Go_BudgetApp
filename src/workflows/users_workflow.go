package workflows

import "github.com/gretchelg/Go_BudgetApp/src/models"

type UsersStorage interface {
	GetAllUsers() ([]models.User, error)
}

type UsersWorkflow struct {
	storage UsersStorage
}

func NewUsersWorkflow(storage UsersStorage) *UsersWorkflow {
	return &UsersWorkflow{
		storage: storage,
	}
}

func (u *UsersWorkflow) GetAllUsers() ([]models.User, error) {
	return u.storage.GetAllUsers()
}
