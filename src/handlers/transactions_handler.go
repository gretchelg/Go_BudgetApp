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

// NewTransactionsHandler is the constructor for TransactionsHandler, which handles
// all HTTP request relating to Transactions functionality
func NewTransactionsHandler(svc *service.Service) *TransactionsHandler {
	return &TransactionsHandler{
		svc: svc,
	}
}

// GetAllTransactionsResponse defines the schema for GetAllTransactions Response
type GetAllTransactionsResponse struct {
	Data []models.Transaction `json:"data"`
}

func (t *TransactionsHandler) GetAllTransactions(w http.ResponseWriter, _ *http.Request) {
	// do get all txns
	txns, err := t.svc.GetAllTransactions()
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

func (t *TransactionsHandler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	//id := r.URL.Query().Get("id")
	tranID := chi.URLParam(r, "tran_id")

	// do get one txn
	txn, err := t.svc.GetTransactionByID(tranID)

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

// PostTransactionResponse defines the schema for the PostTransaction Response
type PostTransactionResponse struct {
	TranID string `json:"tran_id"`
}

func (t *TransactionsHandler) PostTransaction(w http.ResponseWriter, r *http.Request) {
	// read request body
	var request models.Transaction
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errMsg := fmt.Sprintf("unable to decode request body: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	// do insert
	generatedTranID, err := t.svc.InsertTransaction(request)
	if err != nil {
		errMsg := fmt.Sprintf("insert failed: %s", err.Error())
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	// write response header
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// write response body
	resp := PostTransactionResponse{
		TranID: generatedTranID,
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func (t *TransactionsHandler) PatchTransaction(w http.ResponseWriter, r *http.Request) {
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
	err = t.svc.UpdateTransaction(tranID, request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// respond
	w.WriteHeader(http.StatusOK)
}
