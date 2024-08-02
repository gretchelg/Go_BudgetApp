package server

import (
	"encoding/json"
	"github.com/gretchelg/Go_BudgetApp/src/models"
	"github.com/gretchelg/Go_BudgetApp/src/service"
	"net/http"
)

// UsersHandler handles all requests related to users
type UsersHandler struct {
	svc *service.Service
}

func NewUsersHandler(svc *service.Service) *UsersHandler {
	return &UsersHandler{
		svc: svc,
	}
}

type GetAllUsersResponse struct {
	Data []models.User `json:"data"`
}

func (h *UsersHandler) GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	// do get all users
	users, err := h.svc.Users.GetAllUsers()
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
