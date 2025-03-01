package service

import (
	"context"
	"github.com/auth-service/internal/dto/repository"
	"github.com/auth-service/internal/dto/service"
	"github.com/auth-service/pkg"
	"github.com/pkg/errors"
)

type UserRepository interface {
	Create(context.Context, repository.CreateUserInput) (*repository.CreateUserOutput, error)
}

type AuthService struct {
	userRepository UserRepository
}

func NewAuthService(userRepository UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) GenerateToken(_ context.Context, _ service.GenerateTokenInput) (*service.GenerateTokenOutput, error) {
	return &service.GenerateTokenOutput{}, nil
}

func (s *AuthService) CreateUser(ctx context.Context, input service.CreateUserInput) (*service.CreateUserOutput, error) {
	password, err := pkg.HashPassword(input.Password)
	if err != nil {
		return nil, errors.Wrap(err, "hashing password")
	}

	createResult, err := s.userRepository.Create(ctx, repository.CreateUserInput{
		Login: input.Login,
		Email: input.Email,
		Hash:  password,
	})
	if err != nil {
		return nil, errors.Wrap(err, "creating user")
	}

	return &service.CreateUserOutput{
		User: createResult.User,
	}, nil
}
