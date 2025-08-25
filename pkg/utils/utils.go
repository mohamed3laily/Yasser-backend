package utils

import (
	"fmt"
	"strconv"

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

func GetRequiredUintFromHeader(c *gin.Context, headerName string) (uint, error) {
	headerValue := c.GetHeader(headerName)
	if headerValue == "" {
		return 0, fmt.Errorf("required header '%s' is missing", headerName)
	}

	id, err := strconv.ParseUint(headerValue, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid value for header '%s': expected a number", headerName)
	}

	return uint(id), nil
}

