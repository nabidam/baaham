package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IHealthHandler interface {
	HealthCheck(*gin.Context)
}

type IHealthService interface {
	HealthCheck(ctx context.Context) (bool, error)
}

type HealthHandler struct {
	svc IHealthService
}

func NewHealthHandler(svc IHealthService) IHealthHandler {
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
