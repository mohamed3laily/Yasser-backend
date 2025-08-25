package main

import (
	"log"

	"yasser-backend/bootstrap"
	"yasser-backend/internal/city"
	"yasser-backend/internal/vendor-group/category"
	"yasser-backend/internal/vendor-group/vendor"
	"yasser-backend/seeder"
)

func main() {
	deps := bootstrap.NewDependencies()

	if err := category.Seed(deps.DB); err != nil {
		log.Fatalf("❌ Failed to seed category: %v", err)
	}

	if err := city.Seed(deps.DB); err != nil {
		log.Fatalf("❌ Failed to seed city & districts: %v", err)
	}

	if err := vendor.Seed(deps.DB); err != nil {
		log.Fatalf("❌ Failed to seed vendors: %v", err)
	}

	if err := seeder.SeedItemsAndCategories(deps.DB); err != nil {
		log.Fatalf("❌ Failed to seed items: %v", err)
	}

	log.Println("✅ Seeding completed successfully.")
}
