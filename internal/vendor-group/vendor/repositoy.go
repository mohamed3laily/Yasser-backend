// vendor/repository.go
package vendor

import (
	"errors"
	"yasser-backend/pkg/dto"
	customerrors "yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"

	"gorm.io/gorm"
)

type Repository interface {
	GetByID(id uint) (*Vendor, error)
	GetAll(pagination *dto.PaginationQuery) ([]*Vendor, *response.PaginationMeta, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetByID(id uint) (*Vendor, error) {
	var vendor Vendor
	err := r.db.Preload("City").Preload("Area").Preload("Category").First(&vendor, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrNotFound
		}
		return nil, err
	}
	return &vendor, nil
}

func (r *repository) GetAll(pagination *dto.PaginationQuery) ([]*Vendor, *response.PaginationMeta, error) {
	var vendors []*Vendor
	var total int64

	if err := r.db.Model(&Vendor{}).Count(&total).Error; err != nil {
		return nil, nil, err
	}

	err := r.db.Preload("City").Preload("Area").Preload("Category").
		Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Find(&vendors).Error
	if err != nil {
		return nil, nil, err
	}

	meta := &response.PaginationMeta{
		CurrentPage: pagination.Page,
		PerPage:     pagination.PerPage,
		Total:       int(total),
		LastPage:    pagination.CalculateLastPage(int(total)),
	}

	return vendors, meta, nil
}