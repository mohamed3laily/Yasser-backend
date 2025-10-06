package auth

import (
	"yasser-backend/config"
	"yasser-backend/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Routes struct {
	handler *Handler
}

func NewRoutes(service *Service, validator *validator.Validate) *Routes {
	return &Routes{
		handler: NewHandler(service, validator),
	}
}

func (r *Routes) RegisterRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", r.handler.Login)
		auth.POST("/verify-otp", r.handler.VerifyOtp)
	}
}

func SetupAuthModule(db *gorm.DB, validator *validator.Validate, cfg *config.Config) *Routes {
	authRepo := NewRepository(db)
	userRepo := user.NewRepository(db)

	waSender := NewWhatsAppSender(WhatsAppConfig{
		APIURL:      "https://app.wattsi.net/api/send",
		InstanceID:  cfg.WattsiInstanceID,
		AccessToken: cfg.WattsiAccessToken,
	})

	authService := NewService(authRepo, userRepo, waSender, cfg.JWTSecret)

	return &Routes{
		handler: NewHandler(authService, validator),
	}
}