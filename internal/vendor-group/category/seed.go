package category

import (
	"log"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	vendorCategories := []VendorCategory{
		{
			NameEn: "Restaurants",
			NameAr: "مطاعم",
			Color:  "#FF5733",
			Icon:   "https://icon-library.com/images/restaurant-icon/restaurant-icon-0.jpg",
		},
		{
			NameEn: "Supermarkets",
			NameAr: "سوبر ماركت",
			Color:  "#33C1FF",
			Icon:   "https://icon-library.com/images/supermarket-icon/supermarket-icon-1.jpg",
		},
	}

	for _, vc := range vendorCategories {
		var existing VendorCategory
		err := db.Where("name_en = ? AND name_ar = ?", vc.NameEn, vc.NameAr).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&vc).Error; err != nil {
				log.Printf("Failed to seed VendorCategory: %v", err)
			}
		}
	}

	log.Println("Seeded vendor categories")
	return nil
}
