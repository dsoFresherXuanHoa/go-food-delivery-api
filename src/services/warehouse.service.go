package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
	"mime/multipart"
)

type WarehouseStorage interface {
	ImportWareHouse(ctx context.Context, discount *models.DiscountCreatable, product *models.ProductCreatable, productThumb *multipart.FileHeader) (*uint, *uint, error)
	UpdateWareHouseById(ctx context.Context, discountId int, discount *models.DiscountUpdatable, productId int, product *models.ProductUpdatable) (*uint, *int64, error)
}

type warehouseBusiness struct {
	storage WarehouseStorage
}

func NewWarehouseBusiness(storage WarehouseStorage) *warehouseBusiness {
	return &warehouseBusiness{storage: storage}
}

func (business *warehouseBusiness) ImportWareHouse(ctx context.Context, discount *models.DiscountCreatable, product *models.ProductCreatable, productThumb *multipart.FileHeader) (*uint, *uint, error) {
	if discountId, productId, err := business.storage.ImportWareHouse(ctx, discount, product, productThumb); err != nil {
		fmt.Println("Error while create warehouse in service: " + err.Error())
		return nil, nil, err
	} else {
		return discountId, productId, nil
	}
}

func (business *warehouseBusiness) UpdateWareHouseById(ctx context.Context, discountId int, discount *models.DiscountUpdatable, productId int, product *models.ProductUpdatable) (*uint, *int64, error) {
	if discountId, productId, err := business.storage.UpdateWareHouseById(ctx, discountId, discount, productId, product); err != nil {
		fmt.Println("Error while update warehouse in service: " + err.Error())
		return nil, nil, err
	} else {
		return discountId, productId, nil
	}
}
