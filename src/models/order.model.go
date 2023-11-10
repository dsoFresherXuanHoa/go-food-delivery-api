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
	Rejected   bool   `json:"rejected" gorm:"column:rejected;default:false"`
	EmployeeId uint   `json:"employeeId" gorm:"column:employee_id;not null"`
	TableId    uint   `json:"tableId" gorm:"column:table_id;not null"`
	Refundable bool   `json:"refundable" gorm:"column:refundable;not null"`
	Reason     string `json:"reason" gorm:"column:reason"`
}

type OrderResponse struct {
	gorm.Model
	Bills      Bills    `json:"bills" sql:"-" gorm:"-"`
	Note       string   `json:"note"`
	Status     bool     `json:"status"`
	Accepted   bool     `json:"accepted"`
	Rejected   bool     `json:"rejected"`
	Employee   Employee `json:"employee" sql:"-" gorm:"-"`
	Table      Table    `json:"table" sql:"-" gorm:"-"`
	Refundable bool     `json:"refundable"`
	Reason     string   `json:"reason"`
}

type OrderCreatable struct {
	gorm.Model
	Note       *string `json:"note" gorm:"column:note"`
	Status     *bool   `json:"_" gorm:"column:status;default:false"`
	Accepted   bool    `json:"-" gorm:"column:accepted;default:false"`
	EmployeeId *uint   `json:"employeeId" gorm:"column:employee_id;not null"`
	TableId    *uint   `json:"tableId" gorm:"column:table_id;not null"`
	Refundable *bool   `json:"refundable" gorm:"column:refundable;not null"`
	Reason     *string `json:"reason" gorm:"column:reason"`
}

type OrderUpdatable struct {
	gorm.Model
	Note     *string `json:"note" gorm:"column:note"`
	Status   *bool   `json:"-" gorm:"column:status;default:false"`
	Accepted *bool   `json:"-" gorm:"column:accepted;default:false"`
	Rejected *bool   `json:"-" gorm:"column:rejected;default:false"`
	Reason   *string `json:"reason" gorm:"column:reason"`
}

type Orders []Order

func (Order) GetTableName() string          { return "orders" }
func (OrderCreatable) GetTableName() string { return Order{}.GetTableName() }
func (OrderUpdatable) GetTableName() string { return Order{}.GetTableName() }
func (Orders) GetTableName() string         { return Order{}.GetTableName() }
