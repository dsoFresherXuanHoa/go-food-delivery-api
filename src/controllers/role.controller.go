package controllers

import (
	"fmt"
	"go-food-delivery-api/src/constants"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/repositories"
	"go-food-delivery-api/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRole(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var role models.RoleCreatable
		if err := ctx.ShouldBind(&role); err != nil {
			fmt.Println("Error while try parse request body to struct: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidRequestBody))
		} else {
			repositories := repositories.NewSQLStore(db)
			roleService := services.NewRoleBusiness(repositories)

			if roleId, err := roleService.CreateRole(ctx, &role); err != nil {
				fmt.Println("Error while create role in auth controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotCreateRole))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(roleId, http.StatusOK, "", constants.CreateRoleSuccess))
			}
		}
	}
}
