package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	DefaultPage    = 1
	DefaultPerPage = 10
	MaxPerPage     = 100
)

func GetPageQuery(c *gin.Context) int {
	pageStr := c.DefaultQuery("page", strconv.Itoa(DefaultPage))
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		return DefaultPage
	}
	return page
}

func GetPerPageQuery(c *gin.Context) int {
	perPageStr := c.DefaultQuery("per_page", strconv.Itoa(DefaultPerPage))
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage <= 0 {
		perPage = DefaultPerPage
	} else if perPage > MaxPerPage {
		perPage = MaxPerPage
	}
	return perPage
}

func GetIntQuery(c *gin.Context, key string, defaultValue int) int {
	valStr := c.Query(key)
	if valStr == "" {
		return defaultValue
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultValue
	}
	return val
}
