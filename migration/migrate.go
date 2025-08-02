package migration

import (
	"log"
	"yasser-backend/database"
	"yasser-backend/internal/user"
)

func Migrate() {
	err := database.DB.AutoMigrate(
		&user.User{},
	)
	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
}
