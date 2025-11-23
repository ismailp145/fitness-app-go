// config/config.go
package config

import "os"

type Config struct {
    DatabaseURL string
    Port        string
    JWTSecret   string
}

func Load() *Config {
    return &Config{
        DatabaseURL: getEnv("DATABASE_URL", "postgres://user:pass@localhost:5432/dbname?sslmode=disable"),
        Port:        getEnv("PORT", "8080"),
        JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

