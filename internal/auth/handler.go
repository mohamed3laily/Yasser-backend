package auth

import (
	"strings"

	"github.com/gin-gonic/gin"

	customerrors "yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"
)

type LoginRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,min=10,max=15"`
}

type VerifyOtpRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,min=10,max=15"`
	Otp         string `json:"otp" binding:"required,len=6,numeric"`
}

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	phoneNumber := strings.TrimSpace(req.PhoneNumber)

	err := h.service.Login(phoneNumber)
	if err != nil {
		appErr := customerrors.Handle(err, "Failed to send verification code")
		response.Error(c, appErr)
		return
	}

	// Generic response with any data structure
	response.Success(c, "Verification code sent successfully", map[string]interface{}{
		"phoneNumber": phoneNumber,
		"message":     "Please check your WhatsApp for the 6-digit verification code",
	})
}

func (h *Handler) VerifyOtp(c *gin.Context) {
	var req VerifyOtpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	phoneNumber := strings.TrimSpace(req.PhoneNumber)
	otp := strings.TrimSpace(req.Otp)

	userData, err := h.service.VerifyOtp(phoneNumber, otp)
	if err != nil {
		appErr := customerrors.Handle(err, "Failed to verify code")
		response.Error(c, appErr)
		return
	}

	response.Success(c, "Login successful", map[string]interface{}{
		"message": "Welcome! You have been successfully authenticated",
		"user":    userData,
		// "token": "jwt-token-here", // Add when implementing JWT
	})
}