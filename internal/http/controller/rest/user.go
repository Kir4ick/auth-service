package rest

import "github.com/gin-gonic/gin"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) ForgotPassword(ctx *gin.Context) {

}

func (c *UserController) ResetPassword(ctx *gin.Context) {

}

func (c *UserController) Me(ctx *gin.Context) {

}
