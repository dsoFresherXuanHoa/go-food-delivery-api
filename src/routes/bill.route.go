package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BillRouteConfig(router *gin.Engine, db *gorm.DB) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	bills := router.Group("/api/v1/bills")
	{
		bills.POST("/finished/:billId", middlewares.RequiredChiefPermissionOrMore(secretKey), controllers.FinishBill(db))
		bills.PATCH("/compensated", middlewares.RequiredChiefPermissionOrMore(secretKey), controllers.CompensatedBill(db))
	}
}
