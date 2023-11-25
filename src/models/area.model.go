package models

import (
	"gorm.io/gorm"
)

type Area struct {
	gorm.Model
	Tables Tables `json:"-"`
	Pos    string `json:"pos" gorm:"column:pos;not null"`
}

type AreaResponse struct {
	gorm.Model
	Tables Tables `json:"table"`
	Pos    string `json:"pos"`
}

type AreaCreatable struct {
	gorm.Model `json:"-"`
	Pos        *string `json:"pos" gorm:"column:pos;not null"`
}

type AreaUpdatable struct {
	gorm.Model
	Pos *string `json:"pos" gorm:"column:pos;not null"`
}

type Areas []Area

func (Area) GetTableName() string          { return "areas" }
func (AreaCreatable) GetTableName() string { return Area{}.GetTableName() }
func (AreaUpdatable) GetTableName() string { return Area{}.GetTableName() }
func (Areas) GetTableName() string         { return Area{}.GetTableName() }
