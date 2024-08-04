package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gretchelg/Go_BudgetApp/src/models"
	"github.com/gretchelg/Go_BudgetApp/src/service"
)

// UsersHandler handles all requests related to users
type UsersHandler struct {
	svc *service.Service
}

// NewUsersHandler is the constructor for UsersHandler, which handles
// all HTTP request relating to Users functionality
func NewUsersHandler(svc *service.Service) *UsersHandler {
	return &UsersHandler{
		svc: svc,
	}
}

// GetAllUsersResponse defines the schema for GetAllUsers Response
type GetAllUsersResponse struct {
	Data []models.User `json:"data"`
}

func (u *UsersHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// do get all users
	ctx := r.Context()
	users, err := u.svc.GetAllUsers(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// respond w/ success
	resp := GetAllUsersResponse{
		Data: users,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
