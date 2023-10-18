package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

func (s *sqlStorage) CreateOrder(ctx context.Context, order models.OrderCreatable) (*uint, error) {
	if err := s.db.Table(models.OrderCreatable{}.GetTableName()).Create(&order).Error; err != nil {
		fmt.Println("Error while create order in repository: " + err.Error())
		return nil, err
	}
	return &order.ID, nil
}

func (s *sqlStorage) ReadOrderById(ctx context.Context, orderId uint) (*models.Order, error) {
	var order models.Order
	if err := s.db.Where("id = ?", orderId).First(&order).Error; err != nil {
		fmt.Println("Error while find order by id in repository: " + err.Error())
		return nil, err
	}
	return &order, nil
}
