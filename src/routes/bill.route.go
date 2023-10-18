package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func BillRouteConfig(router *gin.Engine) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	bills := router.Group("/api/v1/bills")
	{
		bills.PATCH("/finished/:billId", middlewares.RequiredChiefPermissionOrMore(secretKey), controllers.FinishBill())
	}
}
