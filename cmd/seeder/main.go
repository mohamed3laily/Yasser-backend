package main

import (
	"log"

	"yasser-backend/database"
	"yasser-backend/internal/city"
	"yasser-backend/internal/vendor-group/category"
	"yasser-backend/internal/vendor-group/vendor"
)

func main() {
	database.Init()

	if err := category.Seed(database.DB); err != nil {
		log.Fatalf("❌ Failed to seed category: %v", err)
	}

	if err := city.Seed(database.DB); err != nil {
		log.Fatalf("❌ Failed to seed city & districts: %v", err)
	}

	if err := vendor.Seed(database.DB); err != nil {
		log.Fatalf("❌ Failed to seed vendors: %v", err)
	}

	log.Println("✅ Seeding completed successfully.")
}
