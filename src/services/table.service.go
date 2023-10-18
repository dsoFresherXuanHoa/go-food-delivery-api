package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type TableStorage interface {
	CreateTable(ctx context.Context, table *models.TableCreatable) (*uint, error)
	ReadTableById(ctx context.Context, tableId uint) (*models.Table, error)
	ReadTableByEmployeeId(ctx context.Context, employeeId uint) ([]models.TableResponse, error)
	ReadTableByEmployeeIdAndStatus(ctx context.Context, employeeId uint, status bool) ([]models.TableResponse, error)
	UpdateTable(ctx context.Context, id int, table *models.TableUpdatable) (*int64, error)
	GetDetailTable(ctx context.Context, table models.Table) models.TableResponse
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

func (business *tableBusiness) ReadTableById(ctx context.Context, tableId uint) (*models.Table, error) {
	if table, err := business.storage.ReadTableById(ctx, tableId); err != nil {
		fmt.Println("Error while find Table by in service: " + err.Error())
		return nil, err
	} else {
		return table, nil
	}
}

func (business *tableBusiness) ReadTableByEmployeeId(ctx context.Context, employeeId uint) ([]models.TableResponse, error) {
	if tables, err := business.storage.ReadTableByEmployeeId(ctx, employeeId); err != nil {
		fmt.Println("Error while find Table by employeeId in service: " + err.Error())
		return nil, err
	} else {
		return tables, nil
	}
}

func (business *tableBusiness) ReadTableByEmployeeIdAndStatus(ctx context.Context, employeeId uint, status bool) ([]models.TableResponse, error) {
	if tables, err := business.storage.ReadTableByEmployeeIdAndStatus(ctx, employeeId, status); err != nil {
		fmt.Println("Error while find Table by employeeId and status in service: " + err.Error())
		return nil, err
	} else {
		return tables, nil
	}
}

func (business *tableBusiness) UpdateTable(ctx context.Context, id int, table *models.TableUpdatable) (*int64, error) {
	if tableId, err := business.storage.UpdateTable(ctx, id, table); err != nil {
		fmt.Println("Error while update table by id in service: " + err.Error())
		return nil, err
	} else {
		return tableId, nil
	}
}

func (business *tableBusiness) GetDetailTable(ctx context.Context, table models.Table) models.TableResponse {
	return business.storage.GetDetailTable(ctx, table)
}
