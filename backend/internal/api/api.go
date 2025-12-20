package api

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/nabidam/baaham/cmd/docs"

	"github.com/nabidam/baaham/internal/config"
	"github.com/nabidam/baaham/internal/handler"
	"github.com/nabidam/baaham/internal/route"
)

func New(
	cfg *config.Config,
	h *handler.MainHandler,
) *gin.Engine {
	r := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1"

	// log all requests
	r.Use(ginzap.Ginzap(cfg.Logger, time.RFC3339, true))
	// log panics
	r.Use(ginzap.RecoveryWithZap(cfg.Logger, true))

	route.RegisterRoutes(r, h)

	// swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
