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
		booking.PATCH("/accepted", middlewares.RequiredChiefPermissionOrMore(secretKey), controllers.AcceptOrder())
		booking.PATCH("/rejected", middlewares.RequiredChiefPermissionOrMore(secretKey), controllers.RejectOrder())
		booking.PATCH("/finished", middlewares.RequiredChiefPermissionOrMore(secretKey), controllers.FinishOrder())
		booking.POST("/", middlewares.RequiredWaiterPermissionOrMore(secretKey), controllers.CreateBooking())
		booking.GET("/", middlewares.RequiredAuthorize(secretKey), controllers.GetDetailBooking())
		booking.PATCH("/compensated", middlewares.RequiredWaiterPermissionOrMore(secretKey), controllers.CompensatedOrder())
		booking.GET("/employee", middlewares.RequiredWaiterPermissionOrMore(secretKey), controllers.GetDetailBookingByEmployeeId())
		booking.GET("/table/:tableId", middlewares.RequiredWaiterPermissionOrMore(secretKey), controllers.GetServeBookingByTableId())
		booking.GET("/table/preparing/:tableId", middlewares.RequiredWaiterPermissionOrMore(secretKey), controllers.GetPreparingBookingByTableId())
	}
}
