package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type BookingStorage interface {
	CreateBooking(ctx context.Context, order *models.OrderCreatable, bills []models.BillCreatable, secretCode int) (*uint, error)
	AcceptOrder(ctx context.Context, orderId int) (*uint, error)
	RejectOrder(ctx context.Context, orderId int) (*uint, error)
	FinishOrder(ctx context.Context, orderId int) (*uint, error)
	GetDetailBooking(ctx context.Context, orderId int) (*models.BookingResponse, error)
	CompensatedOrder(ctx context.Context, orderId int, employeeId int) (*uint, error)
	GetOrdersByEmployeeId(ctx context.Context, employeeId int) ([]models.BookingResponse, error)
	GetServeBookingsByTableId(ctx context.Context, tableId int) ([]models.BookingResponse, error)
	GetPreparingBookingsByTableId(ctx context.Context, tableId int) ([]models.BookingResponse, error)
}

type bookingBusiness struct {
	storage BookingStorage
}

func NewBookingBusiness(storage BookingStorage) *bookingBusiness {
	return &bookingBusiness{storage: storage}
}

func (business *bookingBusiness) CreateBooking(ctx context.Context, order *models.OrderCreatable, bills []models.BillCreatable, secretCode int) (*uint, error) {
	if orderId, err := business.storage.CreateBooking(ctx, order, bills, secretCode); err != nil {
		fmt.Println("Error while create order in service: " + err.Error())
		return nil, err
	} else {
		return orderId, nil
	}
}

func (business *bookingBusiness) AcceptOrder(ctx context.Context, orderId int) (*uint, error) {
	if orderId, err := business.storage.AcceptOrder(ctx, orderId); err != nil {
		fmt.Println("Error while update order in service: " + err.Error())
		return nil, err
	} else {
		return orderId, nil
	}
}

func (business *bookingBusiness) RejectOrder(ctx context.Context, orderId int) (*uint, error) {
	if orderId, err := business.storage.RejectOrder(ctx, orderId); err != nil {
		fmt.Println("Error while update order in service: " + err.Error())
		return nil, err
	} else {
		return orderId, nil
	}
}

func (business *bookingBusiness) FinishOrder(ctx context.Context, orderId int) (*uint, error) {
	if orderId, err := business.storage.FinishOrder(ctx, orderId); err != nil {
		fmt.Println("Error while update order in service: " + err.Error())
		return nil, err
	} else {
		return orderId, nil
	}
}

func (business *bookingBusiness) GetDetailBooking(ctx context.Context, orderId int) (*models.BookingResponse, error) {
	if booking, err := business.storage.GetDetailBooking(ctx, orderId); err != nil {
		fmt.Println("Error while read booking in service: " + err.Error())
		return nil, err
	} else {
		return booking, nil
	}
}

func (business *bookingBusiness) CompensatedOrder(ctx context.Context, orderId int, employeeId int) (*uint, error) {
	if orderId, err := business.storage.CompensatedOrder(ctx, orderId, employeeId); err != nil {
		fmt.Println("Error while create compensated booking in service: " + err.Error())
		return nil, err
	} else {
		return orderId, nil
	}
}

func (business *bookingBusiness) GetOrdersByEmployeeId(ctx context.Context, employeeId int) ([]models.BookingResponse, error) {
	if orders, err := business.storage.GetOrdersByEmployeeId(ctx, employeeId); err != nil {
		fmt.Println("Error while get all booking by employeeId in service: " + err.Error())
		return nil, err
	} else {
		return orders, nil
	}
}

func (business *bookingBusiness) GetServeBookingsByTableId(ctx context.Context, tableId int) ([]models.BookingResponse, error) {
	if orders, err := business.storage.GetServeBookingsByTableId(ctx, tableId); err != nil {
		fmt.Println("Error while get serve booking by tableId in service: " + err.Error())
		return nil, err
	} else {
		return orders, nil
	}
}

func (business *bookingBusiness) GetPreparingBookingsByTableId(ctx context.Context, tableId int) ([]models.BookingResponse, error) {
	if orders, err := business.storage.GetPreparingBookingsByTableId(ctx, tableId); err != nil {
		fmt.Println("Error while get preparing booking by tableId in service: " + err.Error())
		return nil, err
	} else {
		return orders, nil
	}
}
