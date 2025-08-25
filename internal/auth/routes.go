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
	Service *Service
}

func NewRoutes(service *Service , validator *validator.Validate) *Routes {
	return &Routes{
		handler: NewHandler(service, validator),
	}
}

func (r *Routes) login(c *gin.Context) {
	r.handler.Login(c)
}

func (r *Routes) verifyOtp(c *gin.Context) {
	r.handler.VerifyOtp(c)
}

func (r *Routes) RegisterRoutes(router *gin.RouterGroup, ) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", r.handler.Login)
		auth.POST("/verify-otp", r.handler.VerifyOtp)
	}
}

func SetupAuthModule(db *gorm.DB, validator *validator.Validate, cfg *config.Config) *Routes {
	authRepo := NewRepository(db)
	userRepo := user.NewRepository(db)

	// The WhatsAppSender is now configured via the central config struct.
	waSender := NewWhatsAppSender(WhatsAppConfig{
		APIURL:      "https://app.wattsi.net/api/send", // This could also be in config
		InstanceID:  cfg.WattsiInstanceID,             // Add this to your config struct
		AccessToken: cfg.WattsiAccessToken,            // Add this to your config struct
	} )

	// The JWT secret is also passed in from the central config.
	authService := NewService(authRepo, userRepo, waSender, cfg.JWTSecret)
	handler := NewHandler(authService, validator)

	return &Routes{
		handler: handler,
		Service: authService,
	}
}