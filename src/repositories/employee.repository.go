package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
	"time"
)

func (s *sqlStorage) CreateEmployee(ctx context.Context, employee *models.EmployeeCreatable) (*uint, error) {
	if err := s.db.Table(models.EmployeeCreatable{}.GetTableName()).Create(&employee).Error; err != nil {
		fmt.Println("Error while save employee in repository: " + err.Error())
		return nil, err
	}
	return &employee.ID, nil
}

func (s *sqlStorage) ReadAllEmployee(ctx context.Context) (models.Employees, error) {
	var employees models.Employees
	if err := s.db.Table(models.Employees{}.GetTableName()).Find(&employees).Error; err != nil {
		fmt.Println("Error while read all employee in repository: " + err.Error())
		return nil, err
	}
	return employees, nil
}

func (s *sqlStorage) ReadEmployeeById(ctx context.Context, id uint) (*models.Employee, error) {
	var employee models.Employee
	if err := s.db.Where("id = ?", id).First(&employee).Error; err != nil {
		fmt.Println("Error while read employee by id in repository: " + err.Error())
		return nil, err
	}
	return &employee, nil
}

func (s *sqlStorage) ReadTopEmployeeOrderNumber(ctx context.Context, startTime int, endTime int) ([]models.StatsTopEmployeeByOrder, error) {
	var res []models.StatsTopEmployeeByOrder
	startTimeUnix := time.Unix(int64(startTime), 0)
	endTimeUnix := time.Unix(int64(endTime), 0)
	if err := s.db.Table(models.StatsTopEmployeeByOrder{}.GetTableName()).Group("employee_id").Order("total_orders DESC").Select("employee_id, COUNT(*) AS total_orders").Where("created_at >= ? AND created_at <= ?", startTimeUnix, endTimeUnix).Find(&res).Error; err != nil {
		fmt.Println("Error while get all number order by employeeId: " + err.Error())
		return nil, err
	}
	return res, nil
}
