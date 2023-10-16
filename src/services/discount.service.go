package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type DiscountStorage interface {
	CreateDiscount(ctx context.Context, discount *models.DiscountCreatable) (*uint, error)
}

type discountBusiness struct {
	storage DiscountStorage
}

func NewDiscountBusiness(storage DiscountStorage) *discountBusiness {
	return &discountBusiness{storage: storage}
}

func (business *discountBusiness) CreateDiscount(ctx context.Context, discount *models.DiscountCreatable) (*uint, error) {
	if discountId, err := business.storage.CreateDiscount(ctx, discount); err != nil {
		fmt.Println("Error while create discount in service: " + err.Error())
		return nil, err
	} else {
		return discountId, nil
	}
}
