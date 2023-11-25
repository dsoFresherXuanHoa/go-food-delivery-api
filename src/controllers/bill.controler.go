package controllers

import (
	"fmt"
	"go-food-delivery-api/src/constants"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/repositories"
	"go-food-delivery-api/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FinishBill godoc
//
//	@Summary		Finish Bill
//	@Description	Finish Bill
//	@Tags			bills
//	@Accept			json
//	@Produce		json
//	@Param  Authorization  header  string  required  "Bearer Token"
//	@Param			billId	path		int	true	"Bill Id"
//	@Success		200		{object}	models.response
//	@Failure		400		{object}	models.response
//	@Failure		500		{object}	models.response
//	@Router			/bills/finished/{billId} [patch]
func FinishBill(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if billId, err := strconv.Atoi(ctx.Param("billId")); err != nil {
			fmt.Println("Error while finish bill by id in bill controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidOrderIDQueryString))
		} else {
			repositories := repositories.NewSQLStore(db)
			billService := services.NewBillBusiness(repositories)
			if orderId, err := billService.FinishBillById(ctx, uint(billId)); err != nil {
				fmt.Println("Error while finish bill in bill controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotFinishBill))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(orderId, http.StatusOK, "", constants.FinishBillSuccess))
			}
		}
	}
}

// CompensatedBill godoc
//
//	@Summary		Compensated Bill
//	@Description	Compensated Bill
//	@Tags			bills
//	@Accept			json
//	@Produce		json
//	@Param  Authorization  header  string  required  "Bearer Token"
//	@Param			orderId	query		int	true	"Order Id"
//	@Param			billId	query		int	true	"Bill Id"
//	@Success		200		{object}	models.response
//	@Failure		400		{object}	models.response
//	@Failure		500		{object}	models.response
//	@Router			/bills/compensated/ [patch]
func CompensatedBill(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if orderId, err := strconv.Atoi(ctx.Query("orderId")); err != nil {
			fmt.Println("Error while get orderId from query string: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidOrderIDQueryString))
		} else if billId, err := strconv.Atoi(ctx.Query("billId")); err != nil {
			fmt.Println("Error while get billId from query string: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidBillIDQueryString))
		} else {
			repositories := repositories.NewSQLStore(db)
			billService := services.NewBillBusiness(repositories)
			if compensatedBillId, newBillId, err := billService.CompensatedBillById(ctx, uint(orderId), uint(billId)); err != nil {
				fmt.Println("Error while compensated a bill in bill controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotCompensatedBill))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(gin.H{
					"compensatedBillId": compensatedBillId,
					"newBillId":         newBillId,
				}, http.StatusOK, "", constants.CompensatedBillSuccess))
			}
		}
	}
}
