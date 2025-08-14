package utils

import (
	"strconv"
	"yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func GetOptionalUintQuery(c *gin.Context, key string) *uint {
	if val := c.Query(key); val != "" {
		if id, err := strconv.ParseUint(val, 10, 32); err == nil {
			uid := uint(id)
			return &uid
		}
	}
	return nil
}

func GetDistrictIDFromHeader(c *gin.Context) (uint, bool) {
	cityIDStr := c.GetHeader("X-District-ID")
	if cityIDStr == "" {
		appErr := errors.BadRequest("common.district_id_required").WithContext(c)
		response.Error(c, appErr)
		return 0, false
	}

	cityID, err := strconv.ParseUint(cityIDStr, 10, 64)
	if err != nil {
		appErr := errors.BadRequest("common.invalid_id").WithContext(c)
		response.Error(c, appErr)
		return 0, false
	}

	return uint(cityID), true
}
