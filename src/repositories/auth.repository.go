package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/services"
	"go-food-delivery-api/src/tokens"
	"go-food-delivery-api/src/tokens/jwt"
	"os"

	exceptions "go-food-delivery-api/src/errors"

	"golang.org/x/crypto/bcrypt"
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

func (s *sqlStorage) SignIn(ctx context.Context, signIn *models.SignIn) (*tokens.Token, error) {
	repository := NewSQLStore(s.db)
	accountService := services.NewAccountBusiness(repository)
	if account, err := accountService.ReadAccountByUsername(ctx, signIn.Password); err != nil {
		fmt.Println("Error while read account by username in repository: " + err.Error())
		return nil, err
	} else if account == nil {
		return nil, exceptions.ErrUserNotFound
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(signIn.Password)); err != nil {
			fmt.Println("Error while compare user password in repository: " + err.Error())
			return nil, err
		} else {
			secretKey := os.Getenv("JWT_ACCESS_SECRET")
			payload := tokens.TokenPayload{AccountId: int(account.ID), RoleId: int(account.RoleId), EmployeeId: int(account.EmployeeId)}
			jwtProvider := jwt.NewJWTProvider(secretKey)
			if accessToken, err := jwtProvider.Generate(payload, 86400); err != nil {
				fmt.Println("Error while try go generate accessToken in repository: " + err.Error())
				return nil, err
			} else {
				return accessToken, nil
			}
		}
	}
}

func (s *sqlStorage) CreateResetPassword(ctx context.Context, resetPassword *models.ResetPasswordCreatable) (*uint, error) {
	currentAccountId := ctx.Value("accountId").(int)
	resetPassword.ManagerId = &currentAccountId
	repository := NewSQLStore(s.db)
	accountService := services.NewAccountBusiness(repository)
	if account, err := accountService.ReadAccountByUsername(ctx, *resetPassword.Email); err != nil {
		fmt.Println("Error while reset password in reset password in repository: " + err.Error())
		return nil, err
	} else {
		hashPasswordBytes, _ := bcrypt.GenerateFromPassword([]byte(account.Email), 5)
		hashPassword := string(hashPasswordBytes)
		accountUpdatable := accountService.GetUpdatableAccount(ctx, account)
		accountUpdatable.Password = &hashPassword
		if _, err := accountService.UpdateAccount(ctx, *resetPassword.Email, &accountUpdatable); err != nil {
			fmt.Println("Error while update account in reset password repository: " + err.Error())
			return nil, err
		} else if result := s.db.Table(models.ResetPasswordCreatable{}.GetTableName()).Create(resetPassword); result.Error != nil {
			fmt.Println("Error while create record reset password in repository: " + result.Error.Error())
			return nil, result.Error
		} else {
			return &resetPassword.ID, nil
		}
	}
}
