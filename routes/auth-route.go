package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thoriqadillah/e-learning-api/controllers"
)

var authController controllers.AuthController = controllers.NewAuthController()

func (r *routes) authRoute(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")

	auth.POST("/login", authController.Login)
	auth.POST("/register", authController.Register)
}
