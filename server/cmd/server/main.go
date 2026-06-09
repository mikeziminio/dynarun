package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mikeziminio/dynarun/server/internal/api"
	"github.com/mikeziminio/dynarun/server/internal/repo/postgres"
	"github.com/mikeziminio/dynarun/server/internal/service"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	address := ":8080"

	ctx := context.Background()

	// db pool
	postgresDSN := "postgresql://dynarun:dynarun@127.0.0.1:5432/dynarun"
	pool, err := postgres.Connect(ctx, postgresDSN, 10)
	if err != nil {
		panic(err)
	}
	defer postgres.Close(pool)

	// repos
	modelRepo, err := postgres.NewModelRepo(pool, logger)
	if err != nil {
		log.Fatalf("failed to init model repo: %v", err)
	}

	// services
	modelService, err := service.NewModelService(modelRepo, logger)
	if err != nil {
		log.Fatalf("failed to init model service: %v", err)
	}

	// main server
	s := api.NewServer(address, modelService, logger)
	err = s.Run()
	fmt.Printf("%v", err)
}
