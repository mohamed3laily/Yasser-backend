package auth

import (
	"database/sql"
	"time"
	"yasser-backend/database"
	"yasser-backend/internal/user"
)

type Repository interface {
	SetPhoneOtp(phone string, otp string) (*user.User, error)
	GetPhoneOtp(phone string) (*user.User, error)
	ClearPhoneOtp(phone string) error

}

type repo struct{}

func NewRepository() Repository {
	return &repo{}
}

var phoneNumberQuery = "phone_number = ?"

func (r *repo) SetPhoneOtp(phone string, otp string) (*user.User, error) {
	var user user.User
	result := database.DB.Where(phoneNumberQuery, phone).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	user.PhoneLoginOtp = sql.NullString{String: otp, Valid: true}
	user.PhoneLoginOtpExpires = sql.NullTime{Time: time.Now().Add(30 * time.Minute), Valid: true}
	return &user, database.DB.Save(&user).Error
}

func (r *repo) GetPhoneOtp(phone string) (*user.User, error) {
    var user user.User
    result := database.DB.Select("id, phone_number, phone_login_otp, phone_login_otp_expires").
        Where(phoneNumberQuery, phone).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func (r *repo) ClearPhoneOtp(phone string) error {
    var user user.User
    result := database.DB.Where(phoneNumberQuery, phone).First(&user)
    if result.Error != nil {
        return result.Error
    }
    user.PhoneLoginOtp =  sql.NullString{Valid: false}
	user.PhoneLoginOtpExpires = sql.NullTime{Valid: false}
    return database.DB.Save(&user).Error
}
