package api

import (
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
	// r.Use(gin.Recovery())

	route.RegisterRoutes(r, h)

	return r
}
