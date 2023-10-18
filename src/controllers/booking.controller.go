package controllers

import (
	"fmt"
	"go-food-delivery-api/src/configs"
	"go-food-delivery-api/src/constants"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/repositories"
	"go-food-delivery-api/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBooking() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		var booking models.BookingCreatable
		if err := ctx.ShouldBind(&booking); err != nil {
			fmt.Println("Error while try parse request body to struct: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidRequestBody))
		} else {
			id := ctx.Value("employeeId").(int)
			employeeId := uint(id)
			order := models.OrderCreatable{Note: booking.Note, EmployeeId: &employeeId, TableId: booking.TableId}
			bills := make([]models.BillCreatable, len(booking.ProductsId))
			for i := 0; i < len(booking.ProductsId); i++ {
				bills[i] = models.BillCreatable{Quantity: &booking.Quantities[i], ProductId: &booking.ProductsId[i]}
			}

			repositories := repositories.NewSQLStore(db)
			bookingService := services.NewBookingBusiness(repositories)

			if orderId, err := bookingService.CreateBooking(ctx, &order, bills); err != nil {
				fmt.Println("Error while create booking in booking controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotCreateBooking))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(gin.H{
					"orderId": orderId,
				}, http.StatusOK, "", constants.CreateBookingSuccess))
			}
		}
	}
}

func AcceptOrder() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		if orderId, err := strconv.Atoi(ctx.Query("orderId")); err != nil {
			fmt.Println("Error while update order by id in category controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidOrderIDQueryString))
		} else {
			repositories := repositories.NewSQLStore(db)
			bookingService := services.NewBookingBusiness(repositories)
			if orderId, err := bookingService.AcceptOrder(ctx, orderId); err != nil {
				fmt.Println("Error while update order in booking controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotAcceptBooking))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(orderId, http.StatusOK, "", constants.AcceptBookingSuccess))
			}
		}
	}
}
