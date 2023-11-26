package main

import (
	"go-food-delivery-api/docs"
	"go-food-delivery-api/src/configs"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/routes"
	"os"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Go Food Delivery CMS - Swagger API Discovery
//	@version		1.0
//	@description	Go Food Delivery CMS - Swagger API Discovery

//	@contact.name	Xuan Hoa Le
//	@contact.email	dso.intern.xuanhoa@gmail.com

// @host		localhost:3001
// @BasePath	/api/v1
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

		// limiter, _ := configs.NewRateLimitClient().Instance()
		router := gin.Default()
		router.Use(cors.AllowAll())
		// router.Use(middlewares.RateLimit(limiter))

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

		docs.SwaggerInfo.BasePath = "/api/v1"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		router.Run(":" + port)
	}
}
