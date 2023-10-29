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

			if orderId, err := bookingService.CreateBooking(ctx, &order, bills, *booking.SecretCode); err != nil {
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

func RejectOrder() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		if orderId, err := strconv.Atoi(ctx.Query("orderId")); err != nil {
			fmt.Println("Error while update order by id in category controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidOrderIDQueryString))
		} else {
			repositories := repositories.NewSQLStore(db)
			bookingService := services.NewBookingBusiness(repositories)
			if orderId, err := bookingService.RejectOrder(ctx, orderId); err != nil {
				fmt.Println("Error while update order in booking controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotRejectBooking))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(orderId, http.StatusOK, "", constants.RejectBookingSuccess))
			}
		}
	}
}

func FinishOrder() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		if orderId, err := strconv.Atoi(ctx.Query("orderId")); err != nil {
			fmt.Println("Error while update order by id in category controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidOrderIDQueryString))
		} else {
			repositories := repositories.NewSQLStore(db)
			bookingService := services.NewBookingBusiness(repositories)
			if orderId, err := bookingService.FinishOrder(ctx, orderId); err != nil {
				fmt.Println("Error while update order in booking controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotRejectBooking))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(orderId, http.StatusOK, "", constants.RejectBookingSuccess))
			}
		}
	}
}

func GetDetailBooking() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		if orderId, err := strconv.Atoi(ctx.Query("orderId")); err != nil {
			fmt.Println("Error while update order by id in category controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidOrderIDQueryString))
		} else {
			repositories := repositories.NewSQLStore(db)
			bookingService := services.NewBookingBusiness(repositories)
			if booking, err := bookingService.GetDetailBooking(ctx, orderId); err != nil {
				fmt.Println("Error while update order in booking controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotRejectBooking))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(booking, http.StatusOK, "", constants.RejectBookingSuccess))
			}
		}
	}
}

func CompensatedOrder() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		if orderId, err := strconv.Atoi(ctx.Query("orderId")); err != nil {
			fmt.Println("Error while update order by id in category controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidOrderIDQueryString))
		} else {
			repositories := repositories.NewSQLStore(db)
			bookingService := services.NewBookingBusiness(repositories)
			id := ctx.Value("employeeId").(int)
			employeeId := uint(id)
			if orderId, err := bookingService.CompensatedOrder(ctx, orderId, int(employeeId)); err != nil {
				fmt.Println("Error while update order in booking controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotCompensateBooking))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(orderId, http.StatusOK, "", constants.CompensatedBookingSuccess))
			}
		}
	}
}

func GetDetailBookingByEmployeeId() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		id := ctx.Value("employeeId").(int)
		employeeId := uint(id)
		repositories := repositories.NewSQLStore(db)
		bookingService := services.NewBookingBusiness(repositories)
		if bookings, err := bookingService.GetOrdersByEmployeeId(ctx, int(employeeId)); err != nil {
			fmt.Println("Error while get all order by employeeId in booking controller: " + err.Error())
			ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotGetAllOrderByEmployeeId))
		} else {
			ctx.JSON(http.StatusOK, models.NewStandardResponse(bookings, http.StatusOK, "", constants.GetAllOrderByEmployeeIdSuccess))
		}
	}
}
