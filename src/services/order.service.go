package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type OrderStorage interface {
	CreateOrder(ctx context.Context, order models.OrderCreatable) (*uint, error)
	ReadOrderById(ctx context.Context, orderId uint) (*models.Order, error)
	UpdateOrderById(ctx context.Context, orderId int, order *models.OrderUpdatable) (*int64, error)
}

type orderBusiness struct {
	storage OrderStorage
}

func NewOrderBusiness(storage OrderStorage) *orderBusiness {
	return &orderBusiness{storage: storage}
}

func (business *orderBusiness) CreateOrder(ctx context.Context, order models.OrderCreatable) (*uint, error) {
	if id, err := business.storage.CreateOrder(ctx, order); err != nil {
		fmt.Println("Error while save order in service: " + err.Error())
		return nil, err
	} else {
		return id, nil
	}
}

func (business *orderBusiness) ReadOrderById(ctx context.Context, orderId uint) (*models.Order, error) {
	if id, err := business.storage.ReadOrderById(ctx, orderId); err != nil {
		fmt.Println("Error while find order by id in service: " + err.Error())
		return nil, err
	} else {
		return id, nil
	}
}

func (business *orderBusiness) UpdateOrderById(ctx context.Context, orderId int, order *models.OrderUpdatable) (*int64, error) {
	if rowsAffected, err := business.storage.UpdateOrderById(ctx, orderId, order); err != nil {
		fmt.Println("Error while update order by id in service: " + err.Error())
		return nil, err
	} else {
		return rowsAffected, nil
	}
}