package main

import (
	"github.com/gretchelg/Go_BudgetApp/src/data_sources/mongodb"
	"github.com/gretchelg/Go_BudgetApp/src/service"
	"log"

	"github.com/vrischmann/envconfig"

	"github.com/gretchelg/Go_BudgetApp/src/handlers"
)

// AppConfig defines the configurations required to run this app
type AppConfig struct {
	MongoURI string
	Plaid    struct {
		Secret string
		Env    string
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
	// setup dependencies
	db, err := mongodb.NewClient(config.MongoURI)
	if err != nil {
		return nil, err
	}

	// setup service
	svc := service.NewService(db)

	// respond with ready service
	return svc, nil
}
