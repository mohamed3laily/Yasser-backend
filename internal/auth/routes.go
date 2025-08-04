package auth

import (
	"os"
	"yasser-backend/internal/user"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	handler *Handler
}

func NewRoutes(service *Service) *Routes {
	return &Routes{
		handler: NewHandler(service),
	}
}

func (r *Routes) login(c *gin.Context) {
	r.handler.Login(c)
}

func (r *Routes) verifyOtp(c *gin.Context) {
	r.handler.VerifyOtp(c)
}

func (r *Routes) RegisterRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", r.login)
		auth.POST("/verify-otp", r.verifyOtp)
	}
}

func SetupAuthModule() *Routes {
	authRepo := NewRepository()
	userRepo := user.NewRepository()
	waSender := NewWhatsAppSender(WhatsAppConfig{
		APIURL:      "https://app.wattsi.net/api/send",
		InstanceID:  os.Getenv("WATTSI_INSTANCE_ID"),
		AccessToken: os.Getenv("WATTSI_ACCESS_TOKEN"),
	})
	authService := NewService(authRepo, userRepo, waSender)

	return NewRoutes(authService)
}