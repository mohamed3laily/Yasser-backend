package routes

import (
	"yasser-backend/bootstrap"
	"yasser-backend/internal/auth"
	"yasser-backend/internal/city"
	"yasser-backend/internal/item-group/item"
	"yasser-backend/internal/search"
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
	itemRoutes *item.Routes
	searchRoutes *search.Routes
}

func NewRegistry(deps *bootstrap.AppDependencies) *Registry {
	return &Registry{

		authRoutes:     auth.SetupAuthModule(deps.DB, deps.Validator, deps.Config),
		userRoutes:     user.SetupUserModule(deps.DB, deps.Validator),
		categoryRoutes: category.SetupCategoryModule(deps.DB, deps.Validator),
		vendorRoutes:   vendor.SetupVendorModule(deps.DB, deps.Validator),
		cityRoutes:     city.SetupCityModule(deps.DB),
		itemRoutes:     item.SetupItemModule(deps.DB, deps.Validator),
		searchRoutes:   search.SetupSearchModule(deps.DB, deps.SearchClient, deps.Validator),
	}
}

func (r *Registry) RegisterAllRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	jwtMiddleware := auth.JWTAuthMiddleware(&user.Service{})

	r.authRoutes.RegisterRoutes(v1)
	r.userRoutes.RegisterRoutes(v1, jwtMiddleware)
	r.categoryRoutes.RegisterRoutes(v1)
	r.vendorRoutes.RegisterRoutes(v1)
	r.cityRoutes.RegisterRoutes(v1)
	r.itemRoutes.RegisterRoutes(v1)
	r.searchRoutes.RegisterRoutes(v1)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})
}