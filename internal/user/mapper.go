package user

import (
	"yasser-backend/internal/city"
	"yasser-backend/pkg/locale"
)

func ToUserResponse(u *User) UserResponse {
    return UserResponse{
        ID:                 u.ID,
        CreatedAt:          u.CreatedAt,
        FullName:           u.FullName,
        ProfilePicture:     u.ProfilePicture,
        PhoneNumber:        u.PhoneNumber,
        LanguagePreference: string(u.LanguagePreference),
        FCMToken:           u.FCMToken,
        LastLogin:          u.LastLogin,
        CityID:             u.CityID,
    }
}

func ToUserWithCityResponse(u *User, city *city.City) UserWithCityResponse {
    userResp := ToUserResponse(u)

    var cityResp *CityResponse

    if u.CityID != nil && city != nil && city.ID > 0 {
        lang := string(u.LanguagePreference)
        if lang == "" {
            lang = "en"
        }

        name := locale.ChooseLang(city.NameEn, city.NameAr, lang)

        resp := CityResponse{
            ID:        city.ID,
            CreatedAt: city.CreatedAt,
            Name:      name,
            Latitude:  city.Latitude,
            Longitude: city.Longitude,
        }
        cityResp = &resp
    }


    return UserWithCityResponse{
        UserResponse: userResp,
        City:         cityResp,
    }
}