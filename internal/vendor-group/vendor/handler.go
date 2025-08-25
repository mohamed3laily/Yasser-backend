package vendor

import (
	"strconv"
	"yasser-backend/pkg/context"
	customerrors "yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"
	"yasser-backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service Service
}

func NewHandler(service Service, validator *validator.Validate) *Handler {
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

	lang := context.GetLanguage(c)
	vendor, err := h.service.GetVendorByID(uint(id), lang)
	if err != nil {
		appErr := customerrors.Handle(c, err, "vendor-group.vendor.failed_to_get")
		response.Error(c, appErr)
		return
	}

	response.Success(c, "vendor-group.vendor.retrieved_successfully", vendor)
}
func (h *Handler) GetAllVendors(c *gin.Context) {
    districtID, err := utils.GetRequiredUintFromHeader(c, "X-District-ID")
    if err != nil {
        appErr := customerrors.BadRequest("vendor.district_required")
        response.Error(c, appErr)
        return
    }

    page := utils.GetPageQuery(c)
    perPage := utils.GetPerPageQuery(c)

    filter := VendorFilter{
        DistrictID: districtID,
        CategoryID: utils.GetOptionalUintQuery(c, "category_id"),
    }

    vendors, paginationResult, err := h.service.GetAllVendors(filter, page, perPage)
    if err != nil {
        return
    }

    lang := context.GetLanguage(c)
    vendorsResponse := make([]*VendorResponse, 0, len(vendors))
    for _, v := range vendors {
        vendorsResponse = append(vendorsResponse, v.ToResponse(lang))
    }

    paginationMeta := response.FromPaginationResult(paginationResult)
    paginatedResponse := response.NewPaginatedResponse(vendorsResponse, paginationMeta)
    response.Success(c, "vendor.retrieved_successfully", paginatedResponse)
}