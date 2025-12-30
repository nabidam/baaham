package route

import (
	"github.com/gin-gonic/gin"
	"github.com/nabidam/baaham/internal/handler"
)

func RegisterAuthRoutes(api gin.IRoutes, h *handler.AuthHandler) {
	api.POST("/login", h.Login)
	// api.POST("/register", h.Register)
}
