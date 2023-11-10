package controllers

import (
	"fmt"
	"go-food-delivery-api/src/configs"
	"go-food-delivery-api/src/constants"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/repositories"
	"go-food-delivery-api/src/services"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func ImportWarehouse() gin.HandlerFunc {
	db, _ := configs.GetGormInstance()
	return func(ctx *gin.Context) {
		var warehouse models.Warehouse
		if err := ctx.ShouldBind(&warehouse); err != nil {
			fmt.Println("Error while parse product from form data: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidProductFormDataStruct))
		} else if uploadFile, err := ctx.FormFile("thumb"); err != nil {
			fmt.Println("Error while parse product from form data: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidProductFormFile))
		} else {
			currentPath, _ := os.Getwd()
			localStorageFilePath := filepath.Join(currentPath, "dist", "thumbnail", uploadFile.Filename)
			if err := ctx.SaveUploadedFile(uploadFile, localStorageFilePath); err != nil {
				fmt.Println("Can't upload image to local storage: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotSaveProductThumbnailsToLocalStorage))
			} else {
				discount := models.DiscountCreatable{MinQuantity: &warehouse.MinQuantity, DiscountPercent: &warehouse.DiscountPercent}
				product := models.ProductCreatable{Name: &warehouse.Name, Description: &warehouse.Description, Price: warehouse.Price, Thumb: &warehouse.Thumb, StockAmount: &warehouse.StockAmount, CategoryId: &warehouse.CategoryId, Uint: &warehouse.Uint}
				repositories := repositories.NewSQLStore(db)
				warehouseService := services.NewWarehouseBusiness(repositories)
				if discountId, productId, err := warehouseService.ImportWareHouse(ctx, &discount, &product, uploadFile); err != nil {
					fmt.Println("Error while create warehouse in warehouse controller: " + err.Error())
					ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotImportWarehouse))
				} else {
					ctx.JSON(http.StatusOK, models.NewStandardResponse(gin.H{
						"discountId": discountId,
						"productId":  productId,
					}, http.StatusOK, "", constants.ImportWarehouseSuccess))
				}
			}
		}
	}
}
