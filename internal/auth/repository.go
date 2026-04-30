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

func (r *authRepository) FindOneByUsername(ctx context.Context, username string) (User, error) {
	res, err := r.queries.FindUserByUsername(ctx, username)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:        int64(res.ID),
		Username:  res.Username,
		Password:  res.Password,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}
