package models

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	Status    bool `json:"status" gorm:"column:status;default:false"`
	Quantity  int  `json:"quantity" gorm:"column:quantity;not null"`
	OrderId   uint `json:"orderId" gorm:"column:order_id;not null"`
	ProductId uint `json:"productId" gorm:"column:product_id;not null"`
}

type BillResponse struct {
	gorm.Model
	Status   bool    `json:"status"`
	Quantity int     `json:"quantity"`
	Order    Order   `json:"order" sql:"-" gorm:"-"`
	Product  Product `json:"product" sql:"-" gorm:"-"`
}

type BillCreatable struct {
	gorm.Model
	Quantity  *int  `json:"quantity" gorm:"column:quantity;not null"`
	OrderId   *uint `json:"orderId" gorm:"column:order_id;not null"`
	ProductId *uint `json:"productId" gorm:"column:product_id;not null"`
}

type BillUpdatable struct {
	gorm.Model `json:"-"`
	Status     *bool `json:"status" gorm:"column:status;default:false"`
}

type Bills []Bill

func (Bill) GetTableName() string          { return "bills" }
func (BillCreatable) GetTableName() string { return Bill{}.GetTableName() }
func (BillUpdatable) GetTableName() string { return Bill{}.GetTableName() }
func (Bills) GetTableName() string         { return Bill{}.GetTableName() }
