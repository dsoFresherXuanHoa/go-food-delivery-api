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

		routes.RoleRouteConfig(router, db)
		routes.CategoriesRouteConfig(router, db)
		routes.AreaRouteConfig(router, db)
		routes.TableRouteConfig(router, db)
		routes.ProductRouteConfig(router, db)
		routes.BillRouteConfig(router, db)
		routes.AuthRouteConfig(router, db)
		routes.WarehouseRouteConfig(router, db)
		routes.BookingRouteConfig(router, db)
		routes.StatsRouteConfig(router, db)
		routes.EmployeeRouteConfig(router, db)

		router.Run(":" + port)
	}
}
