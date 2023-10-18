package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type BillStorage interface {
	CreateBill(ctx context.Context, bill *models.BillCreatable) (*uint, error)
	ReadBillById(ctx context.Context, billId uint) (*models.Bill, error)
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
