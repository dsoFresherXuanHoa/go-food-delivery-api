package models

import (
	"gorm.io/gorm"
)

type ResetPassword struct {
	gorm.Model
	ManagerId int    `json:"-" gorm:"column:managerId;not null"`
	Username  string `json:"-" gorm:"column:username;not null"`
}

type ResetPasswordCreatable struct {
	gorm.Model
	ManagerId *int    `json:"-" gorm:"column:managerId;not null"`
	Username  *string `json:"username" gorm:"column:username;not null"`
}

type ResetPasswords []ResetPassword

func (ResetPassword) GetTableName() string          { return "reset_passwords" }
func (ResetPasswordCreatable) GetTableName() string { return ResetPassword{}.GetTableName() }
func (ResetPasswords) GetTableName() string         { return ResetPassword{}.GetTableName() }
