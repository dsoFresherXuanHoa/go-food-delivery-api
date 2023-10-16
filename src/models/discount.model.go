package models

import (
	"gorm.io/gorm"
)

type Discount struct {
	gorm.Model
	Products        Products `json:"-"`
	MinQuantity     int      `json:"minQuantity" gorm:"column:min_quantity;not null;default:0"`
	DiscountPercent int      `json:"discountPercent" gorm:"column:discount_percent;default:0;not null"`
}

type DiscountResponse struct {
	gorm.Model
	Products        Products `json:"products"`
	MinQuantity     int      `json:"minQuantity"`
	DiscountPercent int      `json:"discountPercent"`
}

type DiscountCreatable struct {
	gorm.Model
	MinQuantity     *int `json:"minQuantity" gorm:"column:min_quantity;not null;default:0"`
	DiscountPercent *int `json:"discountPercent" gorm:"column:discount_percent;default:0;not null"`
}

type DiscountUpdatable struct {
	gorm.Model
	MinQuantity     int `json:"minQuantity" gorm:"column:min_quantity;not null;default:0"`
	DiscountPercent int `json:"discountPercent" gorm:"column:discount_percent;default:0;not null"`
}

type Discounts []Discount

func (Discount) GetTableName() string          { return "discounts" }
func (DiscountCreatable) GetTableName() string { return Discount{}.GetTableName() }
func (DiscountUpdatable) GetTableName() string { return Discount{}.GetTableName() }
func (Discounts) GetTableName() string         { return Discount{}.GetTableName() }
