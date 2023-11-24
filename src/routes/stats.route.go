package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StatsRouteConfig(router *gin.Engine, db *gorm.DB) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	stats := router.Group("/api/v1/stats")
	{
		stats.GET("/employee/top", middlewares.RequiredManagerPermission(secretKey), controllers.ReadTopEmployeeByOrderNumber(db))
		stats.GET("/goods/top", middlewares.RequiredManagerPermission(secretKey), controllers.ReadTopProductByReorderLevel(db))
	}
}
