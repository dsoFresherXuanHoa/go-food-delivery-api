package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type DiscountStorage interface {
	CreateDiscount(ctx context.Context, discount *models.DiscountCreatable) (*uint, error)
	ReadDiscountById(ctx context.Context, discountId uint) (*models.Discount, error)
	UpdateDiscountById(ctx context.Context, discountId uint, discount *models.DiscountUpdatable) (*uint, error)
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

func (business *discountBusiness) ReadDiscountById(ctx context.Context, discountId uint) (*models.Discount, error) {
	if discount, err := business.storage.ReadDiscountById(ctx, discountId); err != nil {
		fmt.Println("Error while read discount by id in service: " + err.Error())
		return nil, err
	} else {
		return discount, nil
	}
}

func (business *discountBusiness) UpdateDiscountById(ctx context.Context, discountId uint, discount *models.DiscountUpdatable) (*uint, error) {
	if updatedDiscountId, err := business.storage.UpdateDiscountById(ctx, discountId, discount); err != nil {
		fmt.Println("Error while update discount by id in service: " + err.Error())
		return nil, err
	} else {
		return updatedDiscountId, nil
	}
}
