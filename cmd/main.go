package main

import (
	"log"
	"yasser-backend/bootstrap"
	"yasser-backend/internal/routes"
	"yasser-backend/migration"
	"yasser-backend/pkg/i18n"

	"github.com/gin-gonic/gin"
)

func main() {
	deps := bootstrap.NewDependencies()

	migration.Migrate(deps.DB)

	if err := i18n.Init("./locales"); err != nil {
		log.Fatal("Failed to initialize i18n:", err)
	}

	route := gin.Default()
	route.Use(i18n.LanguageMiddleware())

	registry := routes.NewRegistry(deps)
	registry.RegisterAllRoutes(route, deps)

	log.Printf("🚀 Starting server on port %s", deps.Config.Port)
	if err := route.Run(":" + deps.Config.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
