package main

import (
	"os"
	"yasser-backend/config"
	"yasser-backend/database"
	"yasser-backend/internal/routes"
	"yasser-backend/migration"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Init()
	migration.Migrate()

	route := gin.Default()
	registry := routes.NewRegistry()
	registry.RegisterAllRoutes(route)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	route.Run(":" + port)
}
