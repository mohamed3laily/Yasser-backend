package search

import (
	"yasser-backend/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Routes struct {
	handler *Handler
}

func NewRoutes(service Service, validator *validator.Validate) *Routes {
	return &Routes{
		handler: NewHandler(service, validator),
	}
}

func SetupSearchModule(client *Client) *Routes {
	repo := NewRepository(database.DB)
	service := NewService(client, repo)

	validator := validator.New()

	return &Routes{
		handler: NewHandler(service, validator),
	}
}

func (r *Routes) RegisterRoutes(router *gin.RouterGroup) {
	searchGroup := router.Group("/search")
	{
		searchGroup.GET("", r.handler.Search) // /api/search?query=shawarma&type=item
	}
}
