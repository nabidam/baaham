package route

import (
	"github.com/gin-gonic/gin"
	"github.com/nabidam/baaham/internal/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.MainHandler) {
	// Define routes
	api := r.Group("/api/v1")
	{
		// Health check
		RegisterHelathRoutes(api, h.HealthHandler)

	}

	// Health check (no auth)
	// r.GET("/health")

	// // API versioning
	// api := r.Group("/api")

	// v1 := api.Group("/v1")
	// {
	// 	v1.GET("/users", handler.ListUsers)
	// 	v1.GET("/users/:id", handler.GetUser)
	// }

	// // Authenticated routes
	// auth := api.Group("/auth")
	// auth.Use(middleware.Auth())
	// {
	// 	auth.POST("/logout", handler.Logout)
	// }

	// // Admin routes
	// admin := api.Group("/admin")
	// admin.Use(
	// 	middleware.Auth(),
	// 	middleware.AdminOnly(),
	// )
	// {
	// 	admin.GET("/users", handler.AdminListUsers)
	// }
}
