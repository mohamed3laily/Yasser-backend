package context

import (
	"github.com/gin-gonic/gin"
)

const (
	DefaultLanguage = "en"
)

func GetLanguage(c *gin.Context) string {
	if lang, exists := c.Get("lang"); exists {
		if language, ok := lang.(string); ok && language != "" {
			return language
		}
	}
	return DefaultLanguage
}

func GetUserID(c *gin.Context) uint {
	if userID, exists := c.Get("user_id"); exists {
		if id, ok := userID.(uint); ok {
			return id
		}
	}
	return 0
}