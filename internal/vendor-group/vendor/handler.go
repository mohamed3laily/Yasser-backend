package vendor

import (
	"strconv"

	"yasser-backend/pkg/context"
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

func (h *Handler) GetVendor(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		appErr := errors.BadRequest("common.invalid_id").WithContext(c)
		response.Error(c, appErr)
		return
	}

	vendor, err := h.service.GetVendorByID(uint(id))
	if err != nil {
		appErr := errors.Handle(c, err, "vendor.failed_to_get")
		response.Error(c, appErr)
		return
	}

	lang := context.GetLanguage(c)
	vendorResponse := vendor.ToResponse(lang)

	response.Success(c, "vendor.retrieved_successfully", vendorResponse)
}

func (h *Handler) GetAllVendors(c *gin.Context) {
	vendors, meta, err := h.service.GetAllVendors(c)
	if err != nil {
		appErr := errors.Handle(c, err, "vendor.failed_to_get_all")
		response.Error(c, appErr)
		return
	}

	lang := context.GetLanguage(c)
	
	var vendorsResponse []*VendorResponse
	for _, vendor := range vendors {
		vendorsResponse = append(vendorsResponse, vendor.ToResponse(lang))
	}

	paginatedResponse := response.NewPaginatedResponse(vendorsResponse, meta)
	response.Success(c, "vendor.retrieved_successfully", paginatedResponse)
}