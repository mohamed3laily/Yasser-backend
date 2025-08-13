package city

import (
	"strconv"

	"yasser-backend/pkg/context"
	"yasser-backend/pkg/errors"
	"yasser-backend/pkg/locale"
	"yasser-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetCities(c *gin.Context) {
	cities, meta, err := h.service.GetAllCities(c)
	if err != nil {
		appErr := errors.Handle(c, err, "location.city.failed_to_get_all")
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

	paginated := response.NewPaginatedResponse(citiesResponse, meta)
	response.Success(c, "location.city.retrieved_successfully", paginated)
}

func (h *Handler) GetDistrictsByCity(c *gin.Context) {
	cidStr := c.Param("id")
	cid, err := strconv.ParseUint(cidStr, 10, 32)
	if err != nil {
		appErr := errors.BadRequest("common.invalid_id").WithContext(c)
		response.Error(c, appErr)
		return
	}
	cityID := uint(cid)

	districts, meta, err := h.service.GetDistricts(c, &cityID)
	if err != nil {
		appErr := errors.Handle(c, err, "location.district.failed_to_get_all")
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

	paginated := response.NewPaginatedResponse(districtsResponse, meta)
	response.Success(c, "location.district.retrieved_successfully", paginated)
}