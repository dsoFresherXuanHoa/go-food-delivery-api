package routes

import (
	"go-food-delivery-api/src/controllers"
	"go-food-delivery-api/src/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookingRouteConfig(router *gin.Engine, db *gorm.DB) {
	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	booking := router.Group("/api/v1/booking")
	{
		booking.POST("/accepted", middlewares.RequiredChiefPermissionOrMore(secretKey), controllers.AcceptOrder(db))
		booking.POST("/rejected", middlewares.RequiredChiefPermissionOrMore(secretKey), controllers.RejectOrder(db))
		booking.PATCH("/finished", middlewares.RequiredWaiterPermissionOrMore(secretKey, db), controllers.FinishOrder(db))
		booking.POST("/", middlewares.RequiredWaiterPermissionOrMore(secretKey, db), controllers.CreateBooking(db))
		booking.GET("/", middlewares.RequiredAuthorize(secretKey), controllers.GetDetailBooking(db))
		booking.GET("/all", middlewares.RequiredAuthorize(secretKey), controllers.GetAllBooking(db))
		booking.GET("/employee", middlewares.RequiredWaiterPermissionOrMore(secretKey, db), controllers.GetTop10DetailBookingByEmployeeId(db))
		booking.GET("/table/:tableId", middlewares.RequiredAuthorize(secretKey), controllers.GetServeBookingByTableId(db))
		booking.GET("/table/preparing/:tableId", middlewares.RequiredAuthorize(secretKey), controllers.GetPreparingBookingByTableId(db))
		booking.GET("/table/rejected/:tableId", middlewares.RequiredAuthorize(secretKey), controllers.GetRejectedBookingByTableId(db))
		booking.PATCH("/refund/:orderId", middlewares.RequiredWaiterPermissionOrMore(secretKey, db), controllers.RefundBooking(db))
	}
}
