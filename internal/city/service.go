package city

import (
	customerrors "yasser-backend/pkg/errors"
	"yasser-backend/pkg/pagination"
)


type Service interface {
	GetAllCities(search string, page, perPage int) ([]*City, *pagination.Result, error)
	GetDistricts(cityID *uint, search string, page, perPage int) ([]*District, *pagination.Result, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAllCities(search string, page, perPage int) ([]*City, *pagination.Result, error) {
	return s.repo.GetAllCities(search, page, perPage)
}

func (s *service) GetDistricts(cityID *uint, search string, page, perPage int) ([]*District, *pagination.Result, error) {

	if cityID != nil && *cityID == 0 {
		return nil, nil, customerrors.BadRequest("city.invalid_id")
	}

	return s.repo.GetDistricts(cityID, search, page, perPage)
}
