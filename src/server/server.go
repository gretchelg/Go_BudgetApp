package server

import (
	"log"
	"net/http"

	"github.com/gretchelg/Go_BudgetApp/src/models/handlers/middleware"
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
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", middleware.EnforceJSONHandler(finalHandler))

	log.Print("Listening on :3000...")

	// start listening.
	return http.ListenAndServe(":3000", mux)
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
