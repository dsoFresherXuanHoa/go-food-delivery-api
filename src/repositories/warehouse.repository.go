package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/services"
)

func (s *sqlStorage) ImportWareHouse(ctx context.Context, discount *models.DiscountCreatable, product *models.ProductCreatable) (*uint, *uint, error) {
	repository := NewSQLStore(s.db)
	discountBusiness := services.NewDiscountBusiness(repository)
	productBusiness := services.NewProductBusiness(repository)
	if discountId, err := discountBusiness.CreateDiscount(ctx, discount); err != nil {
		fmt.Println("Error while create discount in warehouse repository: " + err.Error())
		return nil, nil, err
	} else {
		product.DiscountId = discountId
		if productId, err := productBusiness.CreateProduct(ctx, product); err != nil {
			fmt.Println("Error while create warehouse in repository: " + err.Error())
			return nil, nil, err
		} else {
			return discountId, productId, nil
		}
	}
}
