package migration

import (
	"log"
	"yasser-backend/internal/city"
	"yasser-backend/internal/item-group/item"
	itemcategory "yasser-backend/internal/item-group/item-category"
	"yasser-backend/internal/user"
	"yasser-backend/internal/vendor-group/category"
	"yasser-backend/internal/vendor-group/vendor"

	"gorm.io/gorm"
)


func Migrate(db *gorm.DB) {
	log.Println("🚀 Running database migrations...")

	err := db.AutoMigrate(
		&user.User{},
		&category.VendorCategory{},
		&city.City{},
		&city.District{},
		&vendor.Vendor{},
		&item.Item{},
		&item.ItemAddon{},
		&item.ItemSize{},
		&item.ItemVariant{},
		&itemcategory.ItemsCategory{},
	)
	if err != nil {
		log.Fatalf("❌ Database migration failed: %v", err)
	}

	log.Println("✅ Migrations completed successfully.")
}
