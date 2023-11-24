package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouteConfig(router *gin.Engine, db *gorm.DB) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/sign-up", controllers.SignUp(db))
		auth.POST("/sign-in", controllers.SignIn(db))
		auth.GET("/me", middlewares.RequiredAuthorize(secretKey), controllers.Me(db))
		auth.PATCH("/reset-password", middlewares.RequiredManagerPermission(secretKey), controllers.ResetPassword(db))
	}
}
