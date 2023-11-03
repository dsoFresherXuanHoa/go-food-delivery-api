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
		products.GET("/category/:categoryId", middlewares.RequiredAuthorize(secretKey), controllers.ReadProductByCategoryId())
		products.GET("/recommend/:limit", middlewares.RequiredWaiterPermissionOrMore(secretKey), controllers.ReadRecommendProduct())
		products.GET("/", middlewares.RequiredAuthorize(secretKey), controllers.ReadProduct())
		products.GET("/:productId", middlewares.RequiredAuthorize(secretKey), controllers.ReadProductById())
	}
}
