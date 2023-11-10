package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type BillStorage interface {
	CreateBill(ctx context.Context, bill *models.BillCreatable) (*uint, error)
	ReadBillById(ctx context.Context, billId uint) (*models.Bill, error)
	ReadBillsByOrderId(ctx context.Context, orderId uint) ([]models.Bill, error)
	UpdateBillById(ctx context.Context, billId int, bill *models.BillUpdatable) (*int64, error)
	FinishBillById(ctx context.Context, billId uint) (*uint, error)
	CompensatedBillById(ctx context.Context, orderId uint, billId uint) (*uint, *uint, error)
}

type billBusiness struct {
	storage BillStorage
}

func NewBillBusiness(storage BillStorage) *billBusiness {
	return &billBusiness{storage: storage}
}

func (business *billBusiness) CreateBill(ctx context.Context, bill *models.BillCreatable) (*uint, error) {
	if id, err := business.storage.CreateBill(ctx, bill); err != nil {
		fmt.Println("Error while save bill in service: " + err.Error())
		return nil, err
	} else {
		return id, nil
	}
}

func (business *billBusiness) ReadBillById(ctx context.Context, billId uint) (*models.Bill, error) {
	if id, err := business.storage.ReadBillById(ctx, billId); err != nil {
		fmt.Println("Error while find bill by id in service: " + err.Error())
		return nil, err
	} else {
		return id, nil
	}
}

func (business *billBusiness) ReadBillsByOrderId(ctx context.Context, orderId uint) ([]models.Bill, error) {
	if id, err := business.storage.ReadBillsByOrderId(ctx, orderId); err != nil {
		fmt.Println("Error while find bill by order id in service: " + err.Error())
		return nil, err
	} else {
		return id, nil
	}
}

func (business *billBusiness) UpdateBillById(ctx context.Context, billId int, bill *models.BillUpdatable) (*int64, error) {
	if id, err := business.storage.UpdateBillById(ctx, billId, bill); err != nil {
		fmt.Println("Error while update bill by order id in service: " + err.Error())
		return nil, err
	} else {
		return id, nil
	}
}

func (business *billBusiness) FinishBillById(ctx context.Context, billId uint) (*uint, error) {
	if id, err := business.storage.FinishBillById(ctx, billId); err != nil {
		fmt.Println("Error while finish bill by order id in service: " + err.Error())
		return nil, err
	} else {
		return id, nil
	}
}

func (business *billBusiness) CompensatedBillById(ctx context.Context, orderId uint, billId uint) (*uint, *uint, error) {
	if compensatedBillId, newBillId, err := business.storage.CompensatedBillById(ctx, orderId, billId); err != nil {
		fmt.Println("Error while compensated bill by order id in service: " + err.Error())
		return nil, nil, err
	} else {
		return compensatedBillId, newBillId, nil
	}
}
