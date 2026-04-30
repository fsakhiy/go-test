// internal/auth/repository.go
package auth

import (
	"context"
	"database/sql"
	"gin-test/internal/auth/db"
)

type authRepository struct {
	database *sql.DB
	queries  *db.Queries
}

func NewRepository(database *sql.DB) Repository {
	return &authRepository{
		database: database,
		queries:  db.New(database),
	}
}

func (r *authRepository) CreateUser(ctx context.Context, req CreateUserRequest) (User, error) {
	arg := db.CreateUserParams{
		Username: req.Username,
		Password: req.Password,
	}

	res, err := r.queries.CreateUser(ctx, arg)
	if err != nil {
		return User{}, err
	}

	id, _ := res.LastInsertId()
	user := User{
		ID:        id,
		Username:  req.Username,
		Password:  req.Password,
		CreatedAt: nil,
		UpdatedAt: nil,
	}

	return user, nil
}

func (r *authRepository) Login(ctx context.Context, req LoginRequest) (User, error) {
	var user User
	return user, nil
}
