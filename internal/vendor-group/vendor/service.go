package vendor

import (
	"yasser-backend/internal/item-group/item"
	customerrors "yasser-backend/pkg/errors"
	"yasser-backend/pkg/pagination"
)


type Service interface {
	GetVendorByID(id uint, lang string) (*VendorResponse, error)
	GetAllVendors(filter VendorFilter, page, perPage int) ([]*Vendor, *pagination.Result, error)
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
		return nil, customerrors.BadRequest("vendor.invalid_id")
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

func (s *service) GetAllVendors(filter VendorFilter, page, perPage int) ([]*Vendor, *pagination.Result, error) {
	if filter.DistrictID == 0 {
		return nil, nil, customerrors.BadRequest("vendor.district_required")
	}
	return s.repo.GetAll(filter, page, perPage)
}
