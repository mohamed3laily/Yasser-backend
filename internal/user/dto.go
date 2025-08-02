package user

import "fmt"

type UpdateUserRequest struct {
	FullName           string   `json:"fullName"`
	FCMToken           string   `json:"fcmToken"`
	LanguagePreference Language `json:"languagePreference"`
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