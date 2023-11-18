package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type EmployeeStorage interface {
	CreateEmployee(ctx context.Context, employee *models.EmployeeCreatable) (*uint, error)
	ReadEmployeeById(ctx context.Context, id uint) (*models.Employee, error)
	ReadAllEmployee(ctx context.Context) ([]models.EmployeeResponse, error)
	ReadTopEmployeeOrderNumber(ctx context.Context, startTime int, endTime int) ([]models.StatsTopEmployeeByOrder, error)
	DeleteEmployeeById(ctx context.Context, employeeId int) (*int, error)
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

func (business *employeeBusiness) ReadAllEmployee(ctx context.Context) ([]models.EmployeeResponse, error) {
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

func (business *employeeBusiness) DeleteEmployeeById(ctx context.Context, employeeId int) (*int, error) {
	if deletedEmployee, err := business.storage.DeleteEmployeeById(ctx, employeeId); err != nil {
		fmt.Println("Error while delete employee by id in service: " + err.Error())
		return nil, err
	} else {
		return deletedEmployee, nil
	}
}
