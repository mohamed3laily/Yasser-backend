package bootstrap

import (
	"log"
	"yasser-backend/config"
	"yasser-backend/database"
	"yasser-backend/internal/search"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AppDependencies struct {
	Config       *config.Config
	DB           *gorm.DB
	SearchClient *search.Client
	Validator    *validator.Validate
}

func NewDependencies() *AppDependencies {
	cfg := config.Load()
	db := database.Init(cfg)
	searchClient := search.NewClient(cfg.MeiliHost, cfg.MeiliAPIKey, cfg.MeiliIndexName)
	validator := validator.New()

	log.Println("✅ Core dependencies initialized.")

	return &AppDependencies{
		Config:       cfg,
		DB:           db,
		SearchClient: searchClient,
		Validator:    validator,
	}
}
