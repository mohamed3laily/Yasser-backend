package city

import (
	"errors"
	"yasser-backend/pkg/database"
	customerrors "yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllCities(c *gin.Context) ([]*City, *response.PaginationMeta, error)
	GetDistricts(c *gin.Context, cityID *uint) ([]*District, *response.PaginationMeta, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAllCities(c *gin.Context) ([]*City, *response.PaginationMeta, error) {
	var cities []*City

	search := c.Query("search")

	db := r.db.Model(&City{}).Scopes(SearchCityByName(search))

	err := db.Scopes(database.Paginate(c)).Find(&cities).Error
	if err != nil {
		return nil, nil, err
	}

	paginationInfo, err := database.GetPaginationInfo(c, db, &City{})
	if err != nil {
		return nil, nil, err
	}

	meta := response.FromDatabasePagination(paginationInfo)
	return cities, meta, nil
}

func (r *repository) GetDistricts(c *gin.Context, cityID *uint) ([]*District, *response.PaginationMeta, error) {
	var districts []*District

	search := c.Query("search")
	db := r.db.Model(&District{})

	// Apply city filter if provided
	if cityID != nil {
		db = db.Scopes(DistrictByCity(*cityID))
	}

	db = db.Scopes(SearchDistrictByName(search))

	err := db.Scopes(database.Paginate(c)).Find(&districts).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, customerrors.ErrNotFound
		}
		return nil, nil, err
	}

	paginationInfo, err := database.GetPaginationInfo(c, db, &District{})
	if err != nil {
		return nil, nil, err
	}

	meta := response.FromDatabasePagination(paginationInfo)
	return districts, meta, nil
}