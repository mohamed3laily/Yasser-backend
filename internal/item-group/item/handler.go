package item

import (
	"strconv"
	"yasser-backend/pkg/context"
	customerrors "yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetItem(c *gin.Context) {
	idParam := c.Param("itemId")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		appErr := customerrors.BadRequest("common.invalid_id").WithContext(c)
		response.Error(c, appErr)
		return
	}

	lang := context.GetLanguage(c)
	item, err := h.service.GetItemByID(uint(id), lang)
	if err != nil {
		appErr := customerrors.Handle(c, err, "item.failed_to_get")
		response.Error(c, appErr)
		return
	}

	response.Success(c, "item.retrieved_successfully", item)
}
