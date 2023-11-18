package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func EmployeeRouteConfig(router *gin.Engine) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	employees := router.Group("/api/v1/employees")
	{
		employees.GET("/", middlewares.RequiredManagerPermission(secretKey), controllers.ReadAllEmployee())
		employees.DELETE("/:employeeId", middlewares.RequiredManagerPermission(secretKey), controllers.DeleteEmployee())
	}
}
