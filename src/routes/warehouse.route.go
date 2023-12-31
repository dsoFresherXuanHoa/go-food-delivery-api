package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WarehouseRouteConfig(router *gin.Engine, db *gorm.DB) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	warehouses := router.Group("/api/v1/warehouses")
	{
		warehouses.POST("/", middlewares.RequiredManagerPermission(secretKey), controllers.ImportWarehouse(db))
		warehouses.PATCH("/", middlewares.RequiredManagerPermission(secretKey), controllers.UpdateWarehouse(db))
	}
}
