package user

import (
	"net/http"
	"yasser-backend/pkg/common"

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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req UpdateUserRequest
	if err := common.BindJSONStrict(c, &req); err != nil {
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := h.service.UpdateUser(userID.(uint), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (h *Handler) Ping(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := h.service.UpdateLastLogin(c.Request.Context(), userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update login"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "online"})
}