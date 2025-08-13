package utils

import (
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
