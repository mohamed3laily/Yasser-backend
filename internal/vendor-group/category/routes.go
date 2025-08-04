package category

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

func SetupCategoryModule() *Routes {
	repo := NewRepository(database.DB)
	service := NewService(repo)
	handler := NewHandler(service)

	return &Routes{
		handler: handler,
	}
}

func (r *Routes) RegisterRoutes(router *gin.RouterGroup) {
	categoryGroup := router.Group("/categories")
	{
		categoryGroup.GET("", r.handler.GetAllCategories)
		categoryGroup.GET("/:id", r.handler.GetCategory)
	}
}