package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

func (s *sqlStorage) CreateAccount(ctx context.Context, account *models.AccountCreatable) (*uint, error) {
	if err := s.db.Table(models.AccountCreatable{}.GetTableName()).Create(&account).Error; err != nil {
		fmt.Println("Error while save account in repository: " + err.Error())
		return nil, err
	}
	return &account.ID, nil
}

func (s *sqlStorage) ReadAccountById(ctx context.Context, id uint) (*models.Account, error) {
	var account models.Account
	if err := s.db.Unscoped().Where("id = ?", id).First(&account).Error; err != nil {
		fmt.Println("Error while read account by id in repository: " + err.Error())
		return nil, err
	}
	return &account, nil
}

func (s *sqlStorage) ReadAccountByUsername(ctx context.Context, username string) (*models.Account, error) {
	var account models.Account
	if err := s.db.Where("email = ?", username).First(&account).Error; err != nil {
		fmt.Println("Error while read account by username in repository: " + err.Error())
		return nil, err
	}
	return &account, nil
}

func (s *sqlStorage) ReadAccountByEmail(ctx context.Context, email string) (*models.Account, error) {
	var account models.Account
	if err := s.db.Where("email = ?", email).First(&account).Error; err != nil {
		fmt.Println("Error while read account by email in repository: " + err.Error())
		return nil, err
	}
	return &account, nil
}

func (s *sqlStorage) GetUpdatableAccount(ctx context.Context, account *models.Account) models.AccountUpdatable {
	return models.AccountUpdatable{Model: account.Model, SecretCode: &account.SecretCode, Password: &account.Password}
}

func (s *sqlStorage) UpdateAccount(ctx context.Context, username string, account *models.AccountUpdatable) (*int64, error) {
	if result := s.db.Table(models.AccountUpdatable{}.GetTableName()).Where("email = ?", username).Updates(account); result.Error != nil {
		fmt.Println("Error while update account in repository: " + result.Error.Error())
		return nil, result.Error
	} else {
		return &result.RowsAffected, nil
	}
}

func (s *sqlStorage) ReadAllAccount(ctx context.Context) (models.Accounts, error) {
	var accounts models.Accounts
	if err := s.db.Unscoped().Table(models.Account{}.GetTableName()).Find(&accounts).Error; err != nil {
		fmt.Println("Error while read all accounts in repository: " + err.Error())
		return nil, err
	}
	return accounts, nil
}

func (s *sqlStorage) DeleteAccountById(ctx context.Context, accountId int) (*int, error) {
	var account models.Account
	if err := s.db.Table(models.Account{}.GetTableName()).Where("id = ?", accountId).Delete(&account).Error; err != nil {
		fmt.Println("Error while delete account by id: " + err.Error())
		return nil, err
	}
	return &accountId, nil
}
