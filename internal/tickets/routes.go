// internal/tickets/routes.go
package tickets

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes takes a Gin RouterGroup, the TicketHandler, and the auth middleware.
func RegisterRoutes(router *gin.RouterGroup, h *TicketHandler, authMiddleware gin.HandlerFunc) {

	// 1. Create a specific group for this module.
	// By passing authMiddleware here, every route inside "/ticket" is automatically protected!
	ticketGroup := router.Group("/ticket", authMiddleware)

	// 2. Map the HTTP methods to your Handler functions.
	// Note: Gin uses ALL CAPS for HTTP methods (GET, POST, PUT, DELETE).
	ticketGroup.GET("/test", h.Test)
	ticketGroup.POST("/", h.NewTicket)
	ticketGroup.GET("/", h.GetAll)
	ticketGroup.PUT("/:refNo", h.Update)
	ticketGroup.DELETE("/:refNo", h.Delete)
}
