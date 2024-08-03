package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/gretchelg/Go_BudgetApp/src/service"
)

const (
	port = 3000
)

// Server wraps our Service, to expose its functionality over HTTP request/response
type Server struct {
	svc *service.Service
}

// NewServer wraps the given Service, to expose its functionality over HTTP request/response.
func NewServer(svc *service.Service) (*Server, error) {
	return &Server{
		svc: svc,
	}, nil
}

// Start starts the server and handles requests over HTTP.
// It is a blocking call, and will only return control to the caller once server has shut down.
func (s *Server) Start() error {
	// setup HTTP server + handlers
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("welcome to home page"))
	})

	// setup handlers for the transactions-related routes
	transactionsHandler := NewTransactionsHandler(s.svc)
	r.Get("/api/v1/transaction", transactionsHandler.GetAllTransactions)
	r.Get("/api/v1/transaction/{tran_id}", transactionsHandler.GetTransactionByID)
	r.Post("/api/v1/transaction", transactionsHandler.PostTransaction)

	// setup handlers for the users-related routes
	usersHandler := NewUsersHandler(s.svc)
	r.Get("/api/v1/user", usersHandler.GetAllUsers)

	// start listening.
	log.Print(fmt.Sprintf("Listening on :%d...", port))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
