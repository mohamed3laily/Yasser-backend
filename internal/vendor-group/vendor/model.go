package vendor

import (
	"yasser-backend/internal/city"
	"yasser-backend/internal/vendor-group/category"
	"yasser-backend/pkg/models"
)

type Vendor struct {
	models.BaseModel

	// Information
	NameEn        string `json:"nameEn" gorm:"size:100;not null"`
	NameAr        string `json:"nameAr" gorm:"size:100;not null"`
	DescriptionEn string `json:"descriptionEn" gorm:"size:255"`
	DescriptionAr string `json:"descriptionAr" gorm:"size:255"`
	ProfilePicture string `json:"profilePicture" gorm:"size:255"`
	CoverPicture   string `json:"coverPicture" gorm:"size:255"`
	Phone          string `json:"phone" gorm:"size:20"`
	Email          string `json:"email" gorm:"size:100"`

	// Location
	CityID uint       `json:"cityId"`
	City   city.City  `gorm:"foreignKey:CityID"`

	AreaID uint            `json:"areaId"`
	Area   city.District   `gorm:"foreignKey:AreaID"`

	AddressEn string `json:"addressEn" gorm:"size:255"`
	AddressAr string `json:"addressAr" gorm:"size:255"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	// Opening and Closing Time
	OpeningTime string `json:"openingTime" gorm:"size:255"`
	ClosingTime string `json:"closingTime" gorm:"size:255"`

	// Category
	CategoryID uint                     `json:"categoryId"`
	Category   category.VendorCategory `gorm:"foreignKey:CategoryID"`

	IsActive bool `json:"isActive" gorm:"default:true"`
}

type VendorResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ProfilePicture string `json:"profilePicture"`
	CoverPicture   string `json:"coverPicture"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`

	// Location
	CityID   uint   `json:"cityId"`
	CityName string `json:"cityName"`

	AreaID   uint   `json:"areaId"`
	AreaName string `json:"areaName"`

	Address string `json:"address"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	// Opening and Closing Time
	OpeningTime string `json:"openingTime"`
	ClosingTime string `json:"closingTime"`

	// Category
	CategoryID   uint   `json:"categoryId"`
	CategoryName string `json:"categoryName"`

	IsActive bool `json:"isActive"`
}

func (v *Vendor) ToResponse(lang string) *VendorResponse {
	response := &VendorResponse{
		ID:             v.ID,
		ProfilePicture: v.ProfilePicture,
		CoverPicture:   v.CoverPicture,
		Phone:          v.Phone,
		Email:          v.Email,
		CityID:         v.CityID,
		AreaID:         v.AreaID,
		Latitude:       v.Latitude,
		Longitude:      v.Longitude,
		OpeningTime:    v.OpeningTime,
		ClosingTime:    v.ClosingTime,
		CategoryID:     v.CategoryID,
		IsActive:       v.IsActive,
	}

	if lang == "ar" {
		response.Name = v.NameAr
		response.Description = v.DescriptionAr
		response.Address = v.AddressAr
		if v.City.NameAr != "" {
			response.CityName = v.City.NameAr
		} else {
			response.CityName = v.City.NameEn
		}
		if v.Area.NameAr != "" {
			response.AreaName = v.Area.NameAr
		} else {
			response.AreaName = v.Area.NameEn
		}
		response.CategoryName = v.Category.NameAr
	} else {
		response.Name = v.NameEn
		response.Description = v.DescriptionEn
		response.Address = v.AddressEn
		if v.City.NameEn != "" {
			response.CityName = v.City.NameEn
		} else {
			response.CityName = v.City.NameAr
		}
		if v.Area.NameEn != "" {
			response.AreaName = v.Area.NameEn
		} else {
			response.AreaName = v.Area.NameAr
		}
		response.CategoryName = v.Category.NameEn
	}

	return response
}