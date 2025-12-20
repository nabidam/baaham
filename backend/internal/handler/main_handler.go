package handler

import "github.com/nabidam/baaham/internal/service"

type MainHandler struct {
	HealthHandler *HealthHandler
}

func NewMainHandler(mainSvc *service.MainService) *MainHandler {
	healthHandler := NewHealthHandler(mainSvc.HealthService)

	return &MainHandler{HealthHandler: healthHandler}
}
