package cmd

import (
	"os"

	"github.com/nabidam/baaham/internal/config"
	"github.com/nabidam/baaham/internal/domain"
	"github.com/nabidam/baaham/internal/repository"
	"github.com/nabidam/baaham/pkg/database"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	repo domain.UserRepository
)

var rootCmd = &cobra.Command{
	Use:   "usercli",
	Short: "User management CLI",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.Load()
		defer cfg.Logger.Sync()

		db, err := database.NewPool(cfg)
		if err != nil {
			cfg.Logger.Fatal("db init failed", zap.Error(err))
		}

		repo = repository.NewUserRepository(db)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
