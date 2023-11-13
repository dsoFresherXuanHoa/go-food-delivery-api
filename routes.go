package main

import (
	"go-food-delivery-api/src/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteConfig(db *gorm.DB) {
	port := os.Getenv("PORT")
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router := gin.Default()
	router.Use(cors.New(config))

	routes.RoleRouteConfig(router)
	routes.CategoriesRouteConfig(router)
	routes.AreaRouteConfig(router)
	routes.TableRouteConfig(router)
	routes.ProductRouteConfig(router)
	routes.BillRouteConfig(router)

	routes.AuthRouteConfig(router)
	routes.WarehouseRouteConfig(router)
	routes.BookingRouteConfig(router)
	routes.StatsRouteConfig(router)

	router.Run(":" + port)
}
