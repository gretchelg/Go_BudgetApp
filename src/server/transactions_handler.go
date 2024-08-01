package server

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
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

func (h *TransactionsHandler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	//id := r.URL.Query().Get("id")
	id := chi.URLParam(r, "tran_id")

	// do get one txn
	txn, err := h.svc.Transactions.GetTransactionByID(id)

	// check for a specific "record not found" error
	if errors.Is(err, models.ErrorNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// check for a general error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// respond w/ success
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(txn)
}
