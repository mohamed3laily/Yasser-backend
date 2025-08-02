package routes

import (
	"yasser-backend/internal/auth"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	authRoutes *auth.Routes
}

func NewRegistry() *Registry {
	return &Registry{
		authRoutes: auth.SetupAuthModule(),
	}
}

func (r *Registry) RegisterAllRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	r.authRoutes.RegisterRoutes(v1)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"version": "1.0.0",
		})
	})

	v1.GET("/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"api":     "Yasser Backend",
			"version": "1.0.0",
			"endpoints": gin.H{
				"auth": []string{
					"POST /api/v1/auth/login",
					"POST /api/v1/auth/verify-otp", 
					"POST /api/v1/auth/resend-otp",
				},
				"users": []string{
					"GET /api/v1/users/:id",
					"GET /api/v1/users/phone/:phone",
					"PUT /api/v1/users/:id",
					"DELETE /api/v1/users/:id",
				},
			},
		})
	})
}