package main

import (
	"fmt"
	"log"

	"github.com/vrischmann/envconfig"

	"github.com/gretchelg/Go_BudgetApp/src/data_sources/mongodb"
	"github.com/gretchelg/Go_BudgetApp/src/data_sources/plaid"
	"github.com/gretchelg/Go_BudgetApp/src/handlers"
	"github.com/gretchelg/Go_BudgetApp/src/service"
)

// AppConfig defines the configurations required to run this app
type AppConfig struct {
	MongoURI string
	Plaid    struct {
		ClientId string
		Secret   string
		Env      string
	}
}

func main() {
	// load env vars
	var config AppConfig
	if err := envconfig.Init(&config); err != nil {
		log.Fatalf("failed to load env vars: %s", err)
	}

	// setup core service
	svc, err := initService(config)
	if err != nil {
		log.Fatalf("failed to initialize service layer: %s", err)
	}

	// setup HTTP server + handlers.
	// the Server wraps the Service, to expose its functionalities over HTTP
	svr, err := handlers.NewServer(svc)
	if err != nil {
		log.Fatalf("failed to initialize HTTP server layer: %s", err)
	}

	err = svr.Start()
	log.Fatalf("terminating server: %s", err)
}

// initService bootstraps the service layer and returns a ready service.
func initService(config AppConfig) (*service.Service, error) {
	// setup dependencies: database
	db, err := mongodb.NewDBClient(config.MongoURI)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize DB client: %s", err)
	}

	// setup dependencies: plaid
	plaidConfig := plaid.Config{
		ClientID: config.Plaid.ClientId,
		Secret:   config.Plaid.Secret,
	}

	plaidClient, err := plaid.NewPlaidClient(plaidConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Plaid client: %s", err)
	}

	// setup service
	svc := service.NewService(db, plaidClient)

	// respond with ready service
	return svc, nil
}
