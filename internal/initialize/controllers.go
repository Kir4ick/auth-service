package initialize

import "github.com/auth-service/internal/http/controller/rest"

type Controllers struct {
	authController *rest.AuthController
	userController *rest.UserController
}

func NewControllers(services *Services) *Controllers {
	return &Controllers{
		authController: rest.NewAuthController(services.auth),
		userController: rest.NewUserController(),
	}
}
