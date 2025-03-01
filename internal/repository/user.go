package repository

import (
	"context"
	"github.com/auth-service/internal/dto/repository"
	"github.com/auth-service/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type UserRepository struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (r *UserRepository) Create(ctx context.Context, input repository.CreateUserInput) (*repository.CreateUserOutput, error) {
	var user models.User

	query := "insert into users (login, email, hash) values ($1, $2, $3) returning (id, uuid, login, email, hash, created_at, updated_at, deleted_at, email_verified_at);"

	err := r.conn.QueryRow(ctx, query, input.Login, input.Email, input.Hash).Scan(&user)
	if err != nil {
		return nil, errors.Wrap(err, "insert query for create user")
	}

	return &repository.CreateUserOutput{
		User: &user,
	}, nil
}
