package routes

import (
	"go-food-delivery-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouteConfig(router *gin.Engine) {
	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/sign-up", controllers.SignUp())
		auth.POST("/sign-in", controllers.SignIn())
	}
}
