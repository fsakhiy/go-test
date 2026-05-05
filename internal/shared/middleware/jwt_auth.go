package middleware

import (
	"gin-test/internal/shared/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateAuth(secretKey string) gin.HandlerFunc {
	// Gin middlewares return a gin.HandlerFunc which takes *gin.Context
	return func(c *gin.Context) {

		// Gin uses GetHeader instead of Get
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Missing or invalid token",
			})

			// CRITICAL: You must call Abort() in Gin to stop the request chain
			c.Abort()
			return
		}

		// separate "bearer "
		token := strings.Split(authHeader, " ")[1]

		// validate
		claims, err := utils.ValidateJWT(token, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid token",
			})

			// CRITICAL: You must call Abort() in Gin to stop the request chain
			c.Abort()
			return
		}

		// Extract user_id from claims and store in context.
		// JWT MapClaims stores numbers as float64, so we convert to int64.
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid token: missing user_id",
			})
			c.Abort()
			return
		}
		c.Set("userId", int64(userIDFloat))

		// Pass control to the next middleware/handler
		c.Next()
	}
}
