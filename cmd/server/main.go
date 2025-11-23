// cmd/server/main.go
package main

import (
	"database/sql"
	"log"
	"os"

	"fitness-app-go/config"
	httpDelivery "fitness-app-go/internal/delivery/http"
	"fitness-app-go/internal/repository/postgres"
	"fitness-app-go/internal/usecase"

	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Initialize repositories (infrastructure layer)
	userRepo := postgres.NewUserRepository(db)

	// Initialize use cases (business logic layer)
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Initialize HTTP server (delivery layer)
	router := httpDelivery.SetupRouter(userUseCase)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
