package controllers

import (
	"fmt"
	"go-food-delivery-api/src/configs"
	"go-food-delivery-api/src/constants"
	exceptions "go-food-delivery-api/src/errors"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/repositories"
	"go-food-delivery-api/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
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
			order := models.OrderCreatable{Note: booking.Note, EmployeeId: &employeeId, TableId: booking.TableId, Refundable: booking.Refundable}
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
		} else if reason := ctx.Query("reason"); reason == "" {
			fmt.Println("Error while get reason rejected order in query string: reason can not be empty")
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, exceptions.ErrReasonEmpty.Error(), constants.InvalidReasonQueryString))
		} else {
			repositories := repositories.NewSQLStore(db)
			bookingService := services.NewBookingBusiness(repositories)
			if orderId, err := bookingService.RejectOrder(ctx, orderId, reason); err != nil {
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

func GetTop10DetailBookingByEmployeeId() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		id := ctx.Value("employeeId").(int)
		employeeId := uint(id)
		repositories := repositories.NewSQLStore(db)
		bookingService := services.NewBookingBusiness(repositories)
		if bookings, err := bookingService.GetTop10OrdersByEmployeeId(ctx, int(employeeId)); err != nil {
			fmt.Println("Error while get all order by employeeId in booking controller: " + err.Error())
			ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotGetAllOrderByEmployeeId))
		} else {
			ctx.JSON(http.StatusOK, models.NewStandardResponse(bookings, http.StatusOK, "", constants.GetAllOrderByEmployeeIdSuccess))
		}
	}
}

func GetServeBookingByTableId() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		if tableId, err := strconv.Atoi(ctx.Param("tableId")); err != nil {
			fmt.Println("Error while get tableId in query string in booking controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidTableIdQueryString))
		} else {
			repositories := repositories.NewSQLStore(db)
			bookingService := services.NewBookingBusiness(repositories)
			if bookings, err := bookingService.GetServeBookingsByTableId(ctx, tableId); err != nil {
				fmt.Println("Error while get serve order by tableId in booking controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotGetServeOrderByTableId))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(bookings, http.StatusOK, "", constants.GetServeOrderByTableIdSuccess))
			}
		}
	}
}

func GetPreparingBookingByTableId() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		if tableId, err := strconv.Atoi(ctx.Param("tableId")); err != nil {
			fmt.Println("Error while get tableId in query string in booking controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidTableIdQueryString))
		} else {
			repositories := repositories.NewSQLStore(db)
			bookingService := services.NewBookingBusiness(repositories)
			if bookings, err := bookingService.GetPreparingBookingsByTableId(ctx, tableId); err != nil {
				fmt.Println("Error while get preparing order by tableId in booking controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotGetPreparingOrderByTableId))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(bookings, http.StatusOK, "", constants.GetPreparingOrderByTableIdSuccess))
			}
		}
	}
}

func RefundBooking() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	var booking models.BookingCreatable
	return func(ctx *gin.Context) {
		if orderId, err := strconv.Atoi(ctx.Param("orderId")); err != nil {
			fmt.Println("Error while get orderId in query string in booking controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidOrderIDQueryString))
		} else if err := ctx.ShouldBind(&booking); err != nil {
			fmt.Println("Error while try parse request body to struct: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidRequestBody))
		} else {
			repositories := repositories.NewSQLStore(db)
			bookingService := services.NewBookingBusiness(repositories)
			billService := services.NewBillBusiness(repositories)

			id := ctx.Value("employeeId").(int)
			employeeId := uint(id)
			order := models.OrderCreatable{Note: booking.Note, EmployeeId: &employeeId, TableId: booking.TableId, Refundable: booking.Refundable}
			order.EmployeeId = &employeeId
			order.Accepted = true

			if refundBills, err := billService.ReadBillsByOrderId(ctx, uint(orderId)); err != nil {
				fmt.Println("Error while find bills by order id before refund: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.RefundTargetNotFound))
			} else {
				bills := make([]models.BillCreatable, len(refundBills))
				refundProductId := make([]uint, len(refundBills))
				refundProductQuantity := make([]int, len(refundBills))
				requestProductId := make([]uint, len(booking.ProductsId))
				requestProductQuantity := make([]int, len(booking.ProductsId))
				for i := 0; i < len(refundBills); i++ {
					refundProductId[i] = refundBills[i].ProductId
					refundProductQuantity[i] = refundBills[i].Quantity
				}
				for i := 0; i < len(booking.ProductsId); i++ {
					requestProductId[i] = booking.ProductsId[i]
					requestProductQuantity[i] = booking.Quantities[i]
				}
				for i := 0; i < len(refundProductId); i++ {
					current := refundProductId[i]
					if index := slices.Index(requestProductId, current); index != -1 {
						quantity := refundProductQuantity[i] - requestProductQuantity[index]
						bills[i] = models.BillCreatable{Quantity: &quantity, ProductId: &refundBills[i].ProductId}
					} else {
						bills[i] = models.BillCreatable{Quantity: &refundBills[i].Quantity, ProductId: &refundBills[i].ProductId}
					}
				}
				if _, newOrderId, err := bookingService.RefundOrderById(ctx, orderId, &order, bills, *booking.SecretCode); err != nil {
					fmt.Println("Error while refund order by id in booking controller: " + err.Error())
					ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotRefundOrderByOrderId))
				} else {
					ctx.JSON(http.StatusOK, models.NewStandardResponse(gin.H{
						"refundOrderId": orderId,
						"newOrderId":    *newOrderId,
					}, http.StatusOK, "", constants.RefundOrderByOrderIdSuccess))
				}
			}
		}
	}
}
