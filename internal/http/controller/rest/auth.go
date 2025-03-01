package rest

import (
	"context"
	"github.com/auth-service/internal/dto/service"
	"github.com/auth-service/internal/http/request"
	"github.com/auth-service/internal/http/response"
	"github.com/auth-service/pkg"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type authService interface {
	GenerateToken(context.Context, service.GenerateTokenInput) (*service.GenerateTokenOutput, error)
	CreateUser(context.Context, service.CreateUserInput) (*service.CreateUserOutput, error)
}

type AuthController struct {
	service authService
}

func NewAuthController(service authService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (c *AuthController) SignIn(ctx *gin.Context) {
	username, pass, ok := ctx.Request.BasicAuth()
	if !ok {
		pkg.SendErrorResponse(http.StatusUnauthorized, errors.New("Unauthorized"), ctx)
		return
	}

	input := service.GenerateTokenInput{
		Username: username,
		Password: pass,
	}

	result, err := c.service.GenerateToken(ctx, input)
	if err != nil {
		pkg.SendErrorResponseByError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, response.SignInResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	})
}

func (c *AuthController) SignUp(ctx *gin.Context) {
	var req request.SignUpRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		pkg.SendErrorResponse(http.StatusBadRequest, err, ctx)
		return
	}

	result, err := c.service.CreateUser(ctx, service.CreateUserInput{
		Login:    req.Login,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		pkg.SendErrorResponseByError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, response.SignUpResponse{
		response.User{
			Login:           result.User.Login,
			Email:           result.User.Email,
			Id:              result.User.Id,
			Uuid:            result.User.Uuid,
			CreatedAt:       result.User.CreatedAt,
			UpdatedAt:       result.User.UpdatedAt,
			DeletedAt:       result.User.DeletedAt,
			EmailVerifiedAt: result.User.EmailVerifiedAt,
		},
	})
}

func (c *AuthController) Logout(ctx *gin.Context) {

}

func (c *AuthController) Refresh(ctx *gin.Context) {

}
