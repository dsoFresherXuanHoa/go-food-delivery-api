package models

import (
	"gorm.io/gorm"
)

type ResetPassword struct {
	gorm.Model `json:"-"`
	ManagerId  int    `json:"-" gorm:"column:managerId;not null"`
	Email      string `json:"-" gorm:"column:email;not null"`
}

type ResetPasswordCreatable struct {
	gorm.Model `json:"-"`
	ManagerId  *int    `json:"-" gorm:"column:managerId;not null"`
	Email      *string `json:"email" gorm:"column:email;not null"`
}

type ResetPasswords []ResetPassword

func (ResetPassword) GetTableName() string          { return "reset_passwords" }
func (ResetPasswordCreatable) GetTableName() string { return ResetPassword{}.GetTableName() }
func (ResetPasswords) GetTableName() string         { return ResetPassword{}.GetTableName() }
