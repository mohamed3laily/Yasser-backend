package category

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Routes struct {
	handler *Handler
}

func NewRoutes(service Service, validator *validator.Validate) *Routes {
	return &Routes{
		handler: NewHandler(service , validator),
	}
}

func SetupCategoryModule(db *gorm.DB, validator *validator.Validate) *Routes {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service, validator)

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