// internal/auth/repository.go
package auth

import (
	"context"
	"database/sql"
)

type authRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateUser(ctx context.Context, req CreateUserRequest) (User, error) {
	// TODO: implement actual db insert
	return User{}, nil
}

func (r *authRepository) Login(ctx context.Context, req LoginRequest) (User, error) {
	var user User
	return user, nil
}
