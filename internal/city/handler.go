package city

import (
	"strconv"

	"yasser-backend/pkg/context"
	customerrors "yasser-backend/pkg/errors"
	"yasser-backend/pkg/locale"
	"yasser-backend/pkg/response"
	"yasser-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetCities(c *gin.Context) {
	searchQuery := c.Query("search")
	page := utils.GetPageQuery(c)
	perPage := utils.GetPerPageQuery(c)

	cities, paginationResult, err := h.service.GetAllCities(searchQuery, page, perPage)
	if err != nil {
		appErr := customerrors.Handle(c, err, "location.city.failed_to_get_all")
		response.Error(c, appErr)
		return
	}

	lang := context.GetLanguage(c)
	var citiesResponse []map[string]interface{}
	for _, city := range cities {
		citiesResponse = append(citiesResponse, map[string]interface{}{
			"id":   city.ID,
			"name": locale.ChooseLang(city.NameEn, city.NameAr, lang),
			"lat":  city.Latitude,
			"lng":  city.Longitude,
		})
	}

	paginationMeta := response.FromPaginationResult(paginationResult)
	paginatedResponse := response.NewPaginatedResponse(citiesResponse, paginationMeta)
	response.Success(c, "location.city.retrieved_successfully", paginatedResponse)
}

func (h *Handler) GetDistrictsByCity(c *gin.Context) {
	cidStr := c.Param("id")
	cid, err := strconv.ParseUint(cidStr, 10, 32)
	if err != nil {
		appErr := customerrors.BadRequest("common.invalid_id").WithContext(c)
		response.Error(c, appErr)
		return
	}
	cityID := uint(cid)

	searchQuery := c.Query("search")
	page := utils.GetPageQuery(c)
	perPage := utils.GetPerPageQuery(c)

	districts, paginationResult, err := h.service.GetDistricts(&cityID, searchQuery, page, perPage)
	if err != nil {
		appErr := customerrors.Handle(c, err, "location.district.failed_to_get_all")
		response.Error(c, appErr)
		return
	}

	lang := context.GetLanguage(c)
	var districtsResponse []map[string]interface{}
	for _, district := range districts {
		districtsResponse = append(districtsResponse, map[string]interface{}{
			"id":     district.ID,
			"name":   locale.ChooseLang(district.NameEn, district.NameAr, lang),
			"minLat": district.MinLat,
			"maxLat": district.MaxLat,
			"minLng": district.MinLng,
			"maxLng": district.MaxLng,
		})
	}

	paginationMeta := response.FromPaginationResult(paginationResult)
	paginatedResponse := response.NewPaginatedResponse(districtsResponse, paginationMeta)
	response.Success(c, "location.district.retrieved_successfully", paginatedResponse)
}
