package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func ProductRouteConfig(router *gin.Engine) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	products := router.Group("/api/v1/products")
	{
		products.GET("/", middlewares.RequiredWaiterPermissionOrMore(secretKey), controllers.ReadProduct())
	}
}
