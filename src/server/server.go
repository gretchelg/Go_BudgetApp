package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/gretchelg/Go_BudgetApp/src/service"
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
		_, _ = w.Write([]byte("OK"))
	})

	//r.Mount("/books", BookRoutes())
	transactionsHandler := NewTransactionsHandler(s.svc)
	r.Get("/transactions", transactionsHandler.GetAllTransactions)

	// start listening.
	log.Print("Listening on :3000...")
	return http.ListenAndServe(":3000", r)
}

// Start starts the server and handles requests over HTTP.
// It is a blocking call, and will only return control to the caller once server has shut down.
//func (s *Server) StartBkup() error {
//	// setup HTTP server + handlers
//	mux := http.NewServeMux()
//
//	finalHandler := http.HandlerFunc(final)
//	mux.Handle("/", middleware.EnforceJSONHandler(finalHandler))
//
//	mux.Handle("/transactions", middleware.EnforceJSONHandler(finalHandler))
//
//	// start listening.
//	log.Print("Listening on :3000...")
//	return http.ListenAndServe(":3000", mux)
//}

//func final(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("OK"))
//}
