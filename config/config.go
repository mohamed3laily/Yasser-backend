package config

import (
	"os"

	"github.com/joho/godotenv"
)

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

func Load() *Config {
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

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
