package api

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	"github.com/nabidam/baaham/internal/config"
	"github.com/nabidam/baaham/internal/handler"
	"github.com/nabidam/baaham/internal/route"
)

func New(
	cfg *config.Config,
	h *handler.MainHandler,
) *gin.Engine {
	r := gin.New()

	// log all requests
	r.Use(ginzap.Ginzap(cfg.Logger, time.RFC3339, true))
	// log panics
	r.Use(ginzap.RecoveryWithZap(cfg.Logger, true))

	route.RegisterRoutes(r, h)

	return r
}
