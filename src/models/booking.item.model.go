package models

type BookingItemResponse struct {
	BillId   uint            `json:"billId"`
	Product  ProductResponse `json:"product"`
	Quantity int             `json:"quantity"`
	Status   bool            `json:"status"`
}
