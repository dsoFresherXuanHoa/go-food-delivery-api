package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/utils"
)

type AuthStorage interface {
	SignUp(ctx context.Context, account *models.AccountCreatable, employee *models.EmployeeCreatable) (*uint, *uint, error)
}

type authBusiness struct {
	storage AuthStorage
}

func NewAuthBusiness(storage AuthStorage) *authBusiness {
	return &authBusiness{storage: storage}
}

func (business *authBusiness) SignUp(ctx context.Context, account *models.AccountCreatable, employee *models.EmployeeCreatable) (*uint, *uint, error) {
	if employeeId, accountId, err := business.storage.SignUp(ctx, account, employee); err != nil {
		fmt.Println("Error while sign up in service: " + err.Error())
		return nil, nil, err
	} else if err := utils.SendSignUpSecretCode(*account.Email, nil, *account); err != nil {
		fmt.Println("Error while send email to user email in service: " + err.Error())
		return nil, nil, err
	} else {
		return employeeId, accountId, nil
	}
}
