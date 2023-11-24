package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRouteConfig(router *gin.Engine, db *gorm.DB) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	products := router.Group("/api/v1/products")
	{
		products.GET("/category/:categoryId", middlewares.RequiredAuthorize(secretKey), controllers.ReadProductByCategoryId(db))
		products.GET("/recommend/:limit", middlewares.RequiredWaiterPermissionOrMore(secretKey, db), controllers.ReadRecommendProduct(db))
		products.GET("/", middlewares.RequiredAuthorize(secretKey), controllers.ReadProduct(db))
		products.GET("/:productId", middlewares.RequiredAuthorize(secretKey), controllers.ReadProductById(db))
		products.DELETE("/:productId", middlewares.RequiredManagerPermission(secretKey), controllers.DeleteProductById(db))
	}
}
