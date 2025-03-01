package initialize

import "github.com/auth-service/internal/service"

type Services struct {
	auth *service.AuthService
}

func NewServices(repositories *Repositories) *Services {
	return &Services{
		auth: service.NewAuthService(repositories.userRepository),
	}
}
