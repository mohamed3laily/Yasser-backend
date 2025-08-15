package migration

import (
	"log"
	"yasser-backend/database"
	"yasser-backend/internal/city"
	"yasser-backend/internal/item-group/item"
	itemcategory "yasser-backend/internal/item-group/item-category"
	"yasser-backend/internal/user"
	"yasser-backend/internal/vendor-group/category"
	"yasser-backend/internal/vendor-group/vendor"
)

func Migrate() {
	err := database.DB.AutoMigrate(
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
		log.Fatalf("Database migration failed: %v", err)
	}
}
