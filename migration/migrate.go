package migration

import (
	"log"
	"yasser-backend/database"
	"yasser-backend/internal/city"
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

	)
	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
}
