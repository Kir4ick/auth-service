package initialize

import "github.com/gin-gonic/gin"

type Routes struct {
	controllers *Controllers
}

func NewRoutes(handlers *Controllers) *Routes {
	return &Routes{handlers}
}

func (r *Routes) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/sign-in", r.controllers.authController.SignIn)
	router.POST("/sign-up", r.controllers.authController.SignUp)

	router.POST("/logout", r.controllers.authController.Logout)
	router.POST("/refresh", r.controllers.authController.Refresh)

	router.POST("/forgot-password", r.controllers.userController.ForgotPassword)
	router.POST("/reset-password", r.controllers.userController.ResetPassword)
	router.GET("/me", r.controllers.userController.Me)

	return router
}
