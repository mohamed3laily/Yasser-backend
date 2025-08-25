package search

import (
	"yasser-backend/pkg/context"
	customerrors "yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"
	"yasser-backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service   Service
	validator *validator.Validate
}

func NewHandler(service Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

func (h *Handler) Search(c *gin.Context) {
	var req SearchRequest

	req.Query = c.Query("query")
	req.Lang = context.GetLanguage(c)
	
    districtID, err := utils.GetRequiredUintFromHeader(c, "X-District-ID")
    if err != nil {

        appErr := customerrors.BadRequest("vendor.district_required")
        response.Error(c, appErr)
        return
    }

	req.DistrictID = districtID

	if typeVal := c.Query("type"); typeVal != "" {
		req.Type = &typeVal
	}

	req.Limit = utils.GetIntQuery(c, "limit", 20)
	req.Offset = utils.GetIntQuery(c, "offset", 0)

	if err := h.validator.Struct(req); err != nil {
		appErr := customerrors.BadRequest("common.request_failed")
		response.Error(c, appErr)
		return
	}

	results, err := h.service.Search(req)
	if err != nil {
		appErr := customerrors.Handle(c, err, "common.request_failed")
		response.Error(c, appErr)
		return
	}

	response.Success(c, "request success", map[string]interface{}{
		"results": results,
		"query":   req.Query,
		"count":   len(results),
		"limit":   req.Limit,
		"offset":  req.Offset,
	})
}
