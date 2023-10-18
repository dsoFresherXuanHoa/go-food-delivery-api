package models

type BookingItemResponse struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}
