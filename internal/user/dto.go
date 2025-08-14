package user

import (
	"fmt"
	"time"
)

type UpdateUserRequest struct {
	FullName           string   `json:"fullName"`
	FCMToken           string   `json:"fcmToken"`
	LanguagePreference Language `json:"languagePreference"`
	DistrictID         *uint    `json:"districtId"`
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
    DistrictID         *uint     `json:"districtId"`
}

type UserWithDistrictResponse struct {
    UserResponse
    District *DistrictResponse `json:"district,omitempty"`
}
type DistrictResponse struct {
    ID        uint      `json:"id"`
    CreatedAt time.Time `json:"createdAt"`
    Name      string    `json:"name"`
    MinLat    float64   `json:"minLat"`
    MinLng    float64   `json:"minLng"`
    MaXLat    float64   `json:"maxLat"`
    MaxLng    float64   `json:"maxLng"`
}