package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type ProductStorage interface {
	CreateProduct(ctx context.Context, product *models.ProductCreatable) (*uint, error)
	ReadProductByCategoryId(ctx context.Context, categoryId uint) ([]models.ProductResponse, error)
	ReadProduct(ctx context.Context) ([]models.ProductResponse, error)
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

func (business *productBusiness) ReadProduct(ctx context.Context) ([]models.ProductResponse, error) {
	if products, err := business.storage.ReadProduct(ctx); err != nil {
		fmt.Println("Error while read product in service: " + err.Error())
		return nil, err
	} else {
		return products, nil
	}
}

func (business *productBusiness) ReadProductByCategoryId(ctx context.Context, categoryId uint) ([]models.ProductResponse, error) {
	if products, err := business.storage.ReadProductByCategoryId(ctx, categoryId); err != nil {
		fmt.Println("Error while read product by categoryId in service: " + err.Error())
		return nil, err
	} else {
		return products, nil
	}
}
