package service

import "github.com/auth-service/internal/models"

type GenerateTokenInput struct {
	Username string
	Password string
}

type GenerateTokenOutput struct {
	AccessToken  string
	RefreshToken string
}

type CreateUserInput struct {
	Login    string
	Email    string
	Password string
}

type CreateUserOutput struct {
	User *models.User
}
