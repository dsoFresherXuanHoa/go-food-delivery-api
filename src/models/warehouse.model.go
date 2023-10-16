package models

type Warehouse struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Thumb           string  `json:"thumb"`
	Price           float64 `json:"price"`
	StockAmount     int     `json:"stockAmount"`
	CategoryId      uint    `json:"categoryId"`
	MinQuantity     int     `json:"minQuantity"`
	DiscountPercent int     `json:"discountPercent"`
}
