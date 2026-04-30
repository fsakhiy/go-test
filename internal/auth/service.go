package auth

import (
	"context"
	"fmt"
)

type ticketService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &ticketService{repo: repo}
}

func (s *ticketService) CreateUser(ctx context.Context, req CreateUserRequest) (User, error) {
	// return s.repo.CreateUser(ctx, req)

	fmt.Println(req)

	return User{
		Username: req.Username,
		Password: req.Password,
	}, nil
}

func (s *ticketService) Login(ctx context.Context, req LoginRequest) (User, error) {
	return s.repo.Login(ctx, req)
}
