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

func ReadProduct() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		repositories := repositories.NewSQLStore(db)
		categoryService := services.NewProductBusiness(repositories)

		if categories, err := categoryService.ReadProduct(ctx); err != nil {
			fmt.Println("Error while read product in category controller: " + err.Error())
			ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotReadProduct))
		} else {
			ctx.JSON(http.StatusOK, models.NewStandardResponse(categories, http.StatusOK, "", constants.ReadProductSuccess))
		}
	}
}

func ReadProductByCategoryId() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		if categoryId, err := strconv.Atoi(ctx.Param("categoryId")); err != nil {
			fmt.Println("Error while read product by categoryId in category controller: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidEmployeeIDQueryString))
		} else {
			repositories := repositories.NewSQLStore(db)
			categoryService := services.NewProductBusiness(repositories)
			fmt.Println(categoryId)
			if categories, err := categoryService.ReadProductByCategoryId(ctx, uint(categoryId)); err != nil {
				fmt.Println("Error while read product by category in category controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotReadProduct))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(categories, http.StatusOK, "", constants.ReadProductSuccess))
			}
		}
	}
}
