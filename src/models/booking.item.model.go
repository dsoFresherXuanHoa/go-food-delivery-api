package models

type BookingItemResponse struct {
	Product  ProductResponse `json:"product"`
	Quantity int             `json:"quantity"`
}
