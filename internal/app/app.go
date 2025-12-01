package app

import (
	"log"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/auth"
	"github.com/Vladimirmoscow84/Warehouse_Control/internal/cfg"
	"github.com/Vladimirmoscow84/Warehouse_Control/internal/handlers"
	"github.com/Vladimirmoscow84/Warehouse_Control/internal/service"
	"github.com/Vladimirmoscow84/Warehouse_Control/internal/storage"
	"github.com/Vladimirmoscow84/Warehouse_Control/internal/storage/postgres"
	"github.com/wb-go/wbf/ginext"
)

func Run() {

	config := cfg.Load()

	postgresStore, err := postgres.New(config.DatabaseURI)
	if err != nil {
		log.Fatalf("[app]failed to connect to PG DB: %v", err)
	}
	defer postgresStore.Close()
	log.Println("[app] Connected to Postgres successfully")

	store, err := storage.New(postgresStore, postgresStore, postgresStore, postgresStore)
	if err != nil {
		log.Fatalf("[app] failed to init unified storage: %v", err)
	}
	log.Println("[app]storage initialized successfully")

	service, err := service.New(store, store, store, store)
	if err != nil {
		log.Fatalf("[app] failed to init service: %v", err)
	}
	log.Println("[app] Service initialized successfully")

	authService := auth.New(config.JWTSecret)
	engine := ginext.New("release")
	router := handlers.New(engine, service, service, service, service, service, service, authService)
	router.Routes(config.JWTSecret)

	log.Printf("[app] starting server on %s", config.ServerAddress)
	err = engine.Run(config.ServerAddress)
	if err != nil {
		log.Fatalf("[app] server failed: %v", err)
	}
}
