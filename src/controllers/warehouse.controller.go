package controllers

import (
	"fmt"
	"go-food-delivery-api/src/constants"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/repositories"
	"go-food-delivery-api/src/services"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ImportWarehouse godoc
//
//	@Summary		Import Warehouse
//	@Description	Import Warehouse
//	@Tags			warehouses
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param  Authorization  header  string  required  "Bearer Token"
//	@Param			warehouse	formData		models.Warehouse	true	"Warehouse"
//	@Param			thumb	formData	file	true	"Thumb"
//	@Success		200	{object}		models.response
//	@Failure		400	{object}	models.response
//	@Failure		500	{object}	models.response
//	@Router			/warehouses [post]
func ImportWarehouse(db *gorm.DB) gin.HandlerFunc {
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

// UpdateWarehouse godoc
//
//	@Summary		Update Warehouse
//	@Description	Update Warehouse
//	@Tags			warehouses
//	@Accept			json
//	@Produce		json
//	@Param  Authorization  header  string  required  "Bearer Token"
//	@Param			warehouse	body		models.Warehouse	true	"Warehouse"
//	@Param			productId	query		int	true	"Product Id"
//	@Param			discountId	query		int	true	"DiscountId Id"
//	@Success		200	{object}		models.response
//	@Failure		400	{object}	models.response
//	@Failure		500	{object}	models.response
//	@Router			/warehouses [patch]
func UpdateWarehouse(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var warehouse models.WarehouseUpdatable
		if err := ctx.ShouldBind(&warehouse); err != nil {
			fmt.Println("Error while parse warehouse from form data: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidProductFormDataStruct))
		} else if productId, err := strconv.Atoi(ctx.Query("productId")); err != nil {
			fmt.Println("Error while get productId from query string: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidProductIDQueryString))
		} else if discountId, err := strconv.Atoi(ctx.Query("discountId")); err != nil {
			fmt.Println("Error while get discountId from query string: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidDiscountIDQueryString))
		} else {
			fmt.Println(productId, discountId)

			discount := models.DiscountUpdatable{MinQuantity: &warehouse.MinQuantity, DiscountPercent: &warehouse.DiscountPercent}
			product := models.ProductUpdatable{Name: &warehouse.Name, Description: &warehouse.Description, Price: &warehouse.Price, StockAmount: &warehouse.StockAmount}

			repositories := repositories.NewSQLStore(db)
			warehouseService := services.NewWarehouseBusiness(repositories)
			if discountId, productId, err := warehouseService.UpdateWareHouseById(ctx, discountId, &discount, productId, &product); err != nil {
				fmt.Println("Error while update warehouse in warehouse controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotUpdateWarehouse))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(gin.H{
					"discountId": discountId,
					"productId":  productId,
				}, http.StatusOK, "", constants.UpdateWarehouseSuccess))
			}
		}
	}
}
