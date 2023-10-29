package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

func (s *sqlStorage) GetUpdatableOrder(ctx context.Context, order models.Order) models.OrderUpdatable {
	return models.OrderUpdatable{Model: order.Model, Note: &order.Note, Status: &order.Status, Accepted: &order.Accepted, Rejected: &order.Rejected, Compensate: &order.Compensate}
}

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

func (s *sqlStorage) ReadOrdersByEmployeeId(ctx context.Context, employeeId uint) ([]models.Order, error) {
	var orders []models.Order
	if err := s.db.Where("employee_id = ?", employeeId).Find(&orders).Error; err != nil {
		fmt.Println("Error while find order by employeeId in repository: " + err.Error())
		return nil, err
	}
	return orders, nil
}

func (s *sqlStorage) GetServeOrdersByTableId(ctx context.Context, tableId int) ([]models.Order, error) {
	var orders []models.Order
	if err := s.db.Where("table_id = ?", tableId).Where("status = ?", false).Where("accepted = ?", true).Find(&orders).Error; err != nil {
		fmt.Println("Error while find serve order by tableId in repository: " + err.Error())
		return nil, err
	}
	return orders, nil
}

func (s *sqlStorage) UpdateOrderById(ctx context.Context, orderId int, order *models.OrderUpdatable) (*int64, error) {
	if result := s.db.Table(models.OrderUpdatable{}.GetTableName()).Where("id = ?", orderId).Updates(order); result.Error != nil {
		fmt.Println("Error while update order in repository: " + result.Error.Error())
		return nil, result.Error
	} else {
		return &result.RowsAffected, nil
	}
}
