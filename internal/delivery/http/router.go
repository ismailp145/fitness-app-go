// internal/delivery/http/router.go
package http

import (
	"fitness-app-go/internal/delivery/http/handler"
	"fitness-app-go/internal/delivery/http/middleware"
	"fitness-app-go/internal/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userUseCase usecase.UserUseCase) *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API v1
	v1 := router.Group("/api/v1")
	{
		// User routes
		userHandler := handler.NewUserHandler(userUseCase)
		userHandler.RegisterRoutes(v1)
	}

	return router
}
