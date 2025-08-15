package item

import (
	"sort"
	"yasser-backend/pkg/locale"
)

type ItemResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	BasePrice   int     `json:"basePrice"`
	Price       float64 `json:"price"`
	DiscountPercent float64 `json:"discountPercent"`
	Picture     string  `json:"picture,omitempty"`
	CategoryID  int64   `json:"categoryId"`
	VendorID    int64   `json:"vendorId"`
}

type ItemWithCategory struct {
	ID             int64
	NameEn         string
	NameAr         string
	DescriptionEn  string
	DescriptionAr  string
	BasePrice      int
	DiscountPercent       float64
	Picture        string
	CategoryID     int64
	VendorID       int64
	CategoryNameEn string
	CategoryNameAr string
}

type VariantResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type SizeResponse struct {
	ID        int64 `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	BasePrice int    `json:"basePrice"`
}

type AddonResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	IsRemoval bool   `json:"isRemoval"`
}

type ItemDetailResponse struct {
	ID          int64            `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	BasePrice   int              `json:"basePrice"`
	DiscountPercent float64      `json:"discountPercent"`
	Price       float64          `json:"price"`
	Picture     string           `json:"picture"`
	VendorID    int64            `json:"vendorId"`
	Variants    []VariantResponse `json:"variants"`
	Sizes       []SizeResponse    `json:"sizes"`
	Addons      []AddonResponse   `json:"addons"`
	Removals    []AddonResponse   `json:"removals"`
}

type CategoryWithItemsResponse struct {
	Category string         `json:"category"`
	Items    []ItemResponse `json:"items"`
}

func ToItemDetailResponse(i *Item, lang string) *ItemDetailResponse {
	name := i.NameEn
	desc := i.DescriptionEn
	if lang == "ar" {
		name = i.NameAr
		desc = i.DescriptionAr
	}
	finalPrice := float64(i.BasePrice) * (1 - i.DiscountPercent/100.0)

	variants := make([]VariantResponse, len(i.Variants))
	for idx, v := range i.Variants {
		variants[idx] = VariantResponse{
			ID:   int64(v.ID),
			Name: locale.ChooseLang(v.NameEn, v.NameAr, lang),
		}
	}

	sizes := make([]SizeResponse, len(i.Sizes))
	for idx, s := range i.Sizes {
		finalSizePrice := float64(s.Price) * (1 - i.DiscountPercent/100.0)

		sizes[idx] = SizeResponse{
			ID:        int64(s.ID),
			Name:      locale.ChooseLang(s.Name, s.Name, lang),
			Price:     int(finalSizePrice),
			BasePrice: s.Price,
		}
	}

	var addons []AddonResponse
	var removals []AddonResponse
	for _, a := range i.Addons {
		resp := AddonResponse{
			ID:        int64(a.ID),
			Name:      locale.ChooseLang(a.NameEn, a.NameAr, lang),
			Price:     a.Price,
			IsRemoval: a.IsRemoval,
		}
		if a.IsRemoval {
			removals = append(removals, resp)
		} else {
			addons = append(addons, resp)
		}
	}

	return &ItemDetailResponse{
		ID:          int64(i.ID),
		Name:        name,
		Description: desc,
		BasePrice:   i.BasePrice,
		Price:       finalPrice,
		DiscountPercent: i.DiscountPercent,
		Picture:     i.Picture,
		VendorID:    i.VendorID,
		Variants:    variants,
		Sizes:       sizes,
		Addons:      addons,
		Removals:    removals,
	}
}

func GroupItemsByCategory(items []ItemWithCategory, lang string) []CategoryWithItemsResponse {
	categoryMap := make(map[int64]*CategoryWithItemsResponse)

	categoryNames := make(map[int64]string)

	for _, i := range items {
		name := locale.ChooseLang(i.NameEn, i.NameAr, lang)
		desc := locale.ChooseLang(i.DescriptionEn, i.DescriptionAr, lang)

		categoryName := locale.ChooseLang(i.CategoryNameEn, i.CategoryNameAr, lang)
		categoryNames[i.CategoryID] = categoryName

		finalPrice := float64(i.BasePrice) * (1 - i.DiscountPercent/100.0)

		if _, ok := categoryMap[i.CategoryID]; !ok {
			categoryMap[i.CategoryID] = &CategoryWithItemsResponse{
				Category: categoryName,
				Items:    []ItemResponse{},
			}
		}

		categoryMap[i.CategoryID].Items = append(categoryMap[i.CategoryID].Items, ItemResponse{
			ID:          i.ID,
			Name:        name,
			Description: desc,
			BasePrice:   i.BasePrice,
			Price:       finalPrice,
			Picture:     i.Picture,
			CategoryID:  i.CategoryID,
			VendorID:    i.VendorID,
		})
	}

	result := make([]CategoryWithItemsResponse, 0, len(categoryMap))
	for categoryID, v := range categoryMap {
		v.Category = categoryNames[categoryID]
		result = append(result, *v)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Category < result[j].Category
	})

	return result
}