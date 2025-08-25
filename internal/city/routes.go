package city

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	handler *Handler
}

func NewRoutes(service Service) *Routes {
	return &Routes{
		handler: NewHandler(service),
	}
}

func SetupCityModule(db *gorm.DB) *Routes {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	return &Routes{handler: handler}
}

func (r *Routes) RegisterRoutes(router *gin.RouterGroup) {
	cityGroup := router.Group("/cities")
	{
		cityGroup.GET("", r.handler.GetCities)
		cityGroup.GET("/:id/districts", r.handler.GetDistrictsByCity)
	}
}
