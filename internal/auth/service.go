package auth

import (
	"context"
)

type ticketService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &ticketService{repo: repo}
}

func (s *ticketService) CreateUser(ctx context.Context, req CreateUserRequest) (User, error) {
	return s.repo.CreateUser(ctx, req)
}

func (s *ticketService) Login(ctx context.Context, req LoginRequest) (User, error) {
	return s.repo.Login(ctx, req)
}
