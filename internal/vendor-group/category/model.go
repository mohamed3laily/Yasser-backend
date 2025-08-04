package category

import (
	"yasser-backend/pkg/models"
)

type VendorCategory struct {
	models.BaseModel
	NameEn    string         `json:"nameEn" gorm:"size:100;not null" binding:"required,min=2,max=100"`
	NameAr    string         `json:"nameAr" gorm:"size:100;not null" binding:"required,min=2,max=100"`
	Color     string         `json:"color" gorm:"size:7;not null" binding:"required,len=7"`
	Icon      string         `json:"icon" gorm:"size:255;not null" binding:"required,url"`
}

func (vc *VendorCategory) GetLocalizedName(lang string) string {
	if lang == "ar" {
		return vc.NameAr
	}
	return vc.NameEn
}

type VendorCategoryResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	NameEn   string `json:"nameEn"`
	NameAr   string `json:"nameAr"`
	Color    string `json:"color"`
	Icon     string `json:"icon"`
}

func (vc *VendorCategory) ToResponse(lang string) *VendorCategoryResponse {
	return &VendorCategoryResponse{
		ID:       vc.ID,
		Name:     vc.GetLocalizedName(lang),
		NameEn:   vc.NameEn,
		NameAr:   vc.NameAr,
		Color:    vc.Color,
		Icon:     vc.Icon,
	}
}