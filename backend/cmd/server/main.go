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
	// cfg := config.Load()
	// db := database.Connect(cfg)
	// server := api.NewServer(db)
	// server.Run()

	cfg := config.Load()
	defer cfg.Logger.Sync()

	db, err := database.NewPool(cfg)
	if err != nil {
		cfg.Logger.Fatal("db init failed", zap.Error(err))
	}

	// pool, err := db.NewPool(
	// 	context.Background(),
	// 	db.Config{
	// 		DSN:             cfg.Database.DSN,
	// 		MaxConns:        10,
	// 		MinConns:        2,
	// 		HealthCheckSecs: 30,
	// 	},
	// 	cfg.Logger,
	// )
	// if err != nil {
	// 	cfg.Logger.Fatal("db init failed", zap.Error(err))
	// }

	mainRepo := repository.NewMainRepository(db)
	mainSvc := service.NewMainService(*mainRepo)
	mainHandler := handler.NewMainHandler(*mainSvc)

	r := api.New(cfg, mainHandler)

	serverAddress := fmt.Sprintf(":%s", cfg.Server.Port)
	r.Run(serverAddress)
}
