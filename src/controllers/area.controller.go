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

func CreateArea(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var area models.AreaCreatable
		if err := ctx.ShouldBind(&area); err != nil {
			fmt.Println("Error while try parse request body to struct: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidRequestBody))
		} else {
			repositories := repositories.NewSQLStore(db)
			areaService := services.NewAreaBusiness(repositories)

			if areaId, err := areaService.CreateArea(ctx, &area); err != nil {
				fmt.Println("Error while create area in area controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotCreateArea))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(areaId, http.StatusOK, "", constants.CreateAreaSuccess))
			}
		}
	}
}
