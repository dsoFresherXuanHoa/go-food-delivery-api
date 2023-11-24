package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoriesRouteConfig(router *gin.Engine, db *gorm.DB) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	categories := router.Group("/api/v1/categories")
	{
		categories.GET("/", middlewares.RequiredAuthorize(secretKey), controllers.ReadCategory(db))
		categories.POST("/", middlewares.RequiredManagerPermission(secretKey), controllers.CreateCategory(db))
	}
}
