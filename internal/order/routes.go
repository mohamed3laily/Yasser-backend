package order

// import (
// 	"yasser-backend/internal/cart"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"
// 	"gorm.io/gorm"
// )

// type Routes struct {
// 	handler *Handler
// }

// func SetupOrderModule(db *gorm.DB, validator *validator.Validate, cartService cart.Service) *Routes {
// 	repo := NewRepository(db)
// 	service := NewService(repo, cartService)
// 	handler := NewHandler(service, validator)

// 	return &Routes{
// 		handler: handler,
// 	}
// }

// func (r *Routes) RegisterRoutes(router *gin.RouterGroup) {
// 	orderGroup := router.Group("/orders")
// 	{
// 		orderGroup.POST("", r.handler.CreateOrder)
// 	}
// }
