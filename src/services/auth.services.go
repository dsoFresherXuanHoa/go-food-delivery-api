package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/tokens"
)

type AuthStorage interface {
	SignUp(ctx context.Context, account *models.AccountCreatable, employee *models.EmployeeCreatable) (*uint, *uint, error)
	SignIn(ctx context.Context, signIn *models.SignIn) (*tokens.Token, error)
	ResetPassword(ctx context.Context, resetPassword *models.ResetPasswordCreatable) (*uint, error)
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
	} else
	/*
		if err := utils.SendSignUpSecretCode(*account.Email, nil, *account); err != nil {
			fmt.Println("Error while send email to user email in service: " + err.Error())
			return nil, nil, err
		} else
	*/
	{
		return employeeId, accountId, nil
	}
}

func (business *authBusiness) SignIn(ctx context.Context, signIn *models.SignIn) (*tokens.Token, error) {
	if token, err := business.storage.SignIn(ctx, signIn); err != nil {
		fmt.Println("Error while sign in in service: " + err.Error())
		return nil, err
	} else {
		return token, nil
	}
}

func (business *authBusiness) ResetPassword(ctx context.Context, resetPassword *models.ResetPasswordCreatable) (*uint, error) {
	if resetPasswordId, err := business.storage.ResetPassword(ctx, resetPassword); err != nil {
		fmt.Println("Error while reset password in auth service: " + err.Error())
		return nil, err
	} else {
		return resetPasswordId, nil
	}
}
