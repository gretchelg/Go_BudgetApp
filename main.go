package main

import (
	"log"

	"github.com/vrischmann/envconfig"

	"github.com/gretchelg/Go_BudgetApp/src/server"
	"github.com/gretchelg/Go_BudgetApp/src/service"
)

func main() {
	// load env vars
	var config service.Config
	if err := envconfig.Init(&config); err != nil {
		log.Fatalf("failed to load env vars: %s", err)
	}

	// setup core service
	svc, err := service.NewService(config)
	if err != nil {
		log.Fatalf("failed to initialize service layer: %s", err)
	}

	// setup HTTP server + handlers
	svr, err := server.NewServer(svc)
	if err != nil {
		log.Fatalf("failed to initialize HTTP server layer: %s", err)
	}

	err = svr.Start()
	log.Fatalf("terminating server: %s", err)
}
