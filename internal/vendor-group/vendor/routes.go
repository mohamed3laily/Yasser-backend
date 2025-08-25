package vendor

import (
	"yasser-backend/internal/item-group/item"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Routes struct {
	handler *Handler
}

func NewRoutes(service Service, validator *validator.Validate) *Routes {
	return &Routes{
		handler: NewHandler(service, validator),
	}
}

func SetupVendorModule(db *gorm.DB, validator *validator.Validate) *Routes {
	repo := NewRepository(db)
	
	itemRepo := item.NewRepository(db)
	itemService := item.NewService(itemRepo)

	service := NewService(repo, itemService)
	handler := NewHandler(service, validator)

	return &Routes{
		handler: handler,
	}
}

func (r *Routes) RegisterRoutes(router *gin.RouterGroup) {
	vendorGroup := router.Group("/vendors")
	{
		vendorGroup.GET("", r.handler.GetAllVendors)
		vendorGroup.GET("/:id", r.handler.GetVendor)
	}
}