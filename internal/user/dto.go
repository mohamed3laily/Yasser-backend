package user

import (
	"fmt"
	"time"
)

type UpdateUserRequest struct {
	FullName           string   `json:"fullName"`
	FCMToken           string   `json:"fcmToken"`
	LanguagePreference Language `json:"languagePreference"`
	CityId             *uint    `json:"cityId"`
}

func (r UpdateUserRequest) Validate() error {
	if r.LanguagePreference != "" && !r.LanguagePreference.IsValid() {
		return fmt.Errorf(
			"invalid language preference: %s. Available values: %q",
			r.LanguagePreference,
			[]Language{EN, AR},
		)
	}
	return nil
}

type UserResponse struct {
    ID                 uint      `json:"id"`
    CreatedAt          time.Time `json:"createdAt"`
    FullName           string    `json:"fullName"`
    ProfilePicture     string    `json:"profilePicture"`
    PhoneNumber        string    `json:"phoneNumber"`
    LanguagePreference string    `json:"languagePreference"`
    FCMToken           string    `json:"fcmToken"`
    LastLogin          time.Time `json:"lastLogin"`
    CityID             *uint     `json:"cityId"`
}

type UserWithCityResponse struct {
    UserResponse
    City *CityResponse `json:"city,omitempty"`
}
type CityResponse struct {
    ID        uint      `json:"id"`
    CreatedAt time.Time `json:"createdAt"`
    Name      string    `json:"name"`
    Latitude  float64   `json:"latitude"`
    Longitude float64   `json:"longitude"`
}