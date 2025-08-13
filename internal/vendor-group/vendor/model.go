package vendor

import (
	"yasser-backend/internal/city"
	"yasser-backend/internal/vendor-group/category"
	"yasser-backend/pkg/models"

	"gorm.io/gorm"
)

type Vendor struct {
	models.BaseModel

	// Information
	NameEn        string `json:"nameEn" gorm:"size:100;not null" validate:"required"`
	NameAr        string `json:"nameAr" gorm:"size:100;not null" validate:"required"`
	DescriptionEn string `json:"descriptionEn" gorm:"size:255" validate:"max=255"`
	DescriptionAr string `json:"descriptionAr" gorm:"size:255" validate:"max=255"`
	ProfilePicture string `json:"profilePicture" gorm:"size:255"`
	CoverPicture   string `json:"coverPicture" gorm:"size:255"`
	Phone          string `json:"phone" gorm:"size:20" validate:"required"`
	Email          string `json:"email" gorm:"size:100" validate:"required,email"`

	// Location
	CityID uint      `json:"cityId" validate:"required"`
	City   city.City `gorm:"foreignKey:CityID"`

	DistrictID uint           `json:"districtId" validate:"required"`
	District   city.District  `gorm:"foreignKey:DistrictID"`

	AddressEn string `json:"addressEn" gorm:"size:255" validate:"required"`
	AddressAr string `json:"addressAr" gorm:"size:255" validate:"required"`

	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`

	// Opening and Closing Time
	OpeningTime string `json:"openingTime" gorm:"size:255" validate:"required"`
	ClosingTime string `json:"closingTime" gorm:"size:255" validate:"required"`

	// Category
	CategoryID uint                     `json:"categoryId" validate:"required"`
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

	DistrictID   uint   `json:"districtId"`
	DistrictName string `json:"districtName"`

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

// Scopes
func (r *repository) filterByCity(cityID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("city_id = ?", cityID)
	}
}

func (r *repository) filterByCategory(categoryID *uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if categoryID != nil {
			return db.Where("category_id = ?", *categoryID)
		}
		return db
	}
}

func (r *repository) filterActive() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("is_active = ?", true)
	}
}

func (r *repository) preloadRelations() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("City").
			Preload("District").
			Preload("Category")
	}
}


func (v *Vendor) ToResponse(lang string) *VendorResponse {
	response := &VendorResponse{
		ID:             v.ID,
		ProfilePicture: v.ProfilePicture,
		CoverPicture:   v.CoverPicture,
		Phone:          v.Phone,
		Email:          v.Email,
		CityID:         v.CityID,
		DistrictID:     v.DistrictID,
		Latitude:       v.Latitude,
		Longitude:      v.Longitude,
		OpeningTime:    v.OpeningTime,
		ClosingTime:    v.ClosingTime,
		CategoryID:     v.CategoryID,
		IsActive:       v.IsActive,
	}

	response.localize(lang, v)
	return response
}

func (vr *VendorResponse) localize(lang string, vendor *Vendor) {
	switch lang {
	case "ar":
		vr.Name = vendor.NameAr
		vr.Description = vendor.DescriptionAr
		vr.Address = vendor.AddressAr
		vr.CityName = getLocalizedName(vendor.City.NameAr, vendor.City.NameEn)
		vr.DistrictName = getLocalizedName(vendor.District.NameAr, vendor.District.NameEn)
		vr.CategoryName = vendor.Category.NameAr
	default:
		vr.Name = vendor.NameEn
		vr.Description = vendor.DescriptionEn
		vr.Address = vendor.AddressEn
		vr.CityName = getLocalizedName(vendor.City.NameEn, vendor.City.NameAr)
		vr.DistrictName = getLocalizedName(vendor.District.NameEn, vendor.District.NameAr)
		vr.CategoryName = vendor.Category.NameEn
	}
}

func getLocalizedName(primary, secondary string) string {
	if primary != "" {
		return primary
	}
	return secondary
}

type VendorFilter struct {
	CityID     uint
	CategoryID *uint
}