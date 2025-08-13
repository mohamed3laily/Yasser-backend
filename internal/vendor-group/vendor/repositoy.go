package vendor

import (
	"errors"
	"yasser-backend/pkg/database"
	customerrors "yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	GetByID(id uint) (*Vendor, error)
	GetAll(c *gin.Context, filter VendorFilter) ([]*Vendor, *response.PaginationMeta, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetByID(id uint) (*Vendor, error) {
	var vendor Vendor
	err := r.db.Preload("City").Preload("District").Preload("Category").First(&vendor, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrNotFound
		}
		return nil, err
	}
	return &vendor, nil
}

func (r *repository) GetAll(c *gin.Context, filter VendorFilter) ([]*Vendor, *response.PaginationMeta, error) {
	var vendors []*Vendor

	err := r.db.Scopes(
		database.Paginate(c),
		r.filterByCity(filter.CityID),
		r.filterByCategory(filter.CategoryID),
		r.filterActive(),
		r.preloadRelations(),
	).Find(&vendors).Error

	if err != nil {
		return nil, nil, err
	}

	paginationInfo, err := database.GetPaginationInfo(c, r.db, &Vendor{})
	if err != nil {
		return nil, nil, err
	}

	meta := response.FromDatabasePagination(paginationInfo)
	return vendors, meta, nil
}