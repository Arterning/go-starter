package middleware

import (
	"net/http"
	"strings"

	"go-starter/pkg/utils"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userIDKey           = "user_id"
	userEmailKey        = "user_email"
	userUsernameKey     = "user_username"
)

func AuthMiddleware(jwtManager *utils.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(authorizationHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}

		token := parts[1]
		claims, err := jwtManager.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		c.Set(userIDKey, claims.UserID)
		c.Set(userEmailKey, claims.Email)
		c.Set(userUsernameKey, claims.Username)

		c.Next()
	}
}

func GetUserID(c *gin.Context) int {
	userID, exists := c.Get(userIDKey)
	if !exists {
		return 0
	}
	return userID.(int)
}
