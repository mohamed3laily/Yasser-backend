package main

import (
	"log"

	"yasser-backend/database"
	"yasser-backend/internal/city"
	"yasser-backend/internal/vendor-group/category"
)

func main() {
	database.Init()

	if err := category.Seed(database.DB); err != nil {
		log.Fatalf("❌ Failed to seed category: %v", err)
	}

	if err := city.Seed(database.DB); err != nil {
		log.Fatalf("❌ Failed to seed city & districts: %v", err)
	}

	log.Println("✅ Seeding completed successfully.")
}
