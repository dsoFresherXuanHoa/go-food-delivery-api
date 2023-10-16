package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func RoleRouteConfig(router *gin.Engine) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	roles := router.Group("/api/v1/roles")
	{
		roles.POST("/", middlewares.RequiredManagerPermission(secretKey), controllers.CreateRole())
	}
}
