package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/configs"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/services"
	"mime/multipart"
	"os"
	"path/filepath"

	exceptions "go-food-delivery-api/src/errors"
)

func (s *sqlStorage) ImportWareHouse(ctx context.Context, discount *models.DiscountCreatable, product *models.ProductCreatable, productThumb *multipart.FileHeader) (*uint, *uint, error) {
	discountBusiness := services.NewDiscountBusiness(s)
	if discountId, err := discountBusiness.CreateDiscount(ctx, discount); err != nil {
		fmt.Println("Error while create discount in warehouse repository: " + err.Error())
		return nil, nil, err
	} else {
		cld, _ := configs.GetCloudinaryInstance()
		cloudinaryStorage := NewCloudinaryStore(cld)
		productBusiness := services.NewProductBusiness(s)
		currentPath, _ := os.Getwd()
		localStoragePath := filepath.Join(currentPath, "dist", "thumbnail")
		filePath := filepath.Join(localStoragePath, productThumb.Filename)
		fmt.Println(filePath)

		if res, _ := cloudinaryStorage.UploadProductThumb(ctx, filePath); res.Error.Message != "" {
			fmt.Println("Error while upload image to cloudinary: " + res.Error.Message)
			return nil, nil, exceptions.ErrSaveProductThumbnailsToCloudinaryStorage
		} else {
			fmt.Println("Discount Id: ", discountId)
			product.DiscountId = discountId
			product.Thumb = &res.URL
			if productId, err := productBusiness.CreateProduct(ctx, product); err != nil {
				fmt.Println("Error while create warehouse in repository: " + err.Error())
				return nil, nil, err
			} else {
				return discountId, productId, nil
			}
		}
	}
}
