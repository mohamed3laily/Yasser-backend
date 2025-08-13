package city

import (
	"yasser-backend/pkg/models"

	"gorm.io/gorm"
)

type City struct {
	models.BaseModel
	NameEn      string         `gorm:"not null;uniqueIndex"`
	NameAr      string         `gorm:"not null;uniqueIndex"`
	Latitude  float64 // center of the city
	Longitude float64
	Areas     []District         `gorm:"foreignKey:CityID"`
}

func SearchCityByName(query string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if query == "" {
			return db
		}
		return db.Where("name_en ILIKE ? OR name_ar ILIKE ?", "%"+query+"%", "%"+query+"%")
	}
}