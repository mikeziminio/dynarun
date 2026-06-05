package main

import (
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

	// repos
	modelRepo, err := postgres.NewModelRepo(logger)
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
