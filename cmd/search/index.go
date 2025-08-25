package main

import (
	"fmt"
	"log"
	"os"

	"yasser-backend/database"
	"yasser-backend/internal/search"
)

func main() {
	database.Init()

	meiliHost := getEnv("MEILI_HOST", "http://127.0.0.1:7700")
	meiliKey := getEnv("MEILI_API_KEY", "")

	client := search.NewClient(meiliHost, meiliKey, "Search")
	repo := search.NewRepository(database.DB)
	service := search.NewService(client, repo)

	if err := service.IndexData(); err != nil {
		log.Fatalf("❌ Indexing failed: %v", err)
	}

	fmt.Println("✅ Indexing completed successfully")
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
