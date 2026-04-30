package auth

import (
	"gin-test/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service Service
}

func NewHandler(service Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest

	//Bind the request body to the struct
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, 400, "validation error", err.Error())
		return
	}

	// Call service
	user, err := h.service.CreateUser(c, req)
	if err != nil {
		response.Err(c, 500, "internal server error", err.Error())
		return
	}

	// Respond with 201 Created
	// response.Created(c, "User created successfully", user)
	response.OK(c, 201, "Created User", user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	//Bind the request body to the struct
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, 400, "validation error", err.Error())
		return
	}

	// Call service
	user, err := h.service.Login(c, req)
	if err != nil {
		response.Err(c, 500, "internal server error", err.Error())
		return
	}

	// Respond with 201 Created
	// response.Created(c, "User created successfully", user)
	response.OK(c, 200, "Login", user)
}
