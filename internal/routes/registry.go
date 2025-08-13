package routes

import (
	"yasser-backend/internal/auth"
	"yasser-backend/internal/city"
	"yasser-backend/internal/user"
	"yasser-backend/internal/vendor-group/category"
	"yasser-backend/internal/vendor-group/vendor"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	authRoutes *auth.Routes
	userRoutes *user.Routes
	categoryRoutes *category.Routes
	vendorRoutes *vendor.Routes
	cityRoutes *city.Routes
}

func NewRegistry() *Registry {
	return &Registry{
		authRoutes: auth.SetupAuthModule(),
		userRoutes: user.SetupUserModule(),
		categoryRoutes: category.SetupCategoryModule(),
		vendorRoutes: vendor.SetupVendorModule(),
		cityRoutes: city.SetupCityModule(),
	}
}

func (r *Registry) RegisterAllRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	r.authRoutes.RegisterRoutes(v1)
	r.userRoutes.RegisterRoutes(v1, auth.JWTAuthMiddleware())
	r.categoryRoutes.RegisterRoutes(v1)
	r.vendorRoutes.RegisterRoutes(v1)
	r.cityRoutes.RegisterRoutes(v1)


	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"version": "1.0.0",
		})
	})
}