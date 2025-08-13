package city

import (
	"yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAllCities(c *gin.Context) ([]*City, *response.PaginationMeta, error)
	GetDistricts(c *gin.Context, cityID *uint) ([]*District, *response.PaginationMeta, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAllCities(c *gin.Context) ([]*City, *response.PaginationMeta, error) {
	return s.repo.GetAllCities(c)
}

func (s *service) GetDistricts(c *gin.Context, cityID *uint) ([]*District, *response.PaginationMeta, error) {
	if cityID != nil && *cityID == 0 {
		return nil, nil, errors.ErrInvalid
	}
	return s.repo.GetDistricts(c, cityID)
}
