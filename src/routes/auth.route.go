package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthRouteConfig(router *gin.Engine) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/sign-up", controllers.SignUp())
		auth.POST("/sign-in", controllers.SignIn())
		auth.GET("/me", middlewares.RequiredWaiterPermissionOrMore(secretKey), controllers.Me())
		auth.PATCH("/reset-password", middlewares.RequiredManagerPermission(secretKey), controllers.ResetPassword())
	}
}
