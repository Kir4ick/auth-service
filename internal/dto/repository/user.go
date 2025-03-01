package repository

import "github.com/auth-service/internal/models"

type CreateUserInput struct {
	Login string
	Email string
	Hash  string
}

type CreateUserOutput struct {
	User *models.User
}
