package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/services"
)

func (s *sqlStorage) SignUp(ctx context.Context, account *models.AccountCreatable, employee *models.EmployeeCreatable) (*uint, *uint, error) {
	repository := NewSQLStore(s.db)
	accountBusiness := services.NewAccountBusiness(repository)
	employeeBusiness := services.NewEmployeeBusiness(repository)

	if employeeId, err := employeeBusiness.CreateEmployee(ctx, employee); err != nil {
		fmt.Println("Error while create employee in signUp repository: " + err.Error())
		return nil, nil, err
	} else {
		account.EmployeeId = employeeId
		if accountId, err := accountBusiness.CreateAccount(ctx, account); err != nil {
			fmt.Println("Error while create employee in signUp repository: " + err.Error())
			return nil, nil, err
		} else {
			return employeeId, accountId, nil
		}
	}
}
