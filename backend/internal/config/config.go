package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Server struct {
		Port string
	}

	Database struct {
		Host     string
		Port     string
		UserName string
		Password string
		DBName   string
		DSN      string
	}

	AppEnv    string
	JWTSecret string

	Logger *zap.Logger
}

func Load() *Config {
	v := viper.New()

	v.SetConfigFile(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.SetDefault("APP_ENV", "development")
	v.SetDefault("SERVER_PORT", "8080")
	v.SetDefault("DATABASE_HOST", "localhost")
	v.SetDefault("DATABASE_PORT", "5432")

	if err := v.ReadInConfig(); err != nil {
		log.Println("config: no .env file found, relying on env vars")
	}

	logger := buildLogger(v.GetString("APP_ENV"))

	cfg := &Config{
		AppEnv: v.GetString("APP_ENV"),
		Logger: logger,
	}

	cfg.Server.Port = v.GetString("SERVER_PORT")

	cfg.Database.Host = v.GetString("DATABASE_HOST")
	cfg.Database.Port = v.GetString("DATABASE_PORT")
	cfg.Database.UserName = v.GetString("DATABASE_USERNAME")
	cfg.Database.Password = v.GetString("DATABASE_PASSWORD")
	cfg.Database.DBName = v.GetString("DATABASE_DBNAME")

	cfg.Database.DSN = fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		cfg.Database.UserName, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName,
	)

	cfg.JWTSecret = v.GetString("JWT_SECRET")

	validate(cfg)

	logger.Info("config loaded",
		zap.String("env", cfg.AppEnv),
		zap.String("port", cfg.Server.Port),
	)

	return cfg
}

func buildLogger(env string) *zap.Logger {
	if env == "production" {
		logger, _ := zap.NewProduction()
		return logger
	}

	logger, _ := zap.NewDevelopment()
	return logger
}

func validate(cfg *Config) {
	missing := []string{}

	if cfg.Database.Host == "" {
		missing = append(missing, "DATABASE_HOST")
	}
	if cfg.JWTSecret == "" {
		missing = append(missing, "JWT_SECRET")
	}

	if len(missing) > 0 {
		log.Fatalf("missing required config values: %s", strings.Join(missing, ", "))
	}
}
