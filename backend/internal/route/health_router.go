package route

import (
	"github.com/gin-gonic/gin"
	"github.com/nabidam/baaham/internal/handler"
)

func RegisterHelathRoutes(api gin.IRoutes, h *handler.HealthHandler) {
	api.GET("/health", h.HealthCheck)
}
