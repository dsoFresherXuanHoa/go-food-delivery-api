package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Account  Account `json:"-"`
	Tables   Tables  `json:"-"`
	Orders   Orders  `json:"-"`
	FullName string  `json:"fullName" gorm:"column:full_name;not null"`
	Tel      string  `json:"tel" gorm:"column:tel;not null"`
	Gender   bool    `json:"gender" gorm:"column:gender;default:false"`
}

type EmployeeResponse struct {
	gorm.Model
	Account  Account `json:"account"`
	Tables   Tables  `json:"tables"`
	Orders   Orders  `json:"orders"`
	FullName string  `json:"fullName"`
	Tel      string  `json:"tel"`
	Gender   bool    `json:"gender"`
}

type EmployeeCreatable struct {
	gorm.Model
	FullName *string `json:"fullName" gorm:"column:full_name;not null"`
	Tel      *string `json:"tel" gorm:"column:tel;not null"`
	Gender   *bool   `json:"gender" gorm:"column:gender;default:false"`
}

type EmployeeUpdatable struct {
	gorm.Model
	FullName *string `json:"fullName" gorm:"column:full_name;not null"`
	Tel      *string `json:"tel" gorm:"column:tel;not null"`
}

type Employees []Employee

func (Employee) GetTableName() string          { return "employees" }
func (EmployeeCreatable) GetTableName() string { return Employee{}.GetTableName() }
func (EmployeeUpdatable) GetTableName() string { return Employee{}.GetTableName() }
func (Employees) GetTableName() string         { return Employee{}.GetTableName() }
