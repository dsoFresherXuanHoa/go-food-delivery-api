package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/services"
	"time"

	"gorm.io/gorm"
)

func (s *sqlStorage) CreateEmployee(ctx context.Context, employee *models.EmployeeCreatable) (*uint, error) {
	if err := s.db.Table(models.EmployeeCreatable{}.GetTableName()).Create(&employee).Error; err != nil {
		fmt.Println("Error while save employee in repository: " + err.Error())
		return nil, err
	}
	return &employee.ID, nil
}

func (s *sqlStorage) GetDetailEmployee(ctx context.Context, employee models.Employee) models.EmployeeResponse {
	accountService := services.NewAccountBusiness(s)
	if embedAccount, err := accountService.ReadAccountById(ctx, employee.ID); err != nil {
		return models.EmployeeResponse{Model: gorm.Model{}, Account: models.Account{}, Tables: nil, Orders: nil, FullName: employee.FullName, Tel: employee.Tel, Gender: employee.Gender}
	} else {
		return models.EmployeeResponse{Model: embedAccount.Model, Account: *embedAccount, Tables: nil, Orders: nil, FullName: employee.FullName, Tel: employee.Tel, Gender: employee.Gender}
	}
}

func (s *sqlStorage) ReadAllEmployee(ctx context.Context) ([]models.EmployeeResponse, error) {
	var employees models.Employees
	if err := s.db.Unscoped().Table(models.Employees{}.GetTableName()).Find(&employees).Error; err != nil {
		fmt.Println("Error while read all employee in repository: " + err.Error())
		return nil, err
	}
	var res = make([]models.EmployeeResponse, len(employees))
	for i, employee := range employees {
		res[i] = s.GetDetailEmployee(ctx, employee)
	}
	return res, nil
}

func (s *sqlStorage) ReadEmployeeById(ctx context.Context, id uint) (*models.Employee, error) {
	var employee models.Employee
	if err := s.db.Unscoped().Where("id = ?", id).First(&employee).Error; err != nil {
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

func (s *sqlStorage) DeleteEmployeeById(ctx context.Context, employeeId int) (*int, error) {
	var employee models.Employee
	var deletedAccount models.Account
	var account models.Account
	if err := s.db.Table(models.Employee{}.GetTableName()).Where("id = ?", employeeId).Delete(&employee).Error; err != nil {
		fmt.Println("Error while delete employee by id: " + err.Error())
		return nil, err
	} else if res := s.db.Table(models.Account{}.GetTableName()).Where("employee_id = ?", employeeId).First(&deletedAccount); res.Error != nil {
		fmt.Println("Error while delete account of employee after deleting account: " + res.Error.Error())
		return nil, res.Error
	} else if err := s.db.Table(models.Account{}.GetTableName()).Where("id = ?", deletedAccount.ID).Delete(&account).Error; err != nil {
		fmt.Println("Error while delete account by id: " + err.Error())
		return nil, err
	}
	return &employeeId, nil
}
