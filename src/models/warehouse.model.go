package models

type Warehouse struct {
	Name            string  `form:"name" json:"name"`
	Description     string  `form:"description" json:"description"`
	Thumb           string  `form:"-" json:"thumb"`
	Price           float64 `form:"price" json:"price"`
	StockAmount     int     `form:"stockAmount" json:"stockAmount"`
	CategoryId      uint    `form:"categoryId" json:"categoryId"`
	MinQuantity     int     `form:"minQuantity" json:"minQuantity"`
	DiscountPercent int     `form:"discountPercent" json:"discountPercent"`
	Uint            string  `form:"unit" json:"uint"`
}

type WarehouseUpdatable struct {
	Name            string  `form:"name" json:"name"`
	Description     string  `form:"description" json:"description"`
	Price           float64 `form:"price" json:"price"`
	StockAmount     int     `form:"stockAmount" json:"stockAmount"`
	MinQuantity     int     `form:"minQuantity" json:"minQuantity"`
	DiscountPercent int     `form:"discountPercent" json:"discountPercent"`
	Uint            string  `form:"unit" json:"uint"`
}
