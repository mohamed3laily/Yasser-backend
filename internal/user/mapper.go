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
        DistrictID:         u.DistrictID,
    }
}

func ToUserWithDistrictResponse(u *User, district *city.District) UserWithDistrictResponse {
    userResp := ToUserResponse(u)

    var districtResp *DistrictResponse

    if u.DistrictID != nil && district != nil && district.ID > 0 {
        lang := string(u.LanguagePreference)
        if lang == "" {
            lang = "en"
        }

        name := locale.ChooseLang(district.NameEn, district.NameAr, lang)

        resp := DistrictResponse{
            ID:        district.ID,
            CreatedAt: district.CreatedAt,
            Name:      name,
            MinLat:  district.MinLat,
            MinLng: district.MinLng,
            MaXLat:  district.MaxLat,
            MaxLng: district.MaxLng,
        }
        districtResp = &resp
    }


    return UserWithDistrictResponse{
        UserResponse: userResp,
        District:     districtResp,
    }
}