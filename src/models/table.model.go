package models

import (
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	Orders      Orders  `json:"-"`
	NormalTable bool    `json:"normalTable" gorm:"column:normal_table;default:true"`
	Available   bool    `json:"available" gorm:"column:available;default:true"`
	Capacity    int     `json:"capacity" gorm:"column:capacity;not null"`
	BasePrice   float64 `json:"basePrice" gorm:"column:base_price;default:20000"`
	EmployeeId  uint    `json:"employeeId" gorm:"column:employee_id;not null"`
	AreaId      uint    `json:"areaId" gorm:"column:area_id;not null"`
}

type TableResponse struct {
	gorm.Model
	Orders      Orders   `json:"-" sql:"-" gorm:"-"`
	NormalTable *bool    `json:"normalTable"`
	Available   *bool    `json:"available"`
	Capacity    *int     `json:"capacity"`
	BasePrice   *float64 `json:"basePrice"`
	Employee    Employee `json:"employee" sql:"-" gorm:"-"`
	Area        Area     `json:"area" sql:"-" gorm:"-"`
}

type TableCreatable struct {
	gorm.Model  `json:"-"`
	NormalTable *bool    `json:"normalTable" gorm:"column:normal_table;default:true"`
	Capacity    *int     `json:"capacity" gorm:"column:capacity;not null"`
	BasePrice   *float64 `json:"basePrice" gorm:"column:base_price;default:20000"`
	EmployeeId  *uint    `json:"employeeId" gorm:"column:employee_id;not null"`
	AreaId      *uint    `json:"areaId" gorm:"column:area_id;not null"`
}

type TableUpdatable struct {
	gorm.Model
	NormalTable *bool    `json:"normalTable" gorm:"column:normal_table;default:true"`
	Available   *bool    `json:"available" gorm:"column:available;default:true"`
	Capacity    *int     `json:"capacity" gorm:"column:capacity;not null"`
	BasePrice   *float64 `json:"basePrice" gorm:"column:base_price;default:20000"`
}

type Tables []Table

func (Table) GetTableName() string          { return "tables" }
func (TableCreatable) GetTableName() string { return Table{}.GetTableName() }
func (TableUpdatable) GetTableName() string { return Table{}.GetTableName() }
func (Tables) GetTableName() string         { return Table{}.GetTableName() }
