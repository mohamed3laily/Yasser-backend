package vendor

import (
	"strconv"
	"yasser-backend/pkg/context"
	customerrors "yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"
	"yasser-backend/pkg/utils"

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
		appErr := customerrors.BadRequest("common.invalid_id").WithContext(c)
		response.Error(c, appErr)
		return
	}

	vendor, err := h.service.GetVendorByID(uint(id))
	if err != nil {
		appErr := customerrors.Handle(c, err, "vendor-group.vendor.failed_to_get")
		response.Error(c, appErr)
		return
	}

	lang := context.GetLanguage(c)
	vendorResponse := vendor.ToResponse(lang)

	response.Success(c, "vendor-group.vendor.retrieved_successfully", vendorResponse)
}

func (h *Handler) GetAllVendors(c *gin.Context) {
	DistrictID, ok := utils.GetDistrictIDFromHeader(c)
	if !ok {
		return
	}

	filter := VendorFilter{
		DistrictID:     uint(DistrictID),
		CategoryID: utils.GetOptionalUintQuery(c, "category_id"),
	}

	vendors, meta, err := h.service.GetAllVendors(c, filter)
	if err != nil {
		appErr := customerrors.Handle(c, err, "vendor-group.vendor.failed_to_get_all")
		response.Error(c, appErr)
		return
	}

	lang := context.GetLanguage(c)
	vendorsResponse := make([]*VendorResponse, 0, len(vendors))
	for _, v := range vendors {
		vendorsResponse = append(vendorsResponse, v.ToResponse(lang))
	}

	paginatedResponse := response.NewPaginatedResponse(vendorsResponse, meta)
	response.Success(c, "vendor-group.vendor.retrieved_successfully", paginatedResponse)
}