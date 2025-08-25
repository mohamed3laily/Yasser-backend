package item

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Routes struct {
	handler *Handler
}

func SetupItemModule(db *gorm.DB, validator *validator.Validate) *Routes {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	return &Routes{
		handler: handler,
	}
}


func (r *Routes) RegisterRoutes(router *gin.RouterGroup) {
	itemGroup := router.Group("/vendors/:id/items")
	{
		itemGroup.GET("/:itemId", r.handler.GetItem)
	}
}