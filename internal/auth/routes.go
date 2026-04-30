package auth

import (
	"gin-test/internal/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, h *AuthHandler, jwtSecret string) {
	authGroup := router.Group("/auth")
	authGroup.POST("/register", h.CreateUser)
	authGroup.POST("/login", h.Login)

	authGroup.GET("/", middleware.ValidateAuth(jwtSecret), func(c *gin.Context) {
		c.String(200, "Authorized")
	})
}
