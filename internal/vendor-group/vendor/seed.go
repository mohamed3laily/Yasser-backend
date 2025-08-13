package vendor

import (
	"log"
	"math/rand"

	"yasser-backend/internal/city"
	"yasser-backend/internal/vendor-group/category"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	var dakahlia city.City
	if err := db.Where("name_en = ?", "Dakahlia").First(&dakahlia).Error; err != nil {
		return err
	}

	var districts []city.District
	if err := db.Where("city_id = ?", dakahlia.ID).Find(&districts).Error; err != nil {
		return err
	}

	var categories []category.VendorCategory
	if err := db.Find(&categories).Error; err != nil {
		return err
	}

	if len(districts) == 0 || len(categories) == 0 {
		return nil
	}

	names := map[string][]string{
		"en": {
			"Delicious Eats", "Mama's Kitchen", "Quick Bites", "Spice Garden",
			"Fresh Market", "Daily Grocer", "City Super", "Tasty Corner",
			"Golden Fork", "Ocean Delights",
		},
		"ar": {
			"أكلات لذيذة", "مطبخ ماما", "وجبات سريعة", "حديقة التوابل",
			"سوق طازج", "بقالة اليوم", "مطعم السبروت", "ركن الشهية",
			"الشوكة الذهبية", "نسمات البحر",
		},
	}

	phones := []string{
		"+201001111111", "+201002222222", "+201003333333", "+201004444444",
		"+201005555555", "+201006666666", "+201007777777", "+201008888888",
		"+201009999999", "+201001010101",
	}

	emails := []string{
		"contact@deliciouseats.com", "mama@kitchen.com", "quick@bites.com",
		"spice@garden.com", "fresh@market.com", "daily@grocer.com",
		"city@super.com", "tasty@corner.com", "golden@fork.com",
		"ocean@delights.com",
	}

	profilePics := []string{
		"https://upload.wikimedia.org/wikipedia/ar/a/a1/Albaik_logo.svg  ",
		"https://upload.wikimedia.org/wikipedia/ar/a/a1/Albaik_logo.svg  ",
		"https://upload.wikimedia.org/wikipedia/ar/a/a1/Albaik_logo.svg  ",
		"https://buffaloburger.com/_next/image?url=https%3A%2F%2Fbuffalonlineorderingprod.s3-accelerate.amazonaws.com%2Fmenu_items%2Fd845c9309b0d95d8c5d945b6b2552491.png&w=384&q=75  ",
		"https://buffaloburger.com/_next/image?url=https%3A%2F%2Fbuffalonlineorderingprod.s3-accelerate.amazonaws.com%2Fmenu_items%2Fd845c9309b0d95d8c5d945b6b2552491.png&w=384&q=75  ",
		"https://buffaloburger.com/_next/image?url=https%3A%2F%2Fbuffalonlineorderingprod.s3-accelerate.amazonaws.com%2Fmenu_items%2Fd845c9309b0d95d8c5d945b6b2552491.png&w=384&q=75  ",
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS0SXU1l2VmdHLRxHeMSWN-n8cdeKRucd5Tog&s",
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS0SXU1l2VmdHLRxHeMSWN-n8cdeKRucd5Tog&s",
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS0SXU1l2VmdHLRxHeMSWN-n8cdeKRucd5Tog&s",
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS0SXU1l2VmdHLRxHeMSWN-n8cdeKRucd5Tog&s",
	}

	coverPics := []string{
		"https://watanimg.elwatannews.com/image_archive/648x316/20451042101567622797.jpg  ",
		"https://watanimg.elwatannews.com/image_archive/648x316/20451042101567622797.jpg  ",
		"https://watanimg.elwatannews.com/image_archive/648x316/20451042101567622797.jpg  ",
		"https://buffalonlineorderingprod.s3.eu-west-1.amazonaws.com/carousels/d4540a13ef8f58f47f31e30c9faaa7dc.png  ",
		"https://buffalonlineorderingprod.s3.eu-west-1.amazonaws.com/carousels/d4540a13ef8f58f47f31e30c9faaa7dc.png  ",
		"https://buffalonlineorderingprod.s3.eu-west-1.amazonaws.com/carousels/d4540a13ef8f58f47f31e30c9faaa7dc.png  ",
		"https://www.amnesty.org/ar/wp-content/uploads/2024/10/302466-1-scaled.jpg  ",
		"https://www.amnesty.org/ar/wp-content/uploads/2024/10/302466-1-scaled.jpg  ",
		"https://www.amnesty.org/ar/wp-content/uploads/2024/10/302466-1-scaled.jpg  ",
		"https://www.amnesty.org/ar/wp-content/uploads/2024/10/302466-1-scaled.jpg  ",
	}

	openingTimes := []string{"08:00", "09:00", "10:00", "07:00", "06:00"}
	closingTimes := []string{"22:00", "23:00", "00:00", "21:00", "24:00"}

	addressesEn := []string{
		"Main Street", "Downtown Area", "Near University", "Shopping Mall",
		"Central Square", "Old Town", "New District", "Riverside",
		"Business Zone", "Residential Area",
	}

	addressesAr := []string{
		"الشارع الرئيسي", "المنطقة المركزية", "بالقرب من الجامعة", "المول التجاري",
		"الساحة المركزية", "البلدة القديمة", "الحي الجديد", "على ضفاف النيل",
		"المنطقة التجارية", "المنطقة السكنية",
	}

	// Generate 10 vendors
	for i := 0; i < 10; i++ {
		district := districts[rand.Intn(len(districts))]
		category := categories[rand.Intn(len(categories))]

		vendor := Vendor{
			// Info
			NameEn:         names["en"][i],
			NameAr:         names["ar"][i],
			DescriptionEn:  "Great food and service",
			DescriptionAr:  "طعام رائع وخدمة ممتازة",
			ProfilePicture: profilePics[i],
			CoverPicture:   coverPics[i],
			Phone:          phones[i],
			Email:          emails[i],

			// Location
			CityID:      dakahlia.ID,
			DistrictID:  district.ID,
			AddressEn:   addressesEn[i] + ", " + district.NameEn,
			AddressAr:   addressesAr[i] + ", " + district.NameAr,
			Latitude:    district.MinLat + (district.MaxLat-district.MinLat)*rand.Float64(),
			Longitude:   district.MinLng + (district.MaxLng-district.MinLng)*rand.Float64(),

			// Time
			OpeningTime: openingTimes[rand.Intn(len(openingTimes))],
			ClosingTime: closingTimes[rand.Intn(len(closingTimes))],

			// Category
			CategoryID: category.ID,

			// Status
			IsActive: true,
		}
		var existing Vendor
		
		err := db.Where("name_en = ? AND phone = ?", vendor.NameEn, vendor.Phone).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&vendor).Error; err != nil {
				log.Printf("Failed to create vendor %s: %v", vendor.NameEn, err)
			} else {
				log.Printf("✅ Created vendor: %s", vendor.NameEn)
			}
		} else {
			log.Printf("⏭️ Vendor already exists: %s", vendor.NameEn)
		}
	}

	log.Println("✅ Seeded 10 vendors successfully.")
	return nil
}