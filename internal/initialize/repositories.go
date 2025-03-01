package initialize

import (
	"github.com/auth-service/internal/repository"
	"github.com/jackc/pgx/v5"
)

type Repositories struct {
	userRepository *repository.UserRepository
}

func NewRepositories(conn *pgx.Conn) *Repositories {
	return &Repositories{
		userRepository: repository.NewUserRepository(conn),
	}
}
