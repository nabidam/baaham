package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nabidam/baaham/internal/domain"
)

type AuthHandler struct {
	svc domain.AuthService
}

func NewAuthHandler(svc domain.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

// @Summary	Login user
// @Schemes
// @Description	Login user with username and password
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			login	body		domain.LoginRequest	true	"Login credentials"
// @Success		200		{object}	domain.LoginResponse
// @Produce		json
// @Router			/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp, err := h.svc.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong credentials"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
