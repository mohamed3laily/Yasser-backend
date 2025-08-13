package vendor

import (
	"yasser-backend/database"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	handler *Handler
}

func NewRoutes(service Service) *Routes {
	return &Routes{
		handler: NewHandler(service),
	}
}

func SetupVendorModule() *Routes {
	repo := NewRepository(database.DB)
	service := NewService(repo)
	handler := NewHandler(service)

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