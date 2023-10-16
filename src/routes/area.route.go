package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func AreaRouteConfig(router *gin.Engine) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	areas := router.Group("/api/v1/areas")
	{
		areas.POST("/", middlewares.RequiredManagerPermission(secretKey), controllers.CreateArea())
	}
}
