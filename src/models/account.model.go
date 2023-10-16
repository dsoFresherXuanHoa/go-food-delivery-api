package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username   string `json:"-" gorm:"column:username;unique;not null"`
	Password   string `json:"-" gorm:"column:password;not null"`
	Email      string `json:"-" gorm:"column:email;unique;not null"`
	SecretCode int    `json:"-" gorm:"column:secret_code;not null"`
	EmployeeId uint   `json:"-" gorm:"column:employee_id;unique;not null"`
	RoleId     uint   `json:"-" gorm:"column:role_id;not null"`
}

type AccountResponse struct {
	gorm.Model
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Email      string   `json:"email"`
	SecretCode int      `json:"secretCode"`
	Employee   Employee `json:"employee"`
	Role       Role     `json:"role"`
}

type AccountCreatable struct {
	gorm.Model
	Username   *string `json:"username" gorm:"column:username;unique;not null"`
	Password   *string `json:"password" gorm:"column:password;not null"`
	Email      *string `json:"email" gorm:"column:email;unique;not null"`
	SecretCode *int    `json:"-" gorm:"column:secret_code;not null"`
	EmployeeId *uint   `json:"employeeId" gorm:"column:employee_id;unique;not null"`
	RoleId     *uint   `json:"roleId" gorm:"column:role_id;not null"`
}

type AccountUpdatable struct {
	gorm.Model
	SecretCode *int    `json:"-" gorm:"column:secret_code;not null"`
	Password   *string `json:"password" gorm:"column:password;not null"`
}

type Accounts []Account

func (Account) GetTableName() string          { return "accounts" }
func (AccountCreatable) GetTableName() string { return Account{}.GetTableName() }
func (AccountUpdatable) GetTableName() string { return Account{}.GetTableName() }
func (Accounts) GetTableName() string         { return Account{}.GetTableName() }
