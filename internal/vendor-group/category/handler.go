package category

import (
	"strconv"
	"yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		appErr := errors.BadRequest("common.invalid_id").WithContext(c)
		response.Error(c, appErr)
		return
	}

	category, err := h.service.GetCategoryByID(uint(id))
	if err != nil {
		appErr := errors.Handle(c, err, "vendor-group.category.failed_to_get")
		response.Error(c, appErr)
		return
	}

	lang := getLanguageFromContext(c)
	categoryResponse := category.ToResponse(lang)

	response.Success(c, "vendor-group.category.retrieved_successfully", categoryResponse)
}

func (h *Handler) GetAllCategories(c *gin.Context) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		appErr := errors.Handle(c, err, "vendor-group.category.failed_to_get_all")
		response.Error(c, appErr)
		return
	}

	lang := getLanguageFromContext(c)
	
	var categoriesResponse []*VendorCategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, category.ToResponse(lang))
	}

	response.Success(c, "vendor-group.category.retrieved_successfully", categoriesResponse)
}

func getLanguageFromContext(c *gin.Context) string {
	if lang, exists := c.Get("lang"); exists {
		return lang.(string)
	}
	return "en"
}