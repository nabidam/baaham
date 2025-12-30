package main

import (
	"fmt"

	"github.com/nabidam/baaham/internal/api"
	"github.com/nabidam/baaham/internal/config"
	"github.com/nabidam/baaham/internal/handler"
	"github.com/nabidam/baaham/internal/repository"
	"github.com/nabidam/baaham/internal/service"
	"github.com/nabidam/baaham/pkg/database"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()
	defer cfg.Logger.Sync()

	db, err := database.NewPool(cfg)
	if err != nil {
		cfg.Logger.Fatal("db init failed", zap.Error(err))
	}

	migratiuonErr := database.RunMigrations(db)
	if migratiuonErr != nil {
		cfg.Logger.Fatal("migration failed", zap.Error(migratiuonErr))
	}

	mainRepo := repository.NewMainRepository(db)
	mainSvc := service.NewMainService(mainRepo, cfg)
	mainHandler := handler.NewMainHandler(mainSvc)

	r := api.New(cfg, mainHandler)

	serverAddress := fmt.Sprintf(":%s", cfg.Server.Port)
	r.Run(serverAddress)
}
