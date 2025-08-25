package search

import (
	"strconv"
	"yasser-backend/pkg/context"
	customerrors "yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service   Service
	validator *validator.Validate
}

// NewHandler creates a new search handler with its dependencies.
func NewHandler(service Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

func (h *Handler) Search(c *gin.Context) {
	var req SearchRequest
	
	// Get query parameters
	req.Query = c.Query("query")
	req.Lang = context.GetLanguage(c)
	req.Type = c.Query("type")
	
	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			req.Limit = limit
		}
	}

	// Validate request (optional, but good practice)
	if err := h.validator.Struct(req); err != nil {
		appErr := customerrors.BadRequest("search.invalid_request").WithContext(c)
		response.Error(c, appErr)
		return
	}

	// Perform search
	results, err := h.service.Search(req)
	if err != nil {
		appErr := customerrors.Handle(c, err, "search.failed")
		response.Error(c, appErr)
		return
	}

	response.Success(c, "search.success", map[string]interface{}{
		"results": results,
		"query":   req.Query,
		"count":   len(results),
	})
}
