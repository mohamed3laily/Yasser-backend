package vendor

import (
	"yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetVendorByID(id uint) (*Vendor, error)
	GetAllVendors(c *gin.Context, filter VendorFilter) ([]*Vendor, *response.PaginationMeta, error)
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

func (s *service) GetAllVendors(c *gin.Context, filter VendorFilter) ([]*Vendor, *response.PaginationMeta, error) {
	if filter.CityID == 0 {
		return nil, nil, errors.ErrInvalid
	}
	return s.repo.GetAll(c, filter)
}