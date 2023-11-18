package main

import (
	"go-food-delivery-api/src/configs"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/routes"
	"os"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	if db, err := configs.GetGormInstance(); err != nil {
		panic("Can't connect to db via GORM: " + err.Error())
	} else {
		models := []interface{}{
			&models.Role{},
			&models.Employee{},
			&models.Account{},
			&models.Area{},
			&models.Table{},
			&models.Order{},
			&models.Categories{},
			&models.Discount{},
			&models.Product{},
			&models.Bill{},
			&models.ResetPassword{},
		}
		db.AutoMigrate(models...)

		port := os.Getenv("PORT")
		router := gin.Default()

		router.Use(cors.AllowAll())

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
		routes.EmployeeRouteConfig(router)

		router.Run(":" + port)
	}
}
