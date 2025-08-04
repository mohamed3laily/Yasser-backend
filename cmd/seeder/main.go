package main

import (
	"log"

	"yasser-backend/database"
	"yasser-backend/internal/vendor-group/category"
)

func main() {
	database.Init()

	if err := category.Seed(database.DB); err != nil {
		log.Fatalf("❌ Failed to seed category: %v", err)
	}

	log.Println("✅ Seeding completed successfully.")
}
