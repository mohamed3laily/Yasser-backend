package city

import (
	"yasser-backend/pkg/models"

	"gorm.io/gorm"
)

type District struct {
	models.BaseModel
	NameEn      string         `gorm:"not null"`
	NameAr      string         `gorm:"not null"`
	CityID    uint           `gorm:"index;not null"`
	MinLat float64 `gorm:"not null"`
	MaxLat float64 `gorm:"not null"`
	MinLng float64 `gorm:"not null"`
	MaxLng float64 `gorm:"not null"`
	City      City           `gorm:"foreignKey:CityID"`

}

func DistrictByCity(cityID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("city_id = ?", cityID)
	}
}

func SearchDistrictByName(query string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if query == "" {
			return db
		}
		return db.Where("name_en ILIKE ? OR name_ar ILIKE ?", "%"+query+"%", "%"+query+"%")
	}
}