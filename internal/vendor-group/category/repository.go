package category

import (
	"errors"
	customerrors "yasser-backend/pkg/errors"

	"gorm.io/gorm"
)

type Repository interface {
	GetByID(id uint) (*VendorCategory, error)
	GetAll() ([]*VendorCategory, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(category *VendorCategory) error {
	if err := r.db.Create(category).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return customerrors.ErrInvalid
		}
		return err
	}
	return nil
}

func (r *repository) GetByID(id uint) (*VendorCategory, error) {
	var category VendorCategory
	err := r.db.First(&category, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerrors.ErrNotFound
		}
		return nil, err
	}
	return &category, nil
}

func (r *repository) GetAll() ([]*VendorCategory, error) {
	var categories []*VendorCategory
	query := r.db.Order("name_en ASC")
	
	err := query.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}