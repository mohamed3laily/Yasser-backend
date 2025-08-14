package user

import (
	"database/sql"
	"time"
	"yasser-backend/internal/city"
	"yasser-backend/pkg/models"
)

type Status string
type Language string

const (
	ACTIVE      Status = "ACTIVE"
	DEACTIVATED Status = "DEACTIVATED"
	BANNED      Status = "BANNED"
)

const (
	EN Language = "en"
	AR Language = "ar"
)

type User struct {
	models.BaseModel
	FullName                string         `json:"fullName"`
	ProfilePicture          string         `json:"profilePicture"`
	PhoneNumber             string         `gorm:"unique" json:"phoneNumber"`
	PhoneLoginOtp           sql.NullString `json:"-"`
	PhoneLoginOtpExpires    sql.NullTime   `json:"-"`
	GoogleID                string         `json:"-"`
	Status                  Status         `gorm:"type:varchar(20);default:ACTIVE" json:"-"`
	LanguagePreference      Language       `gorm:"type:varchar(10);default:en" json:"languagePreference"`
	FCMToken                string         `json:"fcmToken"`
	LastLogin               time.Time      `json:"lastLogin"`

	DistrictID 				*uint 		   `json:"districtId"`
	District                city.District  `gorm:"foreignKey:DistrictID"`
}

func (s Status) IsValid() bool {
	switch s {
	case ACTIVE, DEACTIVATED, BANNED:
		return true
	}
	return false
}

func (l Language) IsValid() bool {
	switch l {
	case EN, AR:
		return true
	}
	return false
}