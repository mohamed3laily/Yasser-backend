package main

import (
	"log"
	"os"
	"yasser-backend/config"
	"yasser-backend/database"
	"yasser-backend/internal/routes"
	"yasser-backend/migration"
	"yasser-backend/pkg/i18n"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Init()
	migration.Migrate()
		if err := i18n.Init("./locales"); err != nil {
		log.Fatal("Failed to initialize i18n:", err)
	}


	route := gin.Default()

	route.Use(i18n.LanguageMiddleware())

	registry := routes.NewRegistry()
	registry.RegisterAllRoutes(route)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	route.Run(":" + port)
}
