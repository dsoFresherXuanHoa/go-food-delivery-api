package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func CategoriesRouteConfig(router *gin.Engine) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	categories := router.Group("/api/v1/categories")
	{
		categories.GET("/", middlewares.RequiredAuthorize(secretKey), controllers.ReadCategory())
		categories.POST("/", middlewares.RequiredManagerPermission(secretKey), controllers.CreateCategory())
	}
}
