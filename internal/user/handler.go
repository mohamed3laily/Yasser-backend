package user

import (
	"yasser-backend/pkg/common"
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

func (h *Handler) UpdateUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		appErr := errors.Unauthorized("user.not_authenticated").WithContext(c)
		response.Error(c, appErr)
		return
	}

	var req UpdateUserRequest
	if err := common.BindJSONStrict(c, &req); err != nil {
		response.ValidationError(c, err)
		return
	}

	if err := req.Validate(); err != nil {
		response.ValidationError(c, err)
		return
	}

	updatedUser, err := h.service.UpdateUser(userID.(uint), req)
	if err != nil {
		appErr := errors.Handle(c, err, "user.failed_to_update")
		response.Error(c, appErr)
		return
	}

	response.Success(c, "user.updated_successfully", updatedUser)
}

func (h *Handler) Ping(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		appErr := errors.Unauthorized("user.not_authenticated").WithContext(c)
		response.Error(c, appErr)
		return
	}

	if err := h.service.UpdateLastLogin(c.Request.Context(), userID.(uint)); err != nil {
		appErr := errors.Handle(c, err, "user.failed_to_update_last_login")
		response.Error(c, appErr)
		return
	}

	response.OK(c, "user.is_online")
}

func (h *Handler) GetMe(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		appErr := errors.Unauthorized("user.not_authenticated").WithContext(c)
		response.Error(c, appErr)
		return
	}

	user, err := h.service.GetUserByID(userID.(uint))
	if err != nil {
		appErr := errors.Handle(c, err, "user.failed_to_get")
		response.Error(c, appErr)
		return
	}

	response.Success(c, "user.retrieved_successfully", user)
}