package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Routes struct {
	handler *Handler
}

func NewRoutes(service *Service) *Routes {
	return &Routes{
		handler: NewHandler(*service),
	}
}

func (r *Routes) updateUser(c *gin.Context) {
	r.handler.UpdateUser(c)
}

func (r *Routes) RegisterRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	userGroup := router.Group("/users")
	userGroup.Use(authMiddleware)
	{
		userGroup.PUT("/me", r.updateUser)
		userGroup.GET("/ping", r.handler.Ping)
		userGroup.GET("/me", r.handler.GetMe)
	}
}

func SetupUserModule(db *gorm.DB, validator *validator.Validate) *Routes {
	userRepo := NewRepository(db)
	userService := NewService(userRepo)

	return NewRoutes(userService)
}
