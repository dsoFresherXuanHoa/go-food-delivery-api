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
		booking.POST("/", middlewares.RequiredWaiterPermissionOrMore(secretKey), controllers.CreateBooking())
	}
}
