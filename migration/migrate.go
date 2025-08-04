package migration

import (
	"log"
	"yasser-backend/database"
	"yasser-backend/internal/user"
	"yasser-backend/internal/vendor-group/category"
)

func Migrate() {
	err := database.DB.AutoMigrate(
		&user.User{},
		&category.VendorCategory{},

	)
	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
}
