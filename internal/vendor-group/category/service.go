package category

import (
	customerrors "yasser-backend/pkg/errors"
)

type Service interface {
	GetCategoryByID(id uint) (*VendorCategory, error)
	GetAllCategories() ([]*VendorCategory, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetCategoryByID(id uint) (*VendorCategory, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalid
	}

	return s.repo.GetByID(id)
}

func (s *service) GetAllCategories() ([]*VendorCategory, error) {
	return s.repo.GetAll()
}