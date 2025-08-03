package routes

import (
	"yasser-backend/internal/auth"
	"yasser-backend/internal/user"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	authRoutes *auth.Routes
	userRoutes *user.Routes
}

func NewRegistry() *Registry {
	return &Registry{
		authRoutes: auth.SetupAuthModule(),
		userRoutes: user.SetupUserModule(),
	}
}

func (r *Registry) RegisterAllRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	r.authRoutes.RegisterRoutes(v1)
	r.userRoutes.RegisterRoutes(v1, auth.JWTAuthMiddleware())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"version": "1.0.0",
		})
	})
}