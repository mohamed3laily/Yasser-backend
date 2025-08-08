package vendor

import (
	"yasser-backend/pkg/dto"
	"yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"
)

type Service interface {
	GetVendorByID(id uint) (*Vendor, error)
	GetAllVendors(pagination *dto.PaginationQuery) ([]*Vendor, *response.PaginationMeta, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetVendorByID(id uint) (*Vendor, error) {
	if id == 0 {
		return nil, errors.ErrInvalid
	}
	return s.repo.GetByID(id)
}

func (s *service) GetAllVendors(pagination *dto.PaginationQuery) ([]*Vendor, *response.PaginationMeta, error) {
	return s.repo.GetAll(pagination)
}