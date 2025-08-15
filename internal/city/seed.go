package city

import (
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	// Dakahlia
	dakahlia := City{
		NameEn:   "Dakahlia",
		NameAr:   "الدقهلية",
		Latitude: 31.0379,
		Longitude: 31.3807,
	}
	if err := db.FirstOrCreate(&dakahlia, City{NameEn: "Dakahlia"}).Error; err != nil {
		return err
	}

	dakahliaDistricts := []District{
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

	for _, d := range dakahliaDistricts {
		if err := db.FirstOrCreate(&d, District{NameEn: d.NameEn, CityID: dakahlia.ID}).Error; err != nil {
			return err
		}
	}

	// Damietta
	damietta := City{
		NameEn:   "Damietta",
		NameAr:   "دمياط",
		Latitude: 31.4165,
		Longitude: 31.8133,
	}
	if err := db.FirstOrCreate(&damietta, City{NameEn: "Damietta"}).Error; err != nil {
		return err
	}

	damiettaDistricts := []District{
		{
			NameEn: "Ras El Bar",
			NameAr: "رأس البر",
			CityID: damietta.ID,
			MinLat: 31.5100,
			MaxLat: 31.5300,
			MinLng: 31.7900,
			MaxLng: 31.8200,
		},
		{
			NameEn: "Ezbet El Borj",
			NameAr: "عزبة البرج",
			CityID: damietta.ID,
			MinLat: 31.5200,
			MaxLat: 31.5500,
			MinLng: 31.8100,
			MaxLng: 31.8400,
		},
	}

	for _, d := range damiettaDistricts {
		if err := db.FirstOrCreate(&d, District{NameEn: d.NameEn, CityID: damietta.ID}).Error; err != nil {
			return err
		}
	}

	return nil
}
