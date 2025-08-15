package vendor

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"yasser-backend/internal/city"
	"yasser-backend/internal/vendor-group/category"

	"gorm.io/gorm"
)

type namePair struct {
	En string
	Ar string
}

func Seed(db *gorm.DB) error {
	// Get all cities (Dakahlia + Damietta present from city seeder)
	var cities []city.City
	if err := db.Find(&cities).Error; err != nil {
		return err
	}

	// Categories
	var categories []category.VendorCategory
	if err := db.Find(&categories).Error; err != nil {
		return err
	}
	if len(cities) == 0 || len(categories) == 0 {
		return nil
	}

	// Popular brands (split 10/10 per city)
	namesByCity := map[string][]namePair{
		"Dakahlia": {
			{"KFC", "كنتاكي"},
			{"McDonald's", "ماكدونالدز"},
			{"Pizza Hut", "بيتزا هت"},
			{"Cook Door", "كوك دور"},
			{"Gad", "جاد"},
			{"El Tahrir", "التحرير"},
			{"Abu Shakra", "أبو شقرة"},
			{"Zooba", "زووبا"},
			{"Kazouza", "كازوزة"},
			{"El Shamy", "الشامي"},
		},
		"Damietta": {
			{"Carrefour", "كارفور"},
			{"HyperOne", "هايبر وان"},
			{"Metro Market", "مترو ماركت"},
			{"Seoudi", "سعودي"},
			{"Oscar Grand Stores", "أوسكار"},
			{"Ragab Sons", "أولاد رجب"},
			{"Spinneys", "سبينيس"},
			{"Kazyon", "كازيون"},
			{"Fresh Food Market", "فريش فود ماركت"},
			{"Alfa Market", "ألفا ماركت"},
		},
	}

	// Keep phones as you had (10 items)
	phones := []string{
		"+201001111111", "+201002222222", "+201003333333", "+201004444444",
		"+201005555555", "+201006666666", "+201007777777", "+201008888888",
		"+201009999999", "+201001010101",
	}

	// 👇 EXACT same images you provided originally
	profilePics := []string{
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQyYyH4UZCH7NcomOdvL2Lk8ZLVddiGknorPQ&s",
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQyYyH4UZCH7NcomOdvL2Lk8ZLVddiGknorPQ&s",
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQyYyH4UZCH7NcomOdvL2Lk8ZLVddiGknorPQ&s",
		"https://static.vecteezy.com/system/resources/previews/008/058/752/non_2x/supermarket-logo-design-vector.jpg",
		"https://static.vecteezy.com/system/resources/previews/008/058/752/non_2x/supermarket-logo-design-vector.jpg",
		"https://static.vecteezy.com/system/resources/previews/008/058/752/non_2x/supermarket-logo-design-vector.jpg",
		"https://thumbs.dreamstime.com/b/grocery-shop-emblem-food-store-supermarket-market-vector-icon-sign-silhouette-shopping-bag-products-baguette-361840126.jpg",
		"https://thumbs.dreamstime.com/b/grocery-shop-emblem-food-store-supermarket-market-vector-icon-sign-silhouette-shopping-bag-products-baguette-361840126.jpg",
		"https://thumbs.dreamstime.com/b/grocery-shop-emblem-food-store-supermarket-market-vector-icon-sign-silhouette-shopping-bag-products-baguette-361840126.jpg",
		"https://thumbs.dreamstime.com/b/grocery-shop-emblem-food-store-supermarket-market-vector-icon-sign-silhouette-shopping-bag-products-baguette-361840126.jpg",
	}
	coverPics := []string{
		"https://media.licdn.com/dms/image/v2/D4E12AQGuQGl0Hhpq_w/article-cover_image-shrink_720_1280/article-cover_image-shrink_720_1280/0/1696413376377?e=2147483647&v=beta&t=QKd-dJKtHK2tu4hcB_wU7qM4C4mceHBfUPeyGuXv48M",
		"https://i0.wp.com/www.laundryheap.com/blog/wp-content/uploads/2023/03/table-g9cfa78ffd_1920.jpg?resize=768%2C512&ssl=1",
		"https://images.ctfassets.net/3vektz2wjwwa/7cI24jNdK9dgAw06UnQWB7/46c4c1c3cb99b80d3ce15e72369b35f3/10-things-to-consider-before-setting-up-a-restaurant-1600x727.jpg?fm=webp",
		"https://dynamic-media-cdn.tripadvisor.com/media/photo-o/11/4f/8d/4e/renoufs-sharing-platters.jpg?w=900&h=500&s=1",		"https://dynamic-media-cdn.tripadvisor.com/media/photo-o/11/4f/8d/4e/renoufs-sharing-platters.jpg?w=900&h=500&s=1",
		"https://buffalonlineorderingprod.s3.eu-west-1.amazonaws.com/carousels/d4540a13ef8f58f47f31e30c9faaa7dc.png",
		"https://www.amnesty.org/ar/wp-content/uploads/2024/10/302466-1-scaled.jpg",
		"https://dynamic-media-cdn.tripadvisor.com/media/photo-o/16/e4/e9/64/the-stunning-outside.jpg?w=900&h=500&s=1",
		"https://dynamic-media-cdn.tripadvisor.com/media/photo-o/16/e4/e9/64/the-stunning-outside.jpg?w=900&h=500&s=1",
		"https://dynamic-media-cdn.tripadvisor.com/media/photo-o/11/4f/8d/4e/renoufs-sharing-platters.jpg?w=900&h=500&s=1",	}

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

	// For each city, seed 10 vendors
	for _, c := range cities {
		// pull districts for this city
		var districts []city.District
		if err := db.Where("city_id = ?", c.ID).Find(&districts).Error; err != nil {
			return err
		}
		if len(districts) == 0 {
			continue
		}

		// choose name set by city (fallback: reuse Dakahlia set)
		nameSet, ok := namesByCity[c.NameEn]
		if !ok || len(nameSet) < 10 {
			nameSet = namesByCity["Dakahlia"]
		}

		for i := 0; i < 10; i++ {
			d := districts[rand.Intn(len(districts))]
			category := categories[rand.Intn(len(categories))]
			name := nameSet[i]

			emailLocal := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(name.En, " ", ""), "'", ""))
			emailDomain := strings.ToLower(strings.ReplaceAll(c.NameEn, " ", ""))
			email := fmt.Sprintf("%s@%s.demo", emailLocal, emailDomain)

			vendor := Vendor{
				// Info
				NameEn:         name.En,
				NameAr:         name.Ar,
				DescriptionEn:  "Great food and service",
				DescriptionAr:  "طعام رائع وخدمة ممتازة",
				ProfilePicture: profilePics[i],
				CoverPicture:   coverPics[i],
				Phone:          phones[i],
				Email:          email,

				// Location
				CityID:     c.ID,
				DistrictID: d.ID,
				AddressEn:  addressesEn[rand.Intn(len(addressesEn))] + ", " + d.NameEn,
				AddressAr:  addressesAr[rand.Intn(len(addressesAr))] + ", " + d.NameAr,
				Latitude:   d.MinLat + (d.MaxLat-d.MinLat)*rand.Float64(),
				Longitude:  d.MinLng + (d.MaxLng-d.MinLng)*rand.Float64(),

				// Time
				OpeningTime: openingTimes[rand.Intn(len(openingTimes))],
				ClosingTime: closingTimes[rand.Intn(len(closingTimes))],

				// Category & status
				CategoryID: category.ID,
				IsActive:   true,
			}
			var existing Vendor
			err := db.Where("name_en = ? AND phone = ?", vendor.NameEn, vendor.Phone).First(&existing).Error
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&vendor).Error; err != nil {
					log.Printf("❌ Failed to create vendor %s in %s: %v", vendor.NameEn, c.NameEn, err)
				} else {
					log.Printf("✅ Created vendor %s in %s (%s)", vendor.NameEn, c.NameEn, d.NameEn)
				}
			} else {
				log.Printf("⏭️ Vendor already exists: %s", vendor.NameEn)
			}
		}
	}

	log.Println("✅ Seeded 10 vendors per city successfully.")
	return nil
}
