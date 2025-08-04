package user

import "github.com/gin-gonic/gin"

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
		userGroup.PUT("/:id", r.updateUser)
		userGroup.GET("/ping", r.handler.Ping)
		userGroup.GET("/me", r.handler.GetMe)
	}
}

func SetupUserModule() *Routes {
	userRepo := NewRepository()
	userService := NewService(userRepo)

	return NewRoutes(userService)
}
