package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
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
	txns, err := h.svc.GetAllTransactions()
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
	tranID := chi.URLParam(r, "tran_id")

	// do get one txn
	txn, err := h.svc.GetTransactionByID(tranID)

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

func (h *TransactionsHandler) PostTransaction(w http.ResponseWriter, r *http.Request) {
	// read request body
	var request models.Transaction
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errMsg := fmt.Sprintf("unable to decode request body: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	// do insert
	err = h.svc.InsertTransaction(request)
	if err != nil {
		errMsg := fmt.Sprintf("insert failed: %s", err.Error())
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	// respond w/ success
	w.WriteHeader(http.StatusCreated)
}

func (h *TransactionsHandler) PatchTransaction(w http.ResponseWriter, r *http.Request) {
	// get the tran ID from the url path
	tranID := chi.URLParam(r, "tran_id")

	// read request body
	var request models.Transaction
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errMsg := fmt.Sprintf("unable to decode request body: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	// do update/patch txn
	err = h.svc.UpdateTransaction(tranID, request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// respond
	w.WriteHeader(http.StatusOK)
}
