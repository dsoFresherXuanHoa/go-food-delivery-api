package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Bills         Bills   `json:"-"`
	Name          string  `json:"name" gorm:"column:name;not null"`
	Description   string  `json:"description" gorm:"column:description;not null"`
	Price         float64 `json:"price" gorm:"column:price;not null"`
	DiscountPrice float64 `json:"discountPrice" sql:"-" gorm:"-"`
	Thumb         string  `json:"thumb" gorm:"column:thumb"`
	ReorderLevel  int     `json:"reorderLevel" gorm:"column:reorder_level;default:0"`
	StockAmount   int     `json:"stockAmount" gorm:"column:stock_amount;default:0"`
	CategoryId    uint    `json:"categoryId" gorm:"column:category_id;not null"`
	DiscountId    uint    `json:"discountId" gorm:"column:discount_id;not null"`
}

type ProductResponse struct {
	gorm.Model
	Bills         Bills    `json:"-" sql:"-" gorm:"-"`
	Name          *string  `json:"name"`
	Description   *string  `json:"description"`
	Price         *float64 `json:"price"`
	DiscountPrice *float64 `json:"discountPrice"`
	Thumb         *string  `json:"thumb"`
	ReorderLevel  *int     `json:"reorderLevel"`
	StockAmount   *int     `json:"stockAmount"`
	Category      Category `json:"category" sql:"-" gorm:"-"`
	Discount      Discount `json:"discount" sql:"-" gorm:"-"`
}

type ProductCreatable struct {
	gorm.Model
	Name        *string `json:"name" gorm:"column:name;not null"`
	Description *string `json:"description" gorm:"column:description;not null"`
	Price       float64 `json:"price" gorm:"column:price;not null"`
	Thumb       *string `json:"thumb" gorm:"column:thumb"`
	StockAmount *int    `json:"stockAmount" gorm:"column:stock_amount;default:0"`
	CategoryId  *uint   `json:"categoryId" gorm:"column:category_id;not null"`
	DiscountId  *uint   `json:"discountId" gorm:"column:discount_id;not null"`
}

type ProductUpdatable struct {
	gorm.Model
	Name         *string `json:"name" gorm:"column:name;not null"`
	Description  *string `json:"description" gorm:"column:description;not null"`
	ReorderLevel *int    `json:"reorderLevel" gorm:"column:reorder_level;default:0"`
	StockAmount  *int    `json:"stockAmount" gorm:"column:stock_amount;default:0"`
}

type Products []Product

func (Product) GetTableName() string          { return "products" }
func (ProductCreatable) GetTableName() string { return Product{}.GetTableName() }
func (ProductUpdatable) GetTableName() string { return Product{}.GetTableName() }
func (Products) GetTableName() string         { return Product{}.GetTableName() }
