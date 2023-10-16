package main

import (
	"go-food-delivery-api/src/routes"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteConfig(db *gorm.DB) {
	port := os.Getenv("PORT")
	router := gin.Default()

	routes.RoleRouteConfig(router)
	routes.CategoriesRouteConfig(router)

	routes.AuthRouteConfig(router)
	routes.WarehouseRouteConfig(router)

	router.Run(":" + port)
}
