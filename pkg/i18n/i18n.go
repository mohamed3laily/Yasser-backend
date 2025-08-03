package i18n

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type Localizer struct {
	translations map[string]map[string]string
	fallback     string
}

var globalLocalizer *Localizer

func Init(translationsDir string) error {
	localizer := &Localizer{
		translations: make(map[string]map[string]string),
		fallback:     "en",
	}

	languages := []string{"en", "ar"}
	
	for _, lang := range languages {
		langDir := filepath.Join(translationsDir, lang)
		if err := localizer.loadLanguageDirectory(lang, langDir); err != nil {
			return fmt.Errorf("failed to load translations for %s: %w", lang, err)
		}
	}

	globalLocalizer = localizer
	return nil
}

func (l *Localizer) loadLanguageDirectory(lang, langDir string) error {
	if l.translations[lang] == nil {
		l.translations[lang] = make(map[string]string)
	}

	files, err := filepath.Glob(filepath.Join(langDir, "*.json"))
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := l.loadTranslationFile(lang, file); err != nil {
			return fmt.Errorf("failed to load file %s: %w", file, err)
		}
	}

	return nil
}

func (l *Localizer) loadTranslationFile(lang, filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	var translations map[string]string
	if err := json.Unmarshal(data, &translations); err != nil {
		return err
	}

	filename := filepath.Base(filePath)
	moduleName := strings.TrimSuffix(filename, filepath.Ext(filename))

	for key, value := range translations {
		fullKey := fmt.Sprintf("%s.%s", moduleName, key)
		l.translations[lang][fullKey] = value
	}

	return nil
}

func (l *Localizer) Get(lang, key string, params ...interface{}) string {
	if translations, exists := l.translations[lang]; exists {
		if translation, exists := translations[key]; exists {
			if len(params) > 0 {
				return fmt.Sprintf(translation, params...)
			}
			return translation
		}
	}

	if translations, exists := l.translations[l.fallback]; exists {
		if translation, exists := translations[key]; exists {
			if len(params) > 0 {
				return fmt.Sprintf(translation, params...)
			}
			return translation
		}
	}

	return key
}

func LanguageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := extractLanguage(c)
		c.Set("lang", lang)
		c.Next()
	}
}

func extractLanguage(c *gin.Context) string {
	acceptLang := c.GetHeader("Accept-Language")
	fmt.Println("Accept-Language:", acceptLang)
	if acceptLang != "" {
		langs := strings.Split(acceptLang, ",")
		for _, lang := range langs {
			langCode := strings.TrimSpace(strings.Split(lang, ";")[0])
			
			primaryLang := strings.Split(langCode, "-")[0]
			
			if primaryLang == "en" || primaryLang == "ar" {
				return primaryLang
			}
		}
	}

	return "en"
}

func T(c *gin.Context, key string, params ...interface{}) string {
	if globalLocalizer == nil {
		return key
	}
	
	lang, exists := c.Get("lang")
	if !exists {
		lang = "en"
	}
	
	return globalLocalizer.Get(lang.(string), key, params...)
}

func TWithLang(lang, key string, params ...interface{}) string {
	if globalLocalizer == nil {
		return key
	}
	
	return globalLocalizer.Get(lang, key, params...)
}