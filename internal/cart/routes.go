package cart

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Routes struct {
	handler *Handler
	service Service
}



func SetupCartModule(db *gorm.DB, validator *validator.Validate) *Routes {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service, validator)
	
	return &Routes{
		handler: handler,
	}
}

func (r *Routes) GetService() Service {
    return r.service
}

func (r *Routes) RegisterRoutes(router *gin.RouterGroup) {
	cartGroup := router.Group("/cart")
	{
		cartGroup.POST("/validate", r.handler.ValidateCart)
	}
}