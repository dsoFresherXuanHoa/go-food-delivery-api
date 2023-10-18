package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func BookingRouteConfig(router *gin.Engine) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	booking := router.Group("/api/v1/booking")
	{
		booking.GET("/accepted", middlewares.RequiredChiefPermissionOrMore(secretKey), controllers.AcceptOrder())
		booking.GET("/rejected", middlewares.RequiredChiefPermissionOrMore(secretKey), controllers.RejectOrder())
		booking.GET("/finished", middlewares.RequiredChiefPermissionOrMore(secretKey), controllers.FinishOrder())
		booking.POST("/", middlewares.RequiredWaiterPermissionOrMore(secretKey), controllers.CreateBooking())
		booking.GET("/", middlewares.RequiredChiefPermissionOrMore(secretKey), controllers.GetDetailBooking())
	}
}
