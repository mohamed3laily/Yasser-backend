package vendor

import (
	"errors"
	customerrors "yasser-backend/pkg/errors"
	"yasser-backend/pkg/pagination"

	"gorm.io/gorm"
)

type Repository interface {
	GetByID(id uint) (*Vendor, error)
	GetAll(filter VendorFilter, page, perPage int) ([]*Vendor, *pagination.Result, error)
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

func (r *repository) GetAll(filter VendorFilter, page, perPage int) ([]*Vendor, *pagination.Result, error) {
	var vendors []*Vendor

	baseQuery := r.db.Model(&Vendor{}).Scopes(
		r.filterByDistrict(filter.DistrictID),
		r.filterByCategory(filter.CategoryID),
		r.filterActive(),
	)

	paginationInfo, err := pagination.GetInfo(baseQuery, &Vendor{}, page, perPage)
	if err != nil {
		return nil, nil, err
	}

	dataQuery := baseQuery.Scopes(
		pagination.Paginate(page, perPage),
		r.preloadRelations(),
	)

	if err := dataQuery.Find(&vendors).Error; err != nil {
		return nil, nil, err
	}

	return vendors, paginationInfo, nil
}