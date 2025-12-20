package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nabidam/baaham/internal/domain"
)

type HealthHandler struct {
	svc domain.HealthService
}

func NewHealthHandler(svc domain.HealthService) *HealthHandler {
	return &HealthHandler{svc: svc}
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	status, err := h.svc.HealthCheck(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return
	}

	if status {
		c.JSON(http.StatusOK, gin.H{"status": "healthy", "db": "up"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "unhealthy", "db": "down"})
	}
}
