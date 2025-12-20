package route

import (
	"github.com/gin-gonic/gin"
	"github.com/nabidam/baaham/internal/handler"
)

func RegisterHelathRoutes(r *gin.Engine, h *handler.HealthHandler) {
	r.GET("/health", h.HealthCheck)
}
