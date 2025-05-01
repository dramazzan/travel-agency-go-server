package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-go/internal/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("authToken")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authToken cookie missing"})
			c.Abort()
			return
		}

		tokenString := cookie.Value
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		userIDFloat, ok := claims["userID"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
			c.Abort()
			return
		}

		userID := uint(userIDFloat)

		c.Set("userID", userID)
		c.Set("username", claims["username"])
		c.Set("role", claims["role"])
		c.Next()
	}
}
