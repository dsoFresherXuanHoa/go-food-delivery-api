package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type ProductStorage interface {
	CreateProduct(ctx context.Context, product *models.ProductCreatable) (*uint, error)
}

type productBusiness struct {
	storage ProductStorage
}

func NewProductBusiness(storage ProductStorage) *productBusiness {
	return &productBusiness{storage: storage}
}

func (business *productBusiness) CreateProduct(ctx context.Context, product *models.ProductCreatable) (*uint, error) {
	if productId, err := business.storage.CreateProduct(ctx, product); err != nil {
		fmt.Println("Error while create product in service: " + err.Error())
		return nil, err
	} else {
		return productId, nil
	}
}
