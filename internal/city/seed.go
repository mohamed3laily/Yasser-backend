package city

import (
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	dakahlia := City{
		NameEn:   "Dakahlia",
		NameAr:   "الدقهلية",
		Latitude:  31.0379,
		Longitude: 31.3807,
	}

	if err := db.FirstOrCreate(&dakahlia, City{NameEn: "Dakahlia"}).Error; err != nil {
		return err
	}

	districts := []District{
		{
			NameEn: "Belqas",
			NameAr: "بلقاس",
			CityID: dakahlia.ID,
			MinLat: 31.1700,
			MaxLat: 31.2300,
			MinLng: 31.3200,
			MaxLng: 31.3900,
		},
		{
			NameEn: "Mansoura",
			NameAr: "المنصورة",
			CityID: dakahlia.ID,
			MinLat: 31.0200,
			MaxLat: 31.0700,
			MinLng: 31.3500,
			MaxLng: 31.4200,
		},
	}

	for _, d := range districts {
		if err := db.FirstOrCreate(&d, District{NameEn: d.NameEn, CityID: dakahlia.ID}).Error; err != nil {
			return err
		}
	}

	return nil
}
