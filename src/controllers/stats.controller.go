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

func ReadTopEmployeeByOrderNumber() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		repository := repositories.NewSQLStore(db)
		employeeService := services.NewEmployeeBusiness(repository)
		if startTime, err := strconv.Atoi(ctx.Query("startTime")); err != nil {
			fmt.Println("Error while read startTime in stats controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidTimeStatsQueryString))
		} else if endTime, err := strconv.Atoi(ctx.Query("endTime")); err != nil {
			fmt.Println("Error while read endTime in stats controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidTimeStatsQueryString))
		} else if res, err := employeeService.ReadTopEmployeeOrderNumber(ctx, startTime, endTime); err != nil {
			fmt.Println("Error while get employee by total number of order in stats controller: " + err.Error())
			ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotGetOrderNumberByEmployee))
		} else {
			ctx.JSON(http.StatusOK, models.NewStandardResponse(res, http.StatusOK, "", constants.GetOrderNumberByEmployeeSuccess))
		}
	}
}

func ReadTopProductByReorderLevel() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		repository := repositories.NewSQLStore(db)
		productService := services.NewProductBusiness(repository)
		if startTime, err := strconv.Atoi(ctx.Query("startTime")); err != nil {
			fmt.Println("Error while read startTime in stats controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidTimeStatsQueryString))
		} else if endTime, err := strconv.Atoi(ctx.Query("endTime")); err != nil {
			fmt.Println("Error while read endTime in stats controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidTimeStatsQueryString))
		} else if res, err := productService.ReadTopGoodsByReorderLevel(ctx, startTime, endTime); err != nil {
			fmt.Println("Error while get product by reorderLevel in stats controller: " + err.Error())
			ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotGetOrderNumberByEmployee))
		} else {
			ctx.JSON(http.StatusOK, models.NewStandardResponse(res, http.StatusOK, "", constants.GetOrderNumberByEmployeeSuccess))
		}
	}
}
