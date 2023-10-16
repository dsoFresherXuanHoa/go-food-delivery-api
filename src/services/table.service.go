package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type TableStorage interface {
	CreateTable(ctx context.Context, table *models.TableCreatable) (*uint, error)
	ReadTableByEmployeeId(ctx context.Context, employeeId uint) (models.Tables, error)
}

type tableBusiness struct {
	storage TableStorage
}

func NewTableBusiness(storage TableStorage) *tableBusiness {
	return &tableBusiness{storage: storage}
}

func (business *tableBusiness) CreateTable(ctx context.Context, table *models.TableCreatable) (*uint, error) {
	if tableId, err := business.storage.CreateTable(ctx, table); err != nil {
		fmt.Println("Error while save Table in service: " + err.Error())
		return nil, err
	} else {
		return tableId, nil
	}
}

func (business *tableBusiness) ReadTableByEmployeeId(ctx context.Context, employeeId uint) (models.Tables, error) {
	if tables, err := business.storage.ReadTableByEmployeeId(ctx, employeeId); err != nil {
		fmt.Println("Error while find Table by employeeId in service: " + err.Error())
		return nil, err
	} else {
		return tables, nil
	}
}
