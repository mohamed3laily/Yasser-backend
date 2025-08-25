package pagination

import (
	"math"

	"gorm.io/gorm"
)

type Result struct {
	CurrentPage int   `json:"currentPage"`
	PerPage     int   `json:"perPage"`
	Total       int64 `json:"total"`
	LastPage    int   `json:"lastPage"`
	HasNext     bool  `json:"hasNext"`
	HasPrev     bool  `json:"hasPrev"`
}


func Paginate(page, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}

func GetInfo(db *gorm.DB, model interface{}, page, perPage int) (*Result, error) {
	var total int64
	if err := db.Model(model).Count(&total).Error; err != nil {
		return nil, err
	}

	lastPage := int(math.Ceil(float64(total) / float64(perPage)))
	if lastPage == 0 {
		lastPage = 1
	}

	return &Result{
		CurrentPage: page,
		PerPage:     perPage,
		Total:       total,
		LastPage:    lastPage,
		HasNext:     page < lastPage,
		HasPrev:     page > 1,
	}, nil
}
