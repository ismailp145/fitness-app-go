// internal/delivery/http/router.go
package http

import (
    "your-project/internal/delivery/http/handler"
    "your-project/internal/delivery/http/middleware"
    "your-project/internal/usecase"
    
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

