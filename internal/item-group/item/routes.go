package item

import (
	"yasser-backend/database"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	handler *Handler
}

func SetupItemModule() *Routes {
	repo := NewRepository(database.DB)
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