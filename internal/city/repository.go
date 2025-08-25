package city

import (
	"errors"
	"yasser-backend/pkg/pagination"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllCities(search string, page, perPage int) ([]*City, *pagination.Result, error)
	GetDistricts(cityID *uint, search string, page, perPage int) ([]*District, *pagination.Result, error)
}


type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAllCities(search string, page, perPage int) ([]*City, *pagination.Result, error) {
	var cities []*City

	db := r.db.Model(&City{}).Scopes(SearchCityByName(search))


	paginationInfo, err := pagination.GetInfo(db, &City{}, page, perPage)
	if err != nil {
		return nil, nil, err
	}

	err = db.Scopes(pagination.Paginate(page, perPage)).Find(&cities).Error
	if err != nil {
		return nil, nil, err
	}

	return cities, paginationInfo, nil
}

func (r *repository) GetDistricts(cityID *uint, search string, page, perPage int) ([]*District, *pagination.Result, error) {
	var districts []*District

	db := r.db.Model(&District{})

	// Apply filters.
	if cityID != nil {
		db = db.Scopes(DistrictByCity(*cityID))
	}
	db = db.Scopes(SearchDistrictByName(search))

	paginationInfo, err := pagination.GetInfo(db, &District{}, page, perPage)
	if err != nil {
		return nil, nil, err
	}

	err = db.Scopes(pagination.Paginate(page, perPage)).Find(&districts).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*District{}, paginationInfo, nil
		}
		return nil, nil, err
	}

	return districts, paginationInfo, nil
}