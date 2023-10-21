package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type EmployeeStorage interface {
	CreateEmployee(ctx context.Context, employee *models.EmployeeCreatable) (*uint, error)
	ReadEmployeeById(ctx context.Context, id uint) (*models.Employee, error)
	ReadAllEmployee(ctx context.Context) (models.Employees, error)
	ReadTopEmployeeOrderNumber(ctx context.Context, startTime int, endTime int) ([]models.StatsTopEmployeeByOrder, error)
}

type employeeBusiness struct {
	storage EmployeeStorage
}

func NewEmployeeBusiness(storage EmployeeStorage) *employeeBusiness {
	return &employeeBusiness{storage: storage}
}

func (business *employeeBusiness) CreateEmployee(ctx context.Context, employee *models.EmployeeCreatable) (*uint, error) {
	if id, err := business.storage.CreateEmployee(ctx, employee); err != nil {
		fmt.Println("Error while save employee in service: " + err.Error())
		return nil, err
	} else {
		return id, nil
	}
}

func (business *employeeBusiness) ReadEmployeeById(ctx context.Context, id uint) (*models.Employee, error) {
	if id, err := business.storage.ReadEmployeeById(ctx, id); err != nil {
		fmt.Println("Error while find employee by id in service: " + err.Error())
		return nil, err
	} else {
		return id, nil
	}
}

func (business *employeeBusiness) ReadAllEmployee(ctx context.Context) (models.Employees, error) {
	if employees, err := business.storage.ReadAllEmployee(ctx); err != nil {
		fmt.Println("Error while read all employee id in service: " + err.Error())
		return nil, err
	} else {
		return employees, nil
	}
}

func (business *employeeBusiness) ReadTopEmployeeOrderNumber(ctx context.Context, startTime int, endTime int) ([]models.StatsTopEmployeeByOrder, error) {
	if employees, err := business.storage.ReadTopEmployeeOrderNumber(ctx, startTime, endTime); err != nil {
		fmt.Println("Error while read number order by employeeId in service: " + err.Error())
		return nil, err
	} else {
		return employees, nil
	}
}
