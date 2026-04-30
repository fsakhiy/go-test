package middleware

import (
	"fmt"
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

		fmt.Printf("=== middleware test === : %s \n", authHeader)

		// Pass control to the next middleware/handler
		c.Next()
	}
}
