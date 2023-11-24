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

func ReadAllEmployee(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		repositories := repositories.NewSQLStore(db)
		employeeService := services.NewEmployeeBusiness(repositories)

		if employees, err := employeeService.ReadAllEmployee(ctx); err != nil {
			fmt.Println("Error while read all employee in table controller: " + err.Error())
			ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotReadAllEmployee))
		} else {
			ctx.JSON(http.StatusOK, models.NewStandardResponse(employees, http.StatusOK, "", constants.ReadAllEmployeeSuccess))
		}
	}
}

func DeleteEmployee(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if employeeId, err := strconv.Atoi(ctx.Param("employeeId")); err != nil {
			fmt.Println("Error while read employee by id in employee controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidEmployeeIDQueryString))
		} else {
			repositories := repositories.NewSQLStore(db)
			employeeService := services.NewEmployeeBusiness(repositories)
			if deleteEmployeeId, err := employeeService.DeleteEmployeeById(ctx, employeeId); err != nil {
				fmt.Println("Error while delete employee by id in employee controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotDeleteEmployeeByID))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(deleteEmployeeId, http.StatusOK, "", constants.DeleteEmployeeByIDSuccess))
			}
		}
	}
}
