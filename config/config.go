package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds all typed configuration for the application.

// WATTSI_INSTANCE_ID=68910B3A4BD70
// WATTSI_ACCESS_TOKEN=6891055b330cf
// JWT_SECRET="your_jwt_secret_here"

// DB_HOST=localhost
// DB_PORT=5432
// DB_USER=educatly
// DB_PASSWORD=password1
// DB_NAME=yasser-dev

// DATABASE_URL=postgresql://educatly:password1@localhost:5432/yasser-dev?sslmode=disable

// MEILI_HOST=http://localhost:7700
type Config struct {
	// Server
	Port string

	// Database
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string

	// Meilisearch
	MeiliHost      string
	MeiliAPIKey    string
	MeiliIndexName string

	// WattsI
	WattsiInstanceID  string
	WattsiAccessToken string

	// JWT
	JWTSecret string
}

// Load creates a new Config struct and populates it from environment variables.
func Load() *Config {
	// Load .env file. It's safe to ignore the error if it doesn't exist.
	_ = godotenv.Load()

	return &Config{
		Port:           getEnv("PORT", "3000"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBUser:         getEnv("DB_USER", ""),
		DBPassword:     getEnv("DB_PASSWORD", ""),
		DBName:         getEnv("DB_NAME", ""),
		DBPort:         getEnv("DB_PORT", "5432"),
		MeiliHost:      getEnv("MEILI_HOST", "http://127.0.0.1:7700" ),
		MeiliAPIKey:    getEnv("MEILI_API_KEY", ""),
		MeiliIndexName: getEnv("MEILI_INDEX_NAME", "Search"),
		JWTSecret:      getEnv("JWT_SECRET", "your-default-secret"),

	}
}

// getEnv is a helper to read an environment variable or return a fallback.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
