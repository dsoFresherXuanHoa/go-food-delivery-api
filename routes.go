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

	routes.AuthRouteConfig(router)
	routes.RoleRouteConfig(router)

	router.Run(":" + port)
}
