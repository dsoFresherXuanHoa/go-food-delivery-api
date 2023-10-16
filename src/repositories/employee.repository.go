package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

func (s *sqlStorage) CreateEmployee(ctx context.Context, employee *models.EmployeeCreatable) (*uint, error) {
	if err := s.db.Table(models.EmployeeCreatable{}.GetTableName()).Create(&employee).Error; err != nil {
		fmt.Println("Error while save employee in repository: " + err.Error())
		return nil, err
	}
	return &employee.ID, nil
}

func (s *sqlStorage) ReadEmployeeById(ctx context.Context, id uint) (*models.Employee, error) {
	var employee models.Employee
	if err := s.db.Where("id = ?", id).First(&employee).Error; err != nil {
		fmt.Println("Error while read employee by id in repository: " + err.Error())
		return nil, err
	}
	return &employee, nil
}
