package vendor

import (
	"yasser-backend/internal/item-group/item"
	"yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetVendorByID(id uint , lang string) (*VendorResponse, error )
	GetAllVendors(c *gin.Context, filter VendorFilter) ([]*Vendor, *response.PaginationMeta, error)
}

type service struct {
	repo        Repository
	itemService item.Service
}

func NewService(repo Repository, itemService item.Service) Service {
	return &service{repo: repo, itemService: itemService}
}

func (s *service) GetVendorByID(id uint, lang string) (*VendorResponse, error) {
	if id == 0 {
		return nil, errors.ErrInvalid
	}

	vendor, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	resp := vendor.ToResponse(lang)

	itemsTree, err := s.itemService.GetVendorItemsGrouped(id, lang)
	if err == nil {
		resp.ItemsTree = itemsTree
	}

	return resp, nil
}

func (s *service) GetAllVendors(c *gin.Context, filter VendorFilter) ([]*Vendor, *response.PaginationMeta, error) {
	if filter.DistrictID == 0 {
		return nil, nil, errors.ErrInvalid
	}
	return s.repo.GetAll(c, filter)
}