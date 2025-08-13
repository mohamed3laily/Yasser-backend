package locale
func ChooseLang(en, ar, lang string) string {
	if lang == "ar" {
		return ar
	}
	return en
}
