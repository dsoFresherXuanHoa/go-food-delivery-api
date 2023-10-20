package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AccountStorage interface {
	CreateAccount(ctx context.Context, account *models.AccountCreatable) (*uint, error)
	ReadAccountById(ctx context.Context, id uint) (*models.Account, error)
	ReadAccountByUsername(ctx context.Context, username string) (*models.Account, error)
	ReadAccountByEmail(ctx context.Context, email string) (*models.Account, error)
	GetUpdatableAccount(ctx context.Context, account *models.Account) models.AccountUpdatable
	UpdateAccount(ctx context.Context, username string, account *models.AccountUpdatable) (*int64, error)
}

type accountBusiness struct {
	storage AccountStorage
}

func NewAccountBusiness(storage AccountStorage) *accountBusiness {
	return &accountBusiness{storage: storage}
}

func (business *accountBusiness) CreateAccount(ctx context.Context, account *models.AccountCreatable) (*uint, error) {
	hashPasswordBytes, _ := bcrypt.GenerateFromPassword([]byte(*account.Password), 5)
	hashPassword := string(hashPasswordBytes)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	secretCode := r.Intn(9000) + 1000

	account.Password = &hashPassword
	account.SecretCode = &secretCode
	if id, err := business.storage.CreateAccount(ctx, account); err != nil {
		fmt.Println("Error while save account in service: " + err.Error())
		return nil, err
	} else {
		return id, nil
	}
}

func (business *accountBusiness) ReadAccountByUsername(ctx context.Context, username string) (*models.Account, error) {
	if account, err := business.storage.ReadAccountByUsername(ctx, username); err != nil {
		fmt.Println("Error while read account by username in service: " + err.Error())
		return nil, err
	} else {
		return account, nil
	}
}

func (business *accountBusiness) ReadAccountById(ctx context.Context, id uint) (*models.Account, error) {
	if account, err := business.storage.ReadAccountById(ctx, id); err != nil {
		fmt.Println("Error while read account by id in service: " + err.Error())
		return nil, err
	} else {
		return account, nil
	}
}

func (business *accountBusiness) ReadAccountByEmail(ctx context.Context, email string) (*models.Account, error) {
	if account, err := business.storage.ReadAccountByEmail(ctx, email); err != nil {
		fmt.Println("Error while read account by email in service: " + err.Error())
		return nil, err
	} else {
		return account, nil
	}
}

func (business *accountBusiness) UpdateAccount(ctx context.Context, username string, account *models.AccountUpdatable) (*int64, error) {
	if accountId, err := business.storage.UpdateAccount(ctx, username, account); err != nil {
		fmt.Println("Error while update account service: " + err.Error())
		return nil, err
	} else {
		return accountId, nil
	}
}

func (business *accountBusiness) GetUpdatableAccount(ctx context.Context, account *models.Account) models.AccountUpdatable {
	return business.storage.GetUpdatableAccount(ctx, account)
}
