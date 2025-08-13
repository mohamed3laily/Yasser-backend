package auth

import (
	"net/http"
	"strings"
	"yasser-backend/internal/auth/jwt"
	"yasser-backend/internal/user"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		u, err := user.GetUserByID(claims.UserID)
		if err != nil || u.Status != user.ACTIVE {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found or inactive"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("user", u)
		
		c.Next()
	}
}
