package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/services"
)

func (s *sqlStorage) CreateTable(ctx context.Context, table *models.TableCreatable) (*uint, error) {
	if err := s.db.Table(models.TableCreatable{}.GetTableName()).Create(&table).Error; err != nil {
		fmt.Println("Error while create table in repository: " + err.Error())
		return nil, err
	} else {
		return &table.ID, nil
	}
}

func (s *sqlStorage) ReadTableById(ctx context.Context, tableId uint) (*models.Table, error) {
	var table models.Table
	if err := s.db.Where("id = ?", tableId).Find(&table).Error; err != nil {
		fmt.Println("Error while find table by id in repository: " + err.Error())
		return nil, err
	}
	return &table, nil
}

func (s *sqlStorage) ReadTableByEmployeeId(ctx context.Context, employeeId uint) ([]models.TableResponse, error) {
	var tables models.Tables
	if err := s.db.Where("employee_id = ?", employeeId).Find(&tables).Error; err != nil {
		fmt.Println("Error while find table by id in repository: " + err.Error())
		return nil, err
	} else if len(tables) == 0 {
		fmt.Println("No record found while find table by employeeId in repository: " + err.Error())
		return nil, err
	}

	var res = make([]models.TableResponse, len(tables))
	for i, table := range tables {
		res[i] = s.GetDetailTable(ctx, &table)
	}
	return res, nil
}

func (s *sqlStorage) ReadTableByEmployeeIdAndStatus(ctx context.Context, employeeId uint, status bool) ([]models.TableResponse, error) {
	var tables models.Tables
	if err := s.db.Where("employee_id = ?", employeeId).Where("available = ?", status).Find(&tables).Error; err != nil {
		fmt.Println("Error while find table by id in repository: " + err.Error())
		return nil, err
	} else if len(tables) == 0 {
		fmt.Println("No record found while find table by employeeId and status in repository: " + err.Error())
		return nil, err
	}

	var res = make([]models.TableResponse, len(tables))
	for i, table := range tables {
		res[i] = s.GetDetailTable(ctx, &table)
	}
	return res, nil
}

func (s *sqlStorage) GetDetailTable(ctx context.Context, table *models.Table) models.TableResponse {
	employeeService := services.NewEmployeeBusiness(s)
	areaService := services.NewAreaBusiness(s)
	embedEmployee, _ := employeeService.ReadEmployeeById(ctx, table.EmployeeId)
	embedArea, _ := areaService.ReadAreaById(ctx, table.AreaId)
	return models.TableResponse{Model: table.Model, NormalTable: &table.NormalTable, Available: &table.NormalTable, Capacity: &table.Capacity, BasePrice: &table.BasePrice, Orders: nil, Employee: *embedEmployee, Area: *embedArea}
}
