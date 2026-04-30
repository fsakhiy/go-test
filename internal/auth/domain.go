package auth

import (
	"context"
	"time"
)

type User struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Repository interface {
	CreateUser(ctx context.Context, user CreateUserRequest) (User, error)
	FindOneByUsername(ctx context.Context, username string) (User, error)
}

type Service interface {
	CreateUser(ctx context.Context, user CreateUserRequest) (User, error)
	Login(ctx context.Context, user LoginRequest) (string, error)
}
