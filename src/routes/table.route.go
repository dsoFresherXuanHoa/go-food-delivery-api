package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TableRouteConfig(router *gin.Engine, db *gorm.DB) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	tables := router.Group("/api/v1/tables")
	{
		tables.GET("/status/:status", middlewares.RequiredWaiterPermissionOrMore(secretKey, db), controllers.ReadTableByEmployeeIdAndStatus(db))
		tables.GET("/", middlewares.RequiredAuthorize(secretKey), controllers.ReadTableByEmployeeId(db))
		tables.POST("/", middlewares.RequiredManagerPermission(secretKey), controllers.CreateTable(db))
		tables.GET("/all", middlewares.RequiredAuthorize(secretKey), controllers.ReadAllTable(db))
	}
}
