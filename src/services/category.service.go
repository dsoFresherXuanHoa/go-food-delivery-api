package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type CategoryStorage interface {
	CreateCategory(ctx context.Context, category *models.CategoryCreatable) (*uint, error)
	ReadCategory(ctx context.Context) (models.Categories, error)
	ReadCategoryById(ctx context.Context, categoryId uint) (*models.Category, error)
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

func (business *categoryBusiness) ReadCategory(ctx context.Context) (models.Categories, error) {
	if categories, err := business.storage.ReadCategory(ctx); err != nil {
		fmt.Println("Error while read category in service: " + err.Error())
		return nil, err
	} else {
		return categories, nil
	}
}

func (business *categoryBusiness) ReadCategoryById(ctx context.Context, categoryId uint) (*models.Category, error) {
	if category, err := business.storage.ReadCategoryById(ctx, categoryId); err != nil {
		fmt.Println("Error while read category by id in service: " + err.Error())
		return nil, err
	} else {
		return category, nil
	}
}
