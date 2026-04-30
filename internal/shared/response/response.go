package response

import (
	"github.com/gin-gonic/gin"
)

// import "github.com/gofiber/fiber/v3"

// 1. The unified JSON structure
type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`  // omitempty hides it if it's nil
	Error   any    `json:"error,omitempty"` // omitempty hides it if there's no error
}

// 2. Helper for Success Responses
func OK[T any](c *gin.Context, status int, message string, data T) {
	c.JSON(status, APIResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// 3. Helper for Error Responses
func Err(c *gin.Context, status int, message string, errDetail any) {
	// We use 'any' here because errors usually don't return data
	c.JSON(status, APIResponse[any]{
		Success: false,
		Message: message,
		Error:   errDetail,
	})

	c.Abort()
}
