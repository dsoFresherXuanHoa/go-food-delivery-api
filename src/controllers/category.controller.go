package controllers

import (
	"fmt"
	"go-food-delivery-api/src/configs"
	"go-food-delivery-api/src/constants"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/repositories"
	"go-food-delivery-api/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		var category models.CategoryCreatable
		if err := ctx.ShouldBind(&category); err != nil {
			fmt.Println("Error while try parse request body to struct: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidRequestBody))
		} else {
			repositories := repositories.NewSQLStore(db)
			categoryService := services.NewCategoryBusiness(repositories)

			if categoryId, err := categoryService.CreateCategory(ctx, &category); err != nil {
				fmt.Println("Error while create category in category controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotCreateCategory))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(categoryId, http.StatusOK, "", constants.CreateCategorySuccess))
			}
		}
	}
}

func ReadCategory() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		repositories := repositories.NewSQLStore(db)
		categoryService := services.NewCategoryBusiness(repositories)

		if categories, err := categoryService.ReadCategory(ctx); err != nil {
			fmt.Println("Error while read category in category controller: " + err.Error())
			ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotReadCategory))
		} else {
			ctx.JSON(http.StatusOK, models.NewStandardResponse(categories, http.StatusOK, "", constants.ReadCategorySuccess))
		}
	}
}
