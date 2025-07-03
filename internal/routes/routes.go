package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dsrose15/device-monitor/internal/api"
)

func SetupRoutes(router *gin.Engine, db *pgxpool.Pool) {
	// Initialize handlers
	h := api.NewHandlers(db)

	// Health check route
	router.GET("/health", h.HealthCheck)

	// API routes
	api := router.Group("/api/v1")
	{
		// User routes
		users := api.Group("/users")
		{
			users.GET("", h.GetUsers)
			users.GET("/:id", h.GetUser)
			users.POST("", h.CreateUser)
			users.PUT("/:id", h.UpdateUser)
			users.DELETE("/:id", h.DeleteUser)
		}

		// Example additional routes
		api.GET("/ping", h.Ping)
	}

	// Static files (if needed)
	router.Static("/static", "./static")

	// Serve HTML templates (if needed)
	router.LoadHTMLGlob("templates/*")
	router.GET("/", h.Index)
}
