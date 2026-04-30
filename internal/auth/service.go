package auth

import (
	"context"
	"errors"
	"gin-test/internal/shared/utils"
)

type ticketService struct {
	repo      Repository
	JWTSecret []byte
}

func NewService(repo Repository, jwtSecret string) Service {
	return &ticketService{repo: repo, JWTSecret: []byte(jwtSecret)}
}

func (s *ticketService) CreateUser(ctx context.Context, req CreateUserRequest) (User, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return User{}, err
	}

	return s.repo.CreateUser(ctx, CreateUserRequest{
		Username: req.Username,
		Password: hashedPassword,
	})
}

func (s *ticketService) Login(ctx context.Context, req LoginRequest) (string, error) {
	// return s.repo.FindOneByUsername(ctx, req)

	user, err := s.repo.FindOneByUsername(ctx, req.Username)
	if err != nil {
		return "", err
	}

	if err := utils.CheckHash(user.Password, req.Password); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(uint(user.ID), string(s.JWTSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}
