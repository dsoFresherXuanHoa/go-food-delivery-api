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

func CreateTable() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		var table models.TableCreatable
		if err := ctx.ShouldBind(&table); err != nil {
			fmt.Println("Error while try parse request body to struct: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidRequestBody))
		} else {
			repositories := repositories.NewSQLStore(db)
			tableService := services.NewTableBusiness(repositories)

			if tableId, err := tableService.CreateTable(ctx, &table); err != nil {
				fmt.Println("Error while create table in table controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotCreateTable))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(tableId, http.StatusOK, "", constants.CreateTableSuccess))
			}
		}
	}
}

func ReadTableByEmployeeId() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		repositories := repositories.NewSQLStore(db)
		tableService := services.NewTableBusiness(repositories)

		employeeId := ctx.Value("employeeId").(int)
		if tables, err := tableService.ReadTableByEmployeeId(ctx, uint(employeeId)); err != nil {
			fmt.Println("Error while read table by employeeId in table controller: " + err.Error())
			ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotReadTableByEmployeeId))
		} else {
			ctx.JSON(http.StatusOK, models.NewStandardResponse(tables, http.StatusOK, "", constants.ReadTableByEmployeeIdSuccess))
		}
	}
}

func ReadTableByEmployeeIdAndStatus() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		if status, err := strconv.ParseBool(ctx.Param("status")); err != nil {
			fmt.Println("Error while read product by status in controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidTableStatusQueryString))
		} else {
			repositories := repositories.NewSQLStore(db)
			tableService := services.NewTableBusiness(repositories)

			employeeId := ctx.Value("employeeId").(int)
			if tables, err := tableService.ReadTableByEmployeeIdAndStatus(ctx, uint(employeeId), status); err != nil {
				fmt.Println("Error while read table by employeeId in table controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotReadTableByEmployeeIdAndStatus))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(tables, http.StatusOK, "", constants.ReadTableByEmployeeIdAndStatusSuccess))
			}
		}
	}
}

func ReadAllTable() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		repositories := repositories.NewSQLStore(db)
		tableService := services.NewTableBusiness(repositories)

		if tables, err := tableService.ReadAllTable(ctx); err != nil {
			fmt.Println("Error while read all table in table controller: " + err.Error())
			ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotReadAllTable))
		} else {
			ctx.JSON(http.StatusOK, models.NewStandardResponse(tables, http.StatusOK, "", constants.CannotReadAllTableSuccess))
		}
	}
}
