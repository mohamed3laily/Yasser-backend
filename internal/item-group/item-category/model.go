package itemcategory

import "yasser-backend/pkg/models"

	type ItemsCategory struct {
		models.BaseModel
		NameAr    string `gorm:"type:text" json:"nameAr"`
		NameEn    string `gorm:"type:text" json:"nameEn"`
		Picture string `gorm:"type:text" json:"picture"`
	}