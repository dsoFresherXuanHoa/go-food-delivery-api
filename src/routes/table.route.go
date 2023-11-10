package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func TableRouteConfig(router *gin.Engine) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	tables := router.Group("/api/v1/tables")
	{
		tables.GET("/status/:status", middlewares.RequiredWaiterPermissionOrMore(secretKey), controllers.ReadTableByEmployeeIdAndStatus())
		tables.GET("/", middlewares.RequiredAuthorize(secretKey), controllers.ReadTableByEmployeeId())
		tables.POST("/", middlewares.RequiredManagerPermission(secretKey), controllers.CreateTable())
		tables.GET("/all", middlewares.RequiredAuthorize(secretKey), controllers.ReadAllTable())
	}
}
