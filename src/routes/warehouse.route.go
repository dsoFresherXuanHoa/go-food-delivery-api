package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func WarehouseRouteConfig(router *gin.Engine) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	warehouses := router.Group("/api/v1/warehouses")
	{
		warehouses.POST("/", middlewares.RequiredManagerPermission(secretKey), controllers.ImportWarehouse())
		warehouses.PATCH("/", middlewares.RequiredManagerPermission(secretKey), controllers.UpdateWarehouse())
	}
}
