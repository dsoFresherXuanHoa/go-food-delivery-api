package main

import (
	"fmt"
	"go-food-delivery-api/src/routes"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteConfig(db *gorm.DB) {
	port := os.Getenv("PORT")
	fmt.Println(port)
	router := gin.Default()

	routes.RoleRouteConfig(router)
	routes.CategoriesRouteConfig(router)
	routes.AreaRouteConfig(router)
	routes.TableRouteConfig(router)
	routes.ProductRouteConfig(router)

	routes.AuthRouteConfig(router)
	routes.WarehouseRouteConfig(router)

	router.Run(":" + port)
}
