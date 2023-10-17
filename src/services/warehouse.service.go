package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
	"mime/multipart"
)

type WarehouseStorage interface {
	ImportWareHouse(ctx context.Context, discount *models.DiscountCreatable, product *models.ProductCreatable, productThumb *multipart.FileHeader) (*uint, *uint, error)
}

type warehouseBusiness struct {
	storage WarehouseStorage
}

func NewWarehouseBusiness(storage WarehouseStorage) *warehouseBusiness {
	return &warehouseBusiness{storage: storage}
}

func (business *warehouseBusiness) ImportWareHouse(ctx context.Context, discount *models.DiscountCreatable, product *models.ProductCreatable, productThumb *multipart.FileHeader) (*uint, *uint, error) {
	if discountId, productId, err := business.storage.ImportWareHouse(ctx, discount, product, productThumb); err != nil {
		fmt.Println("Error while create product in service: " + err.Error())
		return nil, nil, err
	} else {
		return discountId, productId, nil
	}
}
