package server

import (
	"encoding/json"
	"github.com/gretchelg/Go_BudgetApp/src/models"
	"github.com/gretchelg/Go_BudgetApp/src/service"
	"net/http"
)

// TransactionsHandler handles all requests related to transactions
type TransactionsHandler struct {
	svc *service.Service
}

func NewTransactionsHandler(svc *service.Service) *TransactionsHandler {
	return &TransactionsHandler{
		svc: svc,
	}
}

type GetAllTransactionsResponse struct {
	Data []models.Transaction `json:"data"`
}

func (h *TransactionsHandler) GetAllTransactions(w http.ResponseWriter, _ *http.Request) {
	// do get all txns
	txns, err := h.svc.Transactions.GetAllTransactions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// respond w/ success
	resp := GetAllTransactionsResponse{
		Data: txns,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
