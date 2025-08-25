package database

import (
	"math"
	"strconv"
	"yasser-backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaginationResult struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	LastPage    int   `json:"last_page"`
	HasNext     bool  `json:"has_next"`
	HasPrev     bool  `json:"has_prev"`
}

const (
	DefaultPage    = 1
	DefaultPerPage = 10
	MaxPerPage     = 100
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := utils.GetPageQuery(c)
		perPage := utils.GetPerPageQuery(c)
		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}

func GetPaginationInfo(c *gin.Context, db *gorm.DB, model interface{}) (*PaginationResult, error) {
    page := utils.GetPageQuery(c)
    perPage := utils.GetPerPageQuery(c)

    var total int64
    if err := db.Model(model).Count(&total).Error; err != nil {
        return nil, err
    }

    lastPage := int(math.Ceil(float64(total) / float64(perPage)))
    if lastPage == 0 {
        lastPage = 1
    }

    return &PaginationResult{
        CurrentPage: page,
        PerPage:     perPage,
        Total:       total,
        LastPage:    lastPage,
        HasNext:     page < lastPage,
        HasPrev:     page > 1,
    }, nil
}

func getPageFromContext(c *gin.Context) int {
	pageStr := c.DefaultQuery("page", strconv.Itoa(DefaultPage))
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		return DefaultPage
	}
	return page
}