package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Bills      Bills  `json:"-"`
	Note       string `json:"note" gorm:"column:note"`
	Status     bool   `json:"status" gorm:"column:status;default:false"`
	Accepted   bool   `json:"accepted" gorm:"column:accepted;default:false"`
	Compensate bool   `json:"compensate" gorm:"column:compensate;default:false"`
	EmployeeId uint   `json:"employeeId" gorm:"column:employee_id;not null"`
	TableId    uint   `json:"tableId" gorm:"column:table_id;not null"`
}

type OrderResponse struct {
	gorm.Model
	Bills      Bills    `json:"bills" sql:"-" gorm:"-"`
	Note       string   `json:"note"`
	Status     bool     `json:"status"`
	Accepted   bool     `json:"accepted"`
	Compensate bool     `json:"compensate"`
	Employee   Employee `json:"employee" sql:"-" gorm:"-"`
	Table      Table    `json:"table" sql:"-" gorm:"-"`
}

type OrderCreatable struct {
	gorm.Model
	Note       *string `json:"note" gorm:"column:note"`
	EmployeeId *uint   `json:"employeeId" gorm:"column:employee_id;not null"`
	TableId    uint    `json:"tableId" gorm:"column:table_id;not null"`
}

type OrderUpdatable struct {
	gorm.Model
	Note       *string `json:"note" gorm:"column:note"`
	Status     *bool   `json:"-" gorm:"column:status;default:false"`
	Accepted   *bool   `json:"_" gorm:"column:accepted;default:false"`
	Compensate *bool   `json:"-" gorm:"column:compensate;default:false"`
}

type Orders []Order

func (Order) GetTableName() string          { return "orders" }
func (OrderCreatable) GetTableName() string { return Order{}.GetTableName() }
func (OrderUpdatable) GetTableName() string { return Order{}.GetTableName() }
func (Orders) GetTableName() string         { return Order{}.GetTableName() }
