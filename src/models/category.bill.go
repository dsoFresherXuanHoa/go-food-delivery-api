package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Products    Products `json:"-"`
	Name        string   `json:"name" gorm:"column:name;not null"`
	Description string   `json:"description" gorm:"column:description;not null"`
	Thumb       string   `json:"thumb" gorm:"column:thumb"`
}

type CategoryResponse struct {
	gorm.Model
	Products    Products `json:"products"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Thumb       string   `json:"thumb"`
}

type CategoryCreatable struct {
	gorm.Model
	Name        *string `json:"name" gorm:"column:name;not null"`
	Description *string `json:"description" gorm:"column:description;not null"`
	Thumb       *string `json:"thumb" gorm:"column:thumb"`
}

type CategoryUpdatable struct {
	gorm.Model
	Name        *string `json:"name" gorm:"column:name;not null"`
	Description *string `json:"description" gorm:"column:description;not null"`
	Thumb       *string `json:"thumb" gorm:"column:thumb"`
}

type Categories []Category

func (Category) GetTableName() string          { return "categories" }
func (CategoryCreatable) GetTableName() string { return Category{}.GetTableName() }
func (CategoryUpdatable) GetTableName() string { return Category{}.GetTableName() }
func (Categories) GetTableName() string        { return Category{}.GetTableName() }
