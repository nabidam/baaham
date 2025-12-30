package handler

import "github.com/nabidam/baaham/internal/service"

type MainHandler struct {
	HealthHandler *HealthHandler
	AuthHandler   *AuthHandler
}

func NewMainHandler(mainSvc *service.MainService) *MainHandler {
	healthHandler := NewHealthHandler(mainSvc.HealthService)
	authHandler := NewAuthHandler(mainSvc.AuthService)

	return &MainHandler{HealthHandler: healthHandler, AuthHandler: authHandler}
}
