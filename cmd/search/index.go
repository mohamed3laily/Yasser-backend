package main

import (
	"fmt"
	"log"
	"yasser-backend/bootstrap"
	"yasser-backend/internal/search"
)

func main() {
	log.Println("🚀 Starting indexing process...")

	deps := bootstrap.NewDependencies()

	searchRepo := search.NewRepository(deps.DB)
	searchService := search.NewService(deps.SearchClient, searchRepo)

	log.Println("📚 Fetching data and indexing in Meilisearch...")
	if err := searchService.IndexData(); err != nil {
		log.Fatalf("❌ Indexing failed: %v", err)
	}

	fmt.Println("✅ Indexing completed successfully.")
}