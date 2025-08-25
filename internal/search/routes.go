package search

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Routes struct {
	handler *Handler
}

func NewRoutes(handler *Handler) *Routes {
	return &Routes{
		handler: handler,
	}
}


func SetupSearchModule(db *gorm.DB, client *Client, validator *validator.Validate) *Routes {
	repo := NewRepository(db)

	service := NewService(client, repo)

	handler := NewHandler(service, validator)

	return NewRoutes(handler)
}

func (r *Routes) RegisterRoutes(router *gin.RouterGroup) {
	searchGroup := router.Group("/search")
	{
		searchGroup.GET("", r.handler.Search)
	}
}
