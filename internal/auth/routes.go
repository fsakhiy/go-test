package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, h *AuthHandler) {
	authGroup := router.Group("/auth")
	authGroup.POST("/register", h.CreateUser)
	authGroup.POST("/login", h.Login)
}
