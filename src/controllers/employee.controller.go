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

// ReadAllEmployee godoc
//
//	@Summary		Read All Employee
//	@Description	Read All Employee
//	@Tags			employees
//	@Accept			json
//	@Produce		json
//	@Param  Authorization  header  string  required  "Bearer Token"
//	@Success		200		{object}	models.response
//	@Failure		400		{object}	models.response
//	@Failure		500		{object}	models.response
//	@Router			/employees [get]
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

// DeleteEmployee godoc
//
//	@Summary		Delete Employee
//	@Description	Delete Employee
//	@Tags			employees
//	@Accept			json
//	@Produce		json
//	@Param  Authorization  header  string  required  "Bearer Token"
//	@Param			employeeId	path		int	true	"Employee Id"
//	@Success		200		{object}	models.response
//	@Failure		400		{object}	models.response
//	@Failure		500		{object}	models.response
//	@Router			/employees/{employeesId} [delete]
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
