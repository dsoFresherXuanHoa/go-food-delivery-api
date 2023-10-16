package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type CategoryStorage interface {
	CreateCategory(ctx context.Context, category *models.CategoryCreatable) (*uint, error)
}

type categoryBusiness struct {
	storage CategoryStorage
}

func NewCategoryBusiness(storage CategoryStorage) *categoryBusiness {
	return &categoryBusiness{storage: storage}
}

func (business *categoryBusiness) CreateCategory(ctx context.Context, category *models.CategoryCreatable) (*uint, error) {
	if categoryId, err := business.storage.CreateCategory(ctx, category); err != nil {
		fmt.Println("Error while create category in service: " + err.Error())
		return nil, err
	} else {
		return categoryId, nil
	}
}
